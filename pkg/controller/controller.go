// Copyright 2018 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controller

import (
	"fmt"
	"time"

	"github.com/cilium/cilium/pkg/lock"
	"github.com/cilium/cilium/pkg/uuid"

	"github.com/sirupsen/logrus"
)

// ControllerFunc is the function that the controller runs
type ControllerFunc func() error

// Controller is a simple pattern that allows to perform the following
// tasks:
//   - Run an operation in the background and retry until it succeeds
//   - Perform a regular sync operation in the background
//
// A controller has configurable retry intervals and will collect statistics
// on number of successful runs, number of failures, last error message,
// and last error timestamp.
//
// Controllers have a name and are tied to a Manager. The manager is typically
// bound to higher level objects such as endpoint. These higher level objects
// can then run multiple controllers to perform async tasks such as:
//  - Annotating k8s resources with values
//  - Synchronizing an object with the kvstore
//  - Any other async operation to may fail and require retries
//
// Embedding the Manager into higher level resources allows to bind controllers
// to the lifetime of that object. Controllers also have a UUID to allow
// correlating all log messages of a controller instance.
type Controller struct {
	// Name is the name of the controller
	Name string

	// DoFunc is the function that will be run until it succeeds and/or
	// using the interval RunInterval if not 0
	DoFunc ControllerFunc

	// If set to any other value than 0, will cause DoFunc to be run in the
	// specified interval. The interval starts from when the DoFunc has
	// returned last
	RunInterval time.Duration

	// ErrorRetryBaseDuration is the initial time to wait to run DoFunc
	// again on return of an error. On each consecutive error, this value
	// is multiplied by the number of consecutive errors to provide a
	// constant back off. The default is 1s.
	ErrorRetryBaseDuration time.Duration

	// NoErrorRetry when set to true, disabled retries on errors
	NoErrorRetry bool

	successCount   int
	failureCount   int
	lastError      error
	lastErrorStamp time.Time
	uuid           string
	stop           chan struct{}
	mutex          lock.RWMutex
}

// GetSuccessCount returns the number of successful controller runs
func (c *Controller) GetSuccessCount() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.successCount
}

// GetFailureCount returns the number of failed controller runs
func (c *Controller) GetFailureCount() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.failureCount
}

// GetLastError returns the last error returned
func (c *Controller) GetLastError() error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.lastError
}

// GetLastErrorTimestamp returns the last error returned
func (c *Controller) GetLastErrorTimestamp() time.Time {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.lastErrorStamp
}

func (c *Controller) runController() {
	errorRetries := 1

	for {
		var (
			err      error
			interval = c.RunInterval
		)

		if c.DoFunc != nil {
			err = c.DoFunc()
		} else {
			err = fmt.Errorf("DoFunc is nil")
		}

		if err != nil {
			log.WithFields(logrus.Fields{
				fieldControllerName:    c.Name,
				fieldUUID:              c.uuid,
				fieldConsecutiveErrors: errorRetries,
			}).WithError(err).Debug("Controller run failed")

			c.mutex.Lock()
			c.lastError = err
			c.lastErrorStamp = time.Now()
			c.failureCount++
			c.mutex.Unlock()

			if !c.NoErrorRetry {
				if c.ErrorRetryBaseDuration != time.Duration(0) {
					interval = time.Duration(errorRetries) * c.ErrorRetryBaseDuration
				} else {
					interval = time.Duration(errorRetries) * time.Second
				}

				errorRetries++
			}
		} else {
			c.mutex.Lock()
			c.successCount++
			c.mutex.Unlock()

			// reset error retries after successful attempt
			errorRetries = 1

			// If no run interval is specified, no further updates
			// are required and we can shut down the controller
			if c.RunInterval == time.Duration(0) {
				goto shutdown
			}
		}

		select {
		case <-c.stop:
			goto shutdown

		case <-time.After(interval):
		}

	}

shutdown:
	log.WithFields(logrus.Fields{
		fieldControllerName: c.Name,
		fieldUUID:           c.uuid,
	}).Debug("Shutting down controller")
}

func (c *Controller) stopController() {
	close(c.stop)
}

type controllerMap map[string]*Controller

// Manager is a list of controllers
type Manager struct {
	controllers controllerMap
	mutex       lock.RWMutex
}

// UpdateController installs or updates a controller in the manager. A
// controller is identified by its name. If a controller with the name already
// exist, the controller will be shut down and replaced with the provided
// controller. Updating a controller will cause the DoFunc to be run
// immediately regardless of any previous conditions. It will also cause any
// statistics to be reset.
func (m *Manager) UpdateController(ctrl *Controller) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.controllers == nil {
		m.controllers = controllerMap{}
	}

	if oldCtrl, ok := m.controllers[ctrl.Name]; ok {
		oldCtrl.stopController()
	}

	ctrl.stop = make(chan struct{}, 0)
	ctrl.uuid = uuid.NewUUID().String()

	m.controllers[ctrl.Name] = ctrl
	go ctrl.runController()

	log.WithFields(logrus.Fields{
		fieldControllerName: ctrl.Name,
		fieldUUID:           ctrl.uuid,
	}).Debug("Updated controller")
}

func (m *Manager) removeController(ctrl *Controller) {
	ctrl.stopController()
	delete(m.controllers, ctrl.Name)

	log.WithFields(logrus.Fields{
		fieldControllerName: ctrl.Name,
		fieldUUID:           ctrl.uuid,
	}).Debug("Removed update controller")
}

// RemoveController stops and removes a controller from the manager. If DoFunc
// is currently running, DoFunc is allowed to complete in the background.
func (m *Manager) RemoveController(name string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.controllers == nil {
		return fmt.Errorf("empty controller map")
	}

	oldCtrl, ok := m.controllers[name]
	if !ok {
		return fmt.Errorf("unable to find controller %s", name)
	}

	m.removeController(oldCtrl)

	return nil
}

// RemoveAll stops and removes all controllers of the manager
func (m *Manager) RemoveAll() {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.controllers == nil {
		return
	}

	for _, ctrl := range m.controllers {
		m.removeController(ctrl)
	}
}
