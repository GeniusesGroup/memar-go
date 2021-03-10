/* For license and copyright information please see LEGAL file in repository */

package sep

import (
	"../../achaemenid"
	er "../../error"
	"../../http"
	"../../json"
	"../../log"
)

// POSCheckOrderReq is request structure of POSCheckOrder()
type POSCheckOrderReq struct {
	Identifier           string
	RRN                  string
	ResNum               string
	TerminalID           string
	CancelPendingRequest bool
}

// POSCheckOrderRes is response structure of POSCheckOrder()
type POSCheckOrderRes struct {
	IsSuccess        bool
	ErrorCode        int64
	ErrorDescription string

	Identifier          string
	TerminalID          int64
	CreatedOn           string
	ResponseCode        string
	ResponseDescription string
	TraceNumber         string
	ResNum              string
	RRN                 string
	State               int64
	StateDescription    string
	HashData            string
	Amount              string
}

// POSCheckOrder use to check status of order sent earlier by SendOrderToPOS service!
func (pos *POS) POSCheckOrder(req *POSCheckOrderReq) (res *POSCheckOrderRes, err *er.Error) {
	err = pos.checkTerminalID(req.TerminalID)
	if err != nil {
		return
	}

	err = pos.IDN.checkAccessToken(pos.Username, pos.Password)
	if err != nil {
		return
	}

	req.Identifier, err = pos.posGetIdentifier()
	if err != nil {
		return
	}

	var httpReq = http.MakeNewRequest()
	httpReq.Method = http.MethodPOST
	httpReq.URI.Path = posPathCheckOrder
	httpReq.Version = http.VersionHTTP11

	httpReq.Header.Set(http.HeaderKeyHost, posDomain)
	httpReq.Header.Set(http.HeaderKeyAcceptContent, "application/json")
	httpReq.Header.Set(http.HeaderKeyContentType, "application/json")
	httpReq.Header.Set(http.HeaderKeyAuthorization, pos.IDN.TokenType+" "+pos.IDN.AccessToken)

	httpReq.Body = req.jsonEncoder()

	// Set some other header data
	httpReq.SetContentLength()
	httpReq.Header.Set(http.HeaderKeyServer, http.DefaultServer)

	var httpRes *http.Response
	httpRes, err = pos.doHTTP(httpReq)
	if err != nil {
		if err.Equal(achaemenid.ErrReceiveResponse) {
			// TODO::: check order
		}
		return
	}

	switch httpRes.StatusCode {
	case http.StatusBadRequestCode:
		err = achaemenid.ErrBadRequest
		return
	case http.StatusUnauthorizedCode:
		log.Warn("Authorization failed with sep services! Double check account username & password")
	case http.StatusInternalServerErrorCode:
		err = achaemenid.ErrInternalError
		return
	}

	switch httpRes.Header.GetContentType().SubType {
	case http.ContentTypeMimeSubTypeHTML:
		err = ErrPOSInternalError
		return
	case http.ContentTypeMimeSubTypeJSON:
		res = &POSCheckOrderRes{}
		err = res.jsonDecoder(httpRes.Body)
		if err != nil {
			err = achaemenid.ErrBadResponse
			return
		}
		if !res.IsSuccess && res.ErrorCode == 0 {
			err = ErrPOSInternalError
			return
		}
		err = getErrorByResCode(res.ErrorCode)
	}
	return
}

func (req *POSCheckOrderReq) jsonEncoder() (buf []byte) {
	var encoder = json.Encoder{
		Buf: make([]byte, 0, req.jsonLen()),
	}

	encoder.EncodeString(`{"Identifier":"`)
	encoder.EncodeString(req.Identifier)

	encoder.EncodeString(`","RRN":"`)
	encoder.EncodeString(req.RRN)

	encoder.EncodeString(`","ResNum":"`)
	encoder.EncodeString(req.ResNum)

	encoder.EncodeString(`","TerminalID":"`)
	encoder.EncodeString(req.TerminalID)

	encoder.EncodeString(`","CancelPendingRequest":`)
	encoder.EncodeBoolean(req.CancelPendingRequest)

	encoder.EncodeByte('}')
	return encoder.Buf
}

func (req *POSCheckOrderReq) jsonLen() (ln int) {
	ln = len(req.Identifier) + len(req.RRN) + len(req.ResNum) + len(req.TerminalID)
	ln += 83
	return
}

func (res *POSCheckOrderRes) jsonDecoder(buf []byte) (err *er.Error) {
	var decoder = json.DecoderUnsafe{
		Buf: buf,
	}
	for err == nil {
		var keyName = decoder.DecodeKey()
		switch keyName {
		case "":
			return
		case "IsSuccess":
			res.IsSuccess, err = decoder.DecodeBool()
		case "ErrorCode":
			res.ErrorCode, err = decoder.DecodeInt64()
		case "ErrorDescription":
			res.ErrorDescription, err = decoder.DecodeString()

		case "Data":
			continue
		case "Identifier":
			res.Identifier, err = decoder.DecodeString()
		case "TerminalID":
			res.TerminalID, err = decoder.DecodeInt64()
		case "CreatedOn":
			res.CreatedOn, err = decoder.DecodeString()
		case "ResponseCode":
			res.ResponseCode, err = decoder.DecodeString()
		case "ResponseDescription":
			res.ResponseDescription, err = decoder.DecodeString()
		case "TraceNumber":
			res.TraceNumber, err = decoder.DecodeString()
		case "ResNum":
			res.ResNum, err = decoder.DecodeString()
		case "RRN":
			res.RRN, err = decoder.DecodeString()
		case "State":
			res.State, err = decoder.DecodeInt64()
		case "StateDescription":
			res.StateDescription, err = decoder.DecodeString()
		case "HashData":
			res.HashData, err = decoder.DecodeString()
		case "Amount":
			res.Amount, err = decoder.DecodeString()
		default:
			err = decoder.NotFoundKey()
		}

		if len(decoder.Buf) < 3 {
			return
		}
	}
	return
}
