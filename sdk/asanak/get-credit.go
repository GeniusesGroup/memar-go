/* For license and copyright information please see LEGAL file in repository */

package asanak

import (
	"../../achaemenid"
	"../../http"
)

// GetCreditRes is response structure of GetCredit()
type GetCreditRes struct {
	Credit string
}

// GetCredit send req to asanak and return res to caller! it block caller until finish proccess.
// https://panel.asanak.com/webservice/v1rest/getcredit?username=$USERNAME&password=$PASSWORD
func (a *Asanak) GetCredit() (res *GetCreditRes, err error) {
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

	// make a slice by needed size to avoid copy data on append!
	httpReq.Body = make([]byte, 0, sendSMSFixedSize+a.fixSizeDatalen)
	// set username
	httpReq.Body = append(httpReq.Body, username...)
	httpReq.Body = append(httpReq.Body, a.UserName...)
	// set password
	httpReq.Body = append(httpReq.Body, password...)
	httpReq.Body = append(httpReq.Body, a.Password...)

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

	// TODO::: decode body by???
	res = &GetCreditRes{
		Credit: string(httpRes.Body),
	}

	return
}
