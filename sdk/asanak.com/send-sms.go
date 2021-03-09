/* For license and copyright information please see LEGAL file in repository */

package asanak

import (
	"../../achaemenid"
	er "../../error"
	"../../http"
	"../../json"
	"../../log"
)

// SendSMSReq is request structure of SendSMS()
type SendSMSReq struct {
	Destination []string // max 100 number
	Message     []byte
}

// SendSMSRes is response structure of SendSMS()
type SendSMSRes []uint64 // max 100 number

// SendSMS send req to asanak and return res to caller! it block caller until finish proccess.
func (a *Asanak) SendSMS(req *SendSMSReq) (res SendSMSRes, err *er.Error) {
	// Due to legal paper not allow to send more than 100 destination at a request!
	var ln = len(req.Destination)
	if ln > 99 {
		log.Warn("Asanak.com - Destination number can't be more than 100!")
		req.Destination = req.Destination[:100]
	}

	return a.sendSMSByHTTP(req)
}

// https://panel.asanak.com/webservice/v1rest/sendsms?username=$USERNAME&password=$PASSWORD&
// source=$SOURCE&destination=$DESTINATION&message=$MESSAGE
func (a *Asanak) sendSMSByHTTP(req *SendSMSReq) (res SendSMSRes, err *er.Error) {
	var httpReq = http.MakeNewRequest()
	httpReq.Method = http.MethodPOST
	httpReq.URI.Path = sendSMSPath
	httpReq.Version = http.VersionHTTP11

	httpReq.Header.Set(http.HeaderKeyHost, domain)
	httpReq.Header.Set(http.HeaderKeyAcceptContent, "application/json")
	httpReq.Header.Set(http.HeaderKeyContentType, "application/x-www-form-urlencoded")

	httpReq.Body = req.formEncoder(a)

	// Set some other header data
	httpReq.SetContentLength()
	httpReq.Header.Set(http.HeaderKeyServer, http.DefaultServer)

	var httpRes *http.Response
	httpRes, err = a.doHTTP(httpReq)
	if err != nil {
		return
	}

	switch httpRes.StatusCode {
	case http.StatusBadRequestCode:
		return nil, achaemenid.ErrBadRequest
	case http.StatusUnauthorizedCode:
		log.Warn("Authorization failed with asanak services! Double check account username & password")
	case http.StatusInternalServerErrorCode:
		return nil, achaemenid.ErrInternalError
	}

	res = SendSMSRes{}
	err = res.jsonDecoder(httpRes.Body)
	if err != nil {
		return nil, achaemenid.ErrBadResponse
	}

	return
}

func (req *SendSMSReq) formEncoder(a *Asanak) (buf []byte) {
	var desLn = len(req.Destination)
	// make a slice by needed size to avoid copy data on append! '+desLN' use for destination seprators(',')!
	buf = make([]byte, 0, sendSMSFixedSize+a.fixSizeDatalen+len(req.Message)+desLn+(desLn*maxNumberLength))
	// set username
	buf = append(buf, username...)
	buf = append(buf, a.Username...)
	// set password
	buf = append(buf, password...)
	buf = append(buf, a.Password...)
	// set source number
	buf = append(buf, source...)
	buf = append(buf, a.SourceNumber...)
	// set destination numbers
	buf = append(buf, destination...)
	for i := 0; i < desLn; i++ {
		buf = append(buf, req.Destination[i]...)
		buf = append(buf, ',') // "," || "-"
	}
	buf = buf[:len(buf)-1] // Remove trailing comma
	// set message text
	buf = append(buf, message...)
	buf = append(buf, req.Message...)
	return
}

func (res *SendSMSRes) jsonDecoder(buf []byte) (err *er.Error) {
	var decoder = json.DecoderUnsafeMinifed{
		Buf: buf,
	}
	for len(decoder.Buf) > 2 {
		decoder.Offset(1)
		*res, err = decoder.DecodeUInt64SliceAsNumber()
	}
	return
}

func (a *Asanak) sendSMSBySRPC(req *SendSMSReq) (res SendSMSRes, err error) {
	// TODO::: when asanak.com impelement SRPC, complete SDK here!
	return
}
