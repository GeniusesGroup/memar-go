/* For license and copyright information please see LEGAL file in repository */

package asanak

import (
	"../http"
	"../json"
	"../log"
)

// SendSMSReq is request structure of SendSMS()
type SendSMSReq struct {
	Destination []string // max 100 number
	Message     string
}

// SendSMSRes is response structure of SendSMS()
type SendSMSRes []uint64 // max 100 number

// SendSMS send req to asanak and return res to caller! it block caller until finish proccess.
func (a *Asanak) SendSMS(req *SendSMSReq) (res SendSMSRes, err error) {
	// Due to legal paper not allow to send more than 100 destination at a request!
	var ln = len(req.Destination)
	if ln == 0 || ln > 99 {
		log.Fatal("Destination number can't be 0 or more than 100")
	}

	return a.sendSMSByHTTP(req)
}

// https://panel.asanak.com/webservice/v1rest/sendsms?username=$USERNAME&password=$PASSWORD&
// source=$SOURCE&destination=$DESTINATION&message=$MESSAGE
func (a *Asanak) sendSMSByHTTP(req *SendSMSReq) (res SendSMSRes, err error) {
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

	var serverRes []byte
	serverRes, err = a.sendHTTPRequest(httpReq.Marshal())

	var httpRes = http.MakeNewResponse()
	err = httpRes.UnMarshal(serverRes)
	if err == nil {
		return nil, err
	}

	switch httpRes.StatusCode {
	case http.StatusBadRequestCode:
		return nil, ErrBadRequest
	case http.StatusUnauthorizedCode:
		log.Fatal("Authorization failed with asanak services! Double check account username & password")
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

func (a *Asanak) sendSMSBySRPC(req *SendSMSReq) (res SendSMSRes, err error) {
	// TODO::: when asanak.com impelement SRPC, complete SDK here!
	return
}
