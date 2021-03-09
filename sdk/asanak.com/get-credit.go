/* For license and copyright information please see LEGAL file in repository */

package asanak

import (
	"../../achaemenid"
	"../../http"
	"../../log"
)

// GetCreditRes is response structure of GetCredit()
type GetCreditRes struct {
	Credit string
}

// GetCredit send req to asanak and return res to caller! it block caller until finish proccess.
func (a *Asanak) GetCredit() (res *GetCreditRes, err error) {
	return a.getCreditByHTTP()
}

// https://panel.asanak.com/webservice/v1rest/getcredit?username=$USERNAME&password=$PASSWORD
func (a *Asanak) getCreditByHTTP() (res *GetCreditRes, err error) {
	var httpReq = http.MakeNewRequest()
	httpReq.Method = http.MethodPOST
	httpReq.URI.Path = sendSMSPath
	httpReq.Version = http.VersionHTTP11

	httpReq.Header.Set(http.HeaderKeyHost, domain)
	httpReq.Header.Set(http.HeaderKeyAcceptContent, "application/json")
	httpReq.Header.Set(http.HeaderKeyContentType, "application/x-www-form-urlencoded")

	// make a slice by needed size to avoid copy data on append!
	httpReq.Body = make([]byte, 0, sendSMSFixedSize+a.fixSizeDatalen)
	// set username
	httpReq.Body = append(httpReq.Body, username...)
	httpReq.Body = append(httpReq.Body, a.Username...)
	// set password
	httpReq.Body = append(httpReq.Body, password...)
	httpReq.Body = append(httpReq.Body, a.Password...)

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
		log.Fatal("Authorization failed with asanak services! Double check account username & password")
	case http.StatusInternalServerErrorCode:
		return nil, achaemenid.ErrInternalError
	}

	// TODO::: decode body by???
	res = &GetCreditRes{
		Credit: string(httpRes.Body),
	}

	return
}

func (a *Asanak) getCreditBySRPC() (res *GetCreditRes, err error) {
	// TODO::: when asanak.com impelement SRPC, complete SDK here!
	return
}
