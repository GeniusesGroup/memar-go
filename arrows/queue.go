//Copyright 2017 SabzCity
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

//
type Queue struct {
	//Add jobs to buffer and dispatcher get jobs from buffer.
	buffer chan interface{}
	//pool of workers.
	pool chan chan interface{}
}

//Send a job to buffer.
//This method don't wait for any workers.
func (queue *Queue) SendJob(job interface{}) {
	queue.buffer <- job
}

//dispatch jobs to workers.
func (queue *Queue) dispatch() {
	for {
		job := <-queue.buffer
		go func(j interface{}) {
			worker := <-queue.pool
			worker <- j
		}(job)
	}
}

//Create a queue and run workers in goroutines.
func CreateQueue(numWorkers int, buffersize int, action func(interface{})) *Queue {
	queue := Queue{
		buffer: make(chan interface{}, buffersize),
		pool:   make(chan chan interface{}, numWorkers)}
	go queue.dispatch()
	for i := 0; i < numWorkers; i++ {
		go func(pool chan chan interface{}, action *func(interface{})) {
			jobchannel := make(chan interface{})
			for {
				//register itself in pool.
				pool <- jobchannel
				//wait for job and do it.
				job := <-jobchannel
				(*action)(job)
			}
		}(queue.pool, &action)
	}
	return &queue
}
