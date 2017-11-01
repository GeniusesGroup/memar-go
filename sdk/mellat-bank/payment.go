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

package mellatbank

import "github.com/SabzCity/go-library/sdk"

const (
	// RequestServer : SOAP server for request transaction.
	RequestServer = "https://pgws.bpm.bankmellat.ir/pgwchannel/services/pgw"

	// PaymentServer : This server do payment for user.
	PaymentServer = "https://pgw.bpm.bankmellat.ir/pgwchannel/illegaloperation.mellat"
)

// NewPayment : Create mellat payment and convert it to interface.
func NewPayment(username, password string, terminalID uint) sdk.Payment {

	return &Payment{Username: username, Password: password, TerminalID: terminalID}
}

// Payment : Implement of payment for mellat bank.
type Payment struct {
	Username    string
	Password    string
	CallBackURL string
	TerminalID  uint
}

// PayRequest : Request a new transaction ID.
func (payment *Payment) PayRequest(orderID uint, amount uint, description string, payerID uint) (string, error) {

	return "", nil
}
