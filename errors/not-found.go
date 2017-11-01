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

// 404 Not Found - X-Error Related

package errors

import (
	"net/http"
)

//Declare Errors SabzCity Code
const (
	userNotFound = 40400 + (iota + 1)

	routerNotFound

	contentNotExist

	recordNotFound

	domainNotRegistered

	domainVersionNotFound

	unitNotFound

	lesseeNotFound

	complexNotFound

	groupNotFound

	userNotJoinedInAnyGroup

	keyNotFound
)

//Declare Errors Detials
var (
	UserNotFound = New("The requested user was not found", userNotFound, http.StatusNotFound)

	RouterNotFound = New("The requested url(rest) or methodName(RPC) not match with any register router in requested service", routerNotFound, http.StatusNotFound)

	ContentNotExist = New("Requested content was not exist", contentNotExist, http.StatusNotFound)

	RecordNotFound = New("User or any special record not found", recordNotFound, http.StatusNotFound)

	DomainNotRegistered = New("Requested domain not registered", domainNotRegistered, http.StatusNotFound)

	DomainVersionNotFound = New("Requested version of domain not found", domainVersionNotFound, http.StatusNotFound)

	UnitNotFound = New("The requested unit was not found", unitNotFound, http.StatusNotFound)

	LesseeNotFound = New("The requested lessee was not found", lesseeNotFound, http.StatusNotFound)

	ComplexNotFound = New("The requested complex was not found", complexNotFound, http.StatusNotFound)

	GroupNotFound = New("The requested group was not found", groupNotFound, http.StatusNotFound)

	UserNotJoinedInAnyGroup = New("User is not joined in any group", userNotJoinedInAnyGroup, http.StatusNotFound)

	KeyNotFound = New("Requested key or directory not found", keyNotFound, http.StatusNotFound)
)
