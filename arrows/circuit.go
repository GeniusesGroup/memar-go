//Copyright 2016 SabzCity
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// build ranjbar.ehsan.1379@gmail.com

package arrows

import (
	"errors"
)

// Errors that are used throughout the Arrows Circuit Breaker API.
var (
	ErrBreakerOpen    = errors.New("circuit breaker is open")
	ErrCircuitRunning = errors.New("circuit is now running")
)

//Work is repeat in circuit, and when success or
//return error consecutive it will break.
type CircuitBreaker struct {
	//How many circuit return error to open breaker ?
	ErrorThreshold int
	//How many circuit do work successfully to end circuit ?
	SuccessThreshold int
	//number of errors.
	errors int
	//number of success
	success int
	//save state of circuit
	state bool
}

//Run work in circuit.
func (circuit *CircuitBreaker) Run(work func() error) error {
	if circuit.state {
		circuit.state = false
		for {
			err := work()
			if err != nil {
				circuit.errors += 1
			} else {
				circuit.success += 1
				circuit.errors = 0
			}
			if circuit.SuccessThreshold <= circuit.success {
				circuit.state = true
				return nil
			}
			if circuit.ErrorThreshold <= circuit.errors {
				circuit.state = true
				return ErrBreakerOpen
			}
		}
	} else {
		circuit.state = true
		return ErrCircuitRunning
	}
}

//Run work in circuit with goroutine.
func (circuit *CircuitBreaker) RunAsync(work func() error, callback func(error)) {
	go func() {
		if circuit.state {
			circuit.state = false
			for {
				err := work()
				if err != nil {
					circuit.errors += 1
				} else {
					circuit.success += 1
					circuit.errors = 0
				}
				if circuit.SuccessThreshold <= circuit.success {
					callback(nil)
					circuit.state = true
					return
				}
				if circuit.ErrorThreshold <= circuit.errors {
					callback(ErrBreakerOpen)
					circuit.state = true
					return
				}
			}
		} else {
			callback(ErrCircuitRunning)
			circuit.state = true
			return
		}
	}()
}

//Reset states of circuit.
func (circuit *CircuitBreaker) Reset() error {
	if circuit.state {
		circuit.state = false
		circuit.success = 0
		circuit.errors = 0
		circuit.state = true
		return nil
	} else {
		return ErrCircuitRunning
	}
}

//Create new circuit with SuccessThreshold and
//ErrorThreshold values.
func CreateCircuit(successThreshold int, errorThreshold int) *CircuitBreaker {
	return &CircuitBreaker{
		state:            true,
		SuccessThreshold: successThreshold,
		ErrorThreshold:   errorThreshold}
}
