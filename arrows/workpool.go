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
	"time"
)

// Errors that are used throughout the Arrows Work Pool API.
var (
	ErrJobNotFunc  = errors.New("generic worker not given a func()")
	ErrJobTimedOut = errors.New("job request timed out")
)

//Implementation basic arrows worker.
type Worker interface {
	//Do something with data and return result.
	Do(interface{}) interface{}
}

//defualt worker struct of arrows.
type DefaultWorker struct {
	Job *func(interface{}) interface{}
}

func (worker *DefaultWorker) Do(data interface{}) interface{} {
	return (*worker.Job)(data)
}

//Wrapper manage worker in pool.
type WorkerWrapper struct {
	//You shoud add this Wrapper to this channel after worker do job.
	pool chan *WorkerWrapper
	//Inner worker of Wrapper.
	Worker Worker
	//Get input argument of Do() function with this channel.
	Input chan interface{}
	//return result of Do() function to this channel.
	Output chan interface{}
}

//open the Wrapper and start loop() function.
func (wrapper *WorkerWrapper) open() {
	wrapper.Input = make(chan interface{})
	wrapper.Output = make(chan interface{})
	go wrapper.loop()
}

//Loop is main function of Wrapper, this will call in a goroutine.
//In this function, you shoud add this Wrapper to pool and wait to
//recive data from Input and return result to Output.
func (wrapper *WorkerWrapper) loop() {
	for {
		//register Wrapper in pool.
		wrapper.pool <- wrapper
		//do work.
		data := <-wrapper.Input
		result := wrapper.Worker.Do(data)
		wrapper.Output <- result
	}
}

//The WorkPool is entry point you need, you can send work
//with methods synchronous and asynchronous to worker,
// or request a worker and use it personally.
type WorkPool struct {
	workers []*WorkerWrapper
	pool    chan *WorkerWrapper
}

//You can request a worker and do something with it.
//Important : you shoud send one work to worker at time and get result
//from Output channel.
func (pool *WorkPool) RequestWorker() <-chan *WorkerWrapper {
	return pool.pool
}

func (pool *WorkPool) SendWork(data interface{}) interface{} {
	worker := <-pool.pool
	worker.Input <- data
	result := <-worker.Output
	return result
}

func (pool *WorkPool) SendWorkTimed(data interface{}, timeout time.Duration) (interface{}, error) {
	before := time.Now()
	select {
	case worker := <-pool.pool:
		worker.Input <- data
		result := <-worker.Output
		return result, nil
	case <-time.After((timeout * time.Millisecond) - time.Since(before)):
		return nil, ErrJobTimedOut
	}
}

func (pool *WorkPool) SendWorkAsync(data interface{}, callback func(interface{})) {
	go func() {
		worker := <-pool.pool
		worker.Input <- data
		result := <-worker.Output
		callback(result)
	}()
}

func (pool *WorkPool) SendWorkTimedAsync(data interface{}, timeout time.Duration, callback func(interface{}, error)) {
	go func() {
		before := time.Now()
		select {
		case worker := <-pool.pool:
			worker.Input <- data
			result := <-worker.Output
			callback(result, nil)
		case <-time.After((timeout * time.Millisecond) - time.Since(before)):
			callback(nil, ErrJobTimedOut)
		}
	}()
}

//Creates a pool of workers, and takes a closure argument which is the action
//to perform for each job.
func CreatePool(numWorkers int, job func(interface{}) interface{}) *WorkPool {
	pool := WorkPool{}
	pool.pool = make(chan *WorkerWrapper, numWorkers)
	pool.workers = make([]*WorkerWrapper, numWorkers)
	for i := range pool.workers {
		newWorker := WorkerWrapper{
			Worker: &(DefaultWorker{&job}),
			pool:   pool.pool}
		pool.workers[i] = &newWorker
		pool.workers[i].open()
	}
	return &pool
}

//Creates a pool of generic workers. When sending work to a pool of
//generic workers you send a closure (func()) which is the job to perform.
func CreatePoolGeneric(numWorkers int) *WorkPool {
	pool := WorkPool{}
	job := func(job interface{}) interface{} {
		if method, ok := job.(func()); ok {
			method()
			return nil
		}
		return ErrJobNotFunc
	}
	pool.pool = make(chan *WorkerWrapper, numWorkers)
	pool.workers = make([]*WorkerWrapper, numWorkers)
	for i := range pool.workers {
		newWorker := WorkerWrapper{
			Worker: &(DefaultWorker{&job}),
			pool:   pool.pool}
		pool.workers[i] = &newWorker
		pool.workers[i].open()
	}
	return &pool
}

// Creates a pool for an array of custom workers. The custom workers
//must implement Worker, and may also optionally implement ExtendedWorker.
func CreateCustomPool(customWorkers []Worker) *WorkPool {
	pool := WorkPool{}
	pool.pool = make(chan *WorkerWrapper, len(customWorkers))
	pool.workers = make([]*WorkerWrapper, len(customWorkers))
	for i := range pool.workers {
		newWorker := WorkerWrapper{
			Worker: customWorkers[i],
			pool:   pool.pool}
		pool.workers[i] = &newWorker
		pool.workers[i].open()
	}
	return &pool
}
