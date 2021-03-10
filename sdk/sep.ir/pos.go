/* For license and copyright information please see LEGAL file in repository */

package sep

import (
	"../../achaemenid"
	er "../../error"
	"../../http"
	"../../json"
	"../../log"
)

const (
	posDomain     = "cpcpos.seppay.ir"
	posDomainPort = "cpcpos.seppay.ir:443"

	posIP   = "91.240.180.189" // must set on POS devices
	posPort = "6000"           // must set on POS devices

	posPathReceiveIdentifier = "/v1/PcPosTransaction/ReciveIdentifier"
	posPathSendOrder         = "/v1/PcPosTransaction/StartPayment"
	posPathCheckOrder        = "/v1/PcPosTransaction/Inquery"
)

// POS store data to send and receive request to posDomain!
type POS struct {
	Username    string
	Password    string
	TerminalIDs map[string]bool
	IDN         idn `json:"-"`
}

// Init use to initialize SEP.ir POS pointer by given server data.
// Don't pass nil POS otherwise panic will occur.
// Place "sep.ir-pos.json" file in "secret" folder in top of repository
func (pos *POS) Init(jsonSettings []byte) (err *er.Error) {
	if jsonSettings == nil {
		return ErrBadPOSSettings
	}

	pos.jsonDecoder(jsonSettings)
	if pos.Username == "" || pos.Password == "" || len(pos.TerminalIDs) == 0 {
		return ErrBadPOSSettings
	}

	// TODO::: Make GP connection to sep.ir, when sep.ir impelement GP!!
	return
}

func (pos *POS) jsonDecoder(buf []byte) (err *er.Error) {
	// TODO::: Change to generated code!
	err = json.UnMarshal(buf, pos)
	return
}

func (pos *POS) checkTerminalID(TerminalID string) (err *er.Error) {
	var active = pos.TerminalIDs[TerminalID]
	if !active {
		return ErrBadTerminalID
	}
	return
}

func (pos *POS) doHTTP(httpReq *http.Request) (httpRes *http.Response, err *er.Error) {
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

	var res = make([]byte, 2048)
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
		log.DeepDebug("sep.ir - Send msg to:::", httpReq.URI.Raw, httpReq.Header, string(httpReq.Body))
		log.DeepDebug("sep.ir - Received msg from:::", httpRes.ReasonPhrase, httpRes.Header, string(httpRes.Body))
	}
	return
}

func (pos *POS) sendSRPCRequest(req []byte) (res []byte, err *er.Error) {
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
