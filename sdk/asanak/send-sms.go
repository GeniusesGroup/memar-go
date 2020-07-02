/* For license and copyright information please see LEGAL file in repository */

package asanak

import (
	"../../achaemenid"
	"../../http"
	"../../json"
)

// SendSMSReq is request structure of SendSMS()
type SendSMSReq struct {
	Destination []string // max 100 number
	Message     string
}

// SendSMSRes is response structure of SendSMS()
type SendSMSRes []uint64 // max 100 number

// SendSMS send req to asanak and return res to caller! it block caller until finish proccess.
// https://panel.asanak.com/webservice/v1rest/sendsms?username=$USERNAME&password=$PASSWORD&
// source=$SOURCE&destination=$DESTINATION&message=$MESSAGE
func (a *Asanak) SendSMS(req *SendSMSReq) (res SendSMSRes, err error) {
	// Due to legal paper not allow to send more than 100 destination at a request!
	var ln = len(req.Destination)
	if ln == 0 || ln > 99 {
		panic("Destination number can't be 0 or more than 100")
	}
	// Check if no connection exist to use!
	if a.conn == nil {
		return nil, ErrNoConnectionToAsanak
	}

	var httpReq = http.MakeNewRequest()
	httpReq.Method = http.MethodPOST
	httpReq.URI.Authority = domain
	httpReq.URI.Path = sendSMSPath
	httpReq.Version = http.VersionHTTP11

	httpReq.Header.SetValue(http.HeaderKeyAcceptContent, "application/json")
	httpReq.Header.SetValue(http.HeaderKeyContentType, "application/x-www-form-urlencoded")
	httpReq.Header.SetValue(http.HeaderKeyCacheControl, "no-cache")

	var desLn = len(req.Destination)
	// make a slice by needed size to avoid copy data on append! '+desLN' use for destination seprators(',')!
	httpReq.Body = make([]byte, 0, sendSMSFixedSize+a.fixSizeDatalen+len(req.Message)+desLn+(desLn*maxNumberLength))
	// set username
	httpReq.Body = append(httpReq.Body, username...)
	httpReq.Body = append(httpReq.Body, a.UserName...)
	// set password
	httpReq.Body = append(httpReq.Body, password...)
	httpReq.Body = append(httpReq.Body, a.Password...)
	// set source number
	httpReq.Body = append(httpReq.Body, source...)
	httpReq.Body = append(httpReq.Body, a.SourceNumber...)
	// set destination numbers
	httpReq.Body = append(httpReq.Body, destination...)
	for i := 0; i < desLn; i++ {
		httpReq.Body = append(httpReq.Body, req.Destination[i]...)
		httpReq.Body = append(httpReq.Body, ',') // "," || "-"
	}
	// set message text
	httpReq.Body = append(httpReq.Body, message...)
	httpReq.Body = append(httpReq.Body, req.Message...)

	// Make new request-response streams
	var reqStream, resStream *achaemenid.Stream
	reqStream, resStream, err = a.conn.MakeBidirectionalStream(0)

	reqStream.Payload = httpReq.Marshal()
	// Block and wait for response
	err = achaemenid.HTTPOutcomeRequestHandler(a.server, reqStream)
	if err != nil {
		return nil, err
	}

	var httpRes = http.MakeNewResponse()
	err = httpRes.UnMarshal(resStream.Payload)
	if err == nil {
		return nil, err
	}

	switch httpRes.StatusCode {
	case http.StatusBadRequestCode:
		return nil, ErrBadRequest
	case http.StatusUnauthorizedCode:
		panic("Authorization failed with asanak services! Double check account username & password")
	case http.StatusInternalServerErrorCode:
		return nil, ErrAsanakServerInternalError
	}

	res = SendSMSRes{}
	err = res.jsonDecoder(httpRes.Body)

	return
}

func (res *SendSMSRes) jsonDecoder(buf []byte) (err error) {
	// TODO::: Help to complete json generator package to have better performance!
	err = json.UnMarshal(buf, res)
	return
}
