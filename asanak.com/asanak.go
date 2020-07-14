/* For license and copyright information please see LEGAL file in repository */

package asanak

import (
	"../achaemenid"
	"../errors"
	"../json"
	"../log"
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
	gpConn         *achaemenid.Connection
}

// Init use to initialize Asanak pointer by given server data.
// Don't pass nil Asanak otherwise panic will occur.
// Place "asanak.com.json" file in "secret" folder in top of repository
func (a *Asanak) Init(jsonSettings []byte) {
	a.jsonDecoder(jsonSettings)
	if a.UserName == "" || a.Password == "" || a.SourceNumber == "" {
		log.Fatal("'asanak.com.json' file in 'secret' folder has not enough information")
	}

	a.RecChan = make(chan *GetSmsByAsanak)
	a.fixSizeDatalen = len(a.UserName) + len(a.Password) + len(a.SourceNumber)

	// TODO::: Make GP connection to asanak.com, when asanak.com impelement GP!!
}

/* decode json data in this format:
{
    "UserName": "",
    "Password": "",
    "SourceNumber": ""
}
*/
func (a *Asanak) jsonDecoder(buf []byte) (err error) {
	// TODO::: Change to generated code!
	err = json.UnMarshal(buf, a)
	return
}

func (a *Asanak) sendHTTPRequest(req []byte) (res []byte, err error) {
	var tlsConn *tlsConn
	tlsConn, err = getTLSConn()
	if err != nil {
		return nil, ErrNoConnectionToAsanak
	}

	_, err = tlsConn.conn.Write(req)
	if err != nil {
		return
	}

	res = make([]byte, 4096)
	var readSize int
	readSize, err = tlsConn.conn.Read(res)
	if err != nil {
		return
	}

	tlsConn.free()

	return res[:readSize], nil
}

func (a *Asanak) sendSRPCRequest(req []byte) (res []byte, err error) {
	// Make new request-response streams
	// var reqStream, resStream *achaemenid.Stream
	// reqStream, resStream, err = a.conn.MakeBidirectionalStream(0)

	// Block and wait for response
	// err = achaemenid.HTTPOutcomeRequestHandler(a.server, reqStream)
	// if err != nil {
	// 	return nil, err
	// }

	return
}
