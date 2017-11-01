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

// 403 Forbidden - X-Error Related

package errors

import (
	"net/http"
)

//Declare Errors SabzCity Code
const (
	accessForbidden = 40300 + (iota + 1)

	codeIncorrect

	rootNotEditable

	usernameLocked

	guestGroupNotJoinable

	groupNotJoinable

	groupNotDeletable

	userNotAcceptApp

	invoiceNotEditable
)

//Declare Errors Detials
var (
	AccessForbidden = New("User can't access to the resource or Change it", accessForbidden, http.StatusForbidden)

	CodeIncorrect = New("The code you entered is incorrect", codeIncorrect, http.StatusForbidden)

	RootNotEditable = New("You can't change or delete root variable", rootNotEditable, http.StatusForbidden)

	UsernameLocked = New("You can't change username. Wait for one month", usernameLocked, http.StatusForbidden)

	GuestGroupNotJoinable = New("No one can join registered or guest group", guestGroupNotJoinable, http.StatusForbidden)

	GroupNotJoinable = New("Group has admin only policy. Ask organization admins to add you to related group", groupNotJoinable, http.StatusForbidden)

	GroupNotDeletable = New("Admin & Registered & Guest Groups is not deletable", groupNotJoinable, http.StatusForbidden)

	UserNotAcceptApp = New("User not accept or remove app permissions, Ask again permission from user", userNotAcceptApp, http.StatusForbidden)

	InvoiceNotEditable = New("User & Org can edit or delete invoice just before Checkout.", invoiceNotEditable, http.StatusForbidden)
)
