/* For license and copyright information please see LEGAL file in repository */

package asanak

import (
	"../../achaemenid"
	er "../../error"
	"../../http"
	"../../json"
	"../../log"
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
	Username       string
	Password       string
	SourceNumber   string
	RecChan        chan *GetSmsByAsanak   `json:"-"`
	fixSizeDatalen int                    `json:"-"`
	gpConn         *achaemenid.Connection `json:"-"`
}

// Init use to initialize Asanak pointer by given server data.
// Don't pass nil Asanak otherwise panic will occur.
// Place "asanak.com.json" file in "secret" folder in top of repository
func (a *Asanak) Init(jsonSettings []byte) {
	a.jsonDecoder(jsonSettings)
	if a.Username == "" || a.Password == "" || a.SourceNumber == "" {
		log.Fatal("'asanak.com.json' file in 'secret' folder has not enough information")
	}

	a.RecChan = make(chan *GetSmsByAsanak)
	a.fixSizeDatalen = len(a.Username) + len(a.Password) + len(a.SourceNumber)

	// TODO::: Make GP connection to asanak.com, when asanak.com impelement GP!!
}

func (a *Asanak) jsonDecoder(buf []byte) (err *er.Error) {
	// TODO::: Change to generated code!
	err = json.UnMarshal(buf, a)
	return
}

func (a *Asanak) doHTTP(httpReq *http.Request) (httpRes *http.Response, err *er.Error) {
	var goErr error
	var tlsConn *tlsConn
	tlsConn, goErr = getTLSConn()
	if goErr != nil || tlsConn == nil {
		return nil, achaemenid.ErrNoConnection
	}

	var req []byte = httpReq.Marshal()
	_, goErr = tlsConn.conn.Write(req)
	if goErr != nil {
		return nil, achaemenid.ErrSendRequest
	}

	var res = make([]byte, 512)
	var readSize int
	readSize, goErr = tlsConn.conn.Read(res)
	if err != nil {
		return nil, achaemenid.ErrReceiveResponse
	}

	httpRes = http.MakeNewResponse()
	err = httpRes.UnMarshal(res)
	if err != nil {
		return
	}

	var contentLength = httpRes.Header.GetContentLength()
	// TODO::: check content length not
	if contentLength > 0 && len(httpRes.Body) == 0 {
		httpRes.Body = make([]byte, contentLength)
		readSize, goErr = tlsConn.conn.Read(httpRes.Body)
		if goErr != nil {
			return
		}
		if readSize != int(contentLength) {
			// err =
			// return
		}
	}

	tlsConn.free()

	if log.DeepDebugMode {
		log.DeepDebug("Asanak.com - Send msg to server:::", httpReq.URI.Raw, httpReq.Header, string(httpReq.Body))
		log.DeepDebug("Asanak.com - Received msg from server:::", httpRes.ReasonPhrase, httpRes.Header, string(httpRes.Body))
	}
	return
}

func (a *Asanak) sendSRPCRequest(req []byte) (res []byte, err *er.Error) {
	// Make new request-response streams
	// var st *achaemenid.Stream
	// st, err = a.conn.MakeBidirectionalStream(0)

	// Block and wait for response
	// err = achaemenid.HTTPOutcomeRequestHandler(a.server, st)
	// if err != nil {
	// 	return nil, err
	// }

	return
}
