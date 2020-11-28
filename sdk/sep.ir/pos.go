/* For license and copyright information please see LEGAL file in repository */

package sep

import (
	er "../../error"
	"../../json"
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

	err = pos.IDN.getAccessToken(pos.Username, pos.Password)

	// TODO::: Make GP connection to sep.ir, when sep.ir impelement GP!!
	return
}

func (pos *POS) jsonDecoder(buf []byte) (err *er.Error) {
	// TODO::: Change to generated code!
	err = json.UnMarshal(buf, pos)
	return
}

func (pos *POS) checkTerminalID(TerminalID string) (err *er.Error) {
	active := pos.TerminalIDs[TerminalID]
	if !active {
		return ErrBadTerminalID
	}
	return
}

func (pos *POS) sendHTTPRequest(req []byte) (res []byte, err *er.Error) {
	var goErr error
	var tlsConn *tlsConn
	tlsConn, goErr = getTLSConn()
	if goErr != nil {
		return nil, ErrNoConnection
	}

	_, goErr = tlsConn.conn.Write(req)
	if goErr != nil {
		return nil, ErrInternalError
	}

	res = make([]byte, 4096)
	var readSize int
	readSize, goErr = tlsConn.conn.Read(res)
	if err != nil {
		return nil, ErrInternalError
	}

	tlsConn.free()

	return res[:readSize], nil
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
