/* For license and copyright information please see LEGAL file in repository */

package asanak

import (
	"../../achaemenid"
	"../../errors"
	"../../json"
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

// Errors
var (
	ErrNoConnectionToAsanak      = errors.New("NoConnectionToAsanak", "No connection exist to Asanak servers due to temporary or long term problem")
	ErrBadRequest                = errors.New("BadRequest", "Some given data in request must be invalid or server not accept them")
	ErrAsanakServerInternalError = errors.New("AsanakServerInternalError", "Asanak server encounter problem due to temporary or long term problem!")
)

// Asanak store data to send and receive sms by asanak provider!
type Asanak struct {
	UserName       string
	Password       string
	SourceNumber   string
	RecChan        chan *GetSmsByAsanak
	fixSizeDatalen int
	server         *achaemenid.Server
	conn           *achaemenid.Connection
}

// Init use to initialize Asanak pointer by given server data.
// Don't pass nil Asanak otherwise panic will occur.
// Place "asanak.com.json" || "asanak.com.syllab" file in "secret" folder in top of repository
func (a *Asanak) Init(s *achaemenid.Server) {
	var f = s.Assets.Secret.GetFile("asanak.com.syllab")
	if f == nil {
		f = s.Assets.Secret.GetFile("asanak.com.json")
	}
	if f != nil {
		a.jsonDecoder(f.Data)
	} else {
		panic("Can't find 'asanak.com.json' or 'asanak.com.syllab' file in 'secret' folder in top of repository")
	}

	a.RecChan = make(chan *GetSmsByAsanak)
	a.fixSizeDatalen = len(a.UserName) + len(a.Password) + len(a.SourceNumber)
	a.server = s
	// make connection to asanak server by domain!
}

/* decode json data in this format:
{
    "UserName": "",
    "Password": "",
    "SourceNumber": ""
}
*/
func (a *Asanak) jsonDecoder(buf []byte) (err error) {
	// TODO::: Change to generator code for decode!
	err = json.UnMarshal(buf, a)
	return
}
