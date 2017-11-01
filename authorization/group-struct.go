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

package authorization

// Group : base structure for group
type Group struct {
	Details       Details       `json:"GroupDetails" xml:"GroupDetails" valid:"required"`
	AccessControl AccessControl `json:"AccessControl" xml:"AccessControl" valid:"required"` //must store arrays in sort for easy read and comparison
}

// Details :
type Details struct {
	Name        string `json:"name" xml:"name" valid:"required"`
	JoinPolicy  bool   `json:"JoinPolicy" xml:"JoinPolicy"` //true:just add by org admin group, false:join by UUID
	Description string `json:"Description" xml:"Description"`
}

// AccessControl :
type AccessControl struct {
	When  []string `json:"When" xml:"When"`   //Iso8601 Time intervals <start>/<end> ["hhmmss/hhmmss", "hhmmss/hhmmss"]	!!!Just use GMT0!!!
	Where []string `json:"Where" xml:"Where"` //["CoordinatesMSID", "CoordinatesMSID"]
	Which []string `json:"Which" xml:"Which"` //["MSName.MethodName", "MSName.MethodName"]
	How   []string `json:"How" xml:"How"`     //
	What  []string `json:"What" xml:"What"`   //["RecordUUID", "RecordUUID"]
	If    []string `json:"If" xml:"If"`       //Have specific header in response from service to user
}
