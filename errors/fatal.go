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

// FATAL Error - X-Error Related

package errors

//Declare Errors SabzCity Code
const (
	settingsNotFound = 00000 + (iota + 1)

	settingsNotValid

	settingsPartNotExist

	planPartWrong

	canNotServeOrListen

	inaccessiblePort

	badStatement
)

//Declare Errors Detials
var (
	SettingsNotFound = New("Settings Data ('.json' file) is not exist in proper location", settingsNotFound, 0)

	SettingsNotValid = New("JSON syntax of 'settings.json' file have some problems", settingsNotValid, 0)

	SettingsPartNotExist = New("One of setting part not exist in settings", settingsPartNotExist, 0)

	PlanPartWrong = New("This plan not registered or something in it is wrong", planPartWrong, 0)

	CanNotServeOrListen = New("The service can't listen at the port", canNotServeOrListen, 0)

	InaccessiblePort = New("Can't run the microservice in this port", inaccessiblePort, 0)

	BadStatement = New("The data request is wrong", badStatement, 0)
)
