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

// 500 Internal Server Error - X-Error Related

package errors

import (
	"net/http"
)

//Declare Errors SabzCity Code
const (
	cantPrepareStatement = 50000 + (iota + 1)

	serverPanic

	storingDataConflict

	storingDataNotComplete

	someThingIsWrong

	contentMustBeStruct

	validatorNotFound

	tlsCertificateNotExist
)

//Declare Errors Detials
var (
	CantPrepareStatement = New("Can't prepare a new statement to database", cantPrepareStatement, http.StatusInternalServerError)

	ServerPanic = New("Server recover from panic situation. Error send to admins, contact admins for fixes time.", serverPanic, http.StatusInternalServerError)

	StoringDataConflict = New("Your data has some problem in their structure to store in our databases. Send your request again! If error exist contact SabzCity platform administrators", storingDataConflict, http.StatusInternalServerError)

	StoringDataNotComplete = New("We have some problem in storing your data in our databases. Send your request again! If error exist contact SabzCity platform administrators", storingDataNotComplete, http.StatusInternalServerError)

	SomeThingIsWrong = New("The error that unpossible for us", someThingIsWrong, http.StatusInternalServerError)

	ContentMustBeStruct = New("The Content type isn't struct", contentMustBeStruct, http.StatusInternalServerError)

	ValidatorNotFound = New("Some validator function for validation package not found", validatorNotFound, http.StatusInternalServerError)

	TLSCertificateNotExist = New("We can't make connection with this host! No related certificate found", tlsCertificateNotExist, http.StatusInternalServerError)
)
