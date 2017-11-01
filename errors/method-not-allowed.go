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

//405 Method Not Allowed - X-Error Related

package errors

import (
	"net/http"
)

//Declare Errors SabzCity Code
const (
	methodNotAllowed = 40500 + (iota + 1)

	methodExpaierd

	contentAlreadyExist

	methodAlreadyDone
)

//Declare Errors Detials
var (
	MethodNotAllowed = New("There is not handler for this Method on given URL. Use OPTIONS method for extra info", methodNotAllowed, http.StatusMethodNotAllowed)

	MethodExpaierd = New("Method was expired", methodExpaierd, http.StatusMethodNotAllowed)

	ContentAlreadyExist = New("This content was already exist", contentAlreadyExist, http.StatusMethodNotAllowed)

	MethodAlreadyDone = New("This method was already done", methodAlreadyDone, http.StatusMethodNotAllowed)
)
