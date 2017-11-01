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

//You can get a ticket from semaphore and do your worker.
type Semaphore struct {
	//this channel is ticket storage.
	tickets chan struct{}
}

//Begin to work with a ticket.
func (semaphore *Semaphore) Begin() {
	<-semaphore.tickets
}

//End work and revoke the ticket.
func (semaphore *Semaphore) End() {
	semaphore.tickets <- struct{}{}
}

//Create a semaphore
func CreateSemaphore(capacity int) *Semaphore {
	semaphore := Semaphore{
		tickets: make(chan struct{}, capacity),
	}
	//fill ticket storage.
	for i := 0; i < capacity; i++ {
		semaphore.tickets <- struct{}{}
	}
	return &semaphore
}
