/* For license and copyright information please see LEGAL file in repository */

package asanak

import (
	"../../achaemenid"
)

// https://asanak.com/media/2019/02/Asanak_REST_Webservice_Documentation_v3.0.pdf

const (
	// maximum length of 2(world-code)+15(local) digits to telephone numbers
	// CC-xx-xxxxxxxxxxxxx
	maxNumberLength = 17

	domain = "panel.asanak.com"

	sendSMSPath      = "/webservice/v1rest/sendsms"
	sendSMSFixedSize = len(username) + len(password) + len(source) + len(destination) + len(message)

	getCreditPath      = "/webservice/v1rest/getcredit"
	getCreditFixedSize = len(username) + len(password)

	msgStatusPath      = "/webservice/v1rest/msgstatus"
	msgStatusFixedSize = len(username) + len(password) + len(msgID)

	username    = "username="
	password    = "&password="
	source      = "&source="
	destination = "&destination="
	message     = "&message="
	msgID       = "&msgid="
)

// Asanak store data to send and receive sms by asanak provider!
type Asanak struct {
	UserName     string
	Password     string
	SourceNumber string
	RecChan      chan *GetSmsByAsanak
	len          int
	server       *achaemenid.Server
	conn         *achaemenid.Connection
}

// Init use to initialize Asanak pointer by given data
func (a *Asanak) Init(s *achaemenid.Server, userName, password, sourceNumber string) {
	a = &Asanak{
		UserName:     userName,
		Password:     password,
		SourceNumber: sourceNumber,
		RecChan:      make(chan *GetSmsByAsanak),
		len:          len(userName) + len(password) + len(sourceNumber),
		server:       s,
	}
	// make connection to asanak server!
}
