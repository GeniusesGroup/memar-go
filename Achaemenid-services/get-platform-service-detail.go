/* For license and copyright information please see LEGAL file in repository */

package ss

import (
	"../achaemenid"
	"../http"
	"../json"
	lang "../language"
	"../srpc"
	"../syllab"
)

var getPlatformServiceDetailService = achaemenid.Service{
	ID:                3799388884,
	IssueDate:         1601109351,
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // English name of favor service just to show off!
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.EnglishLanguage: "GetPlatformServiceDetail",
	},
	Description: map[lang.Language]string{
		lang.EnglishLanguage: "Returns service all details by given service ID or URI",
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: GetPlatformServiceDetailSRPC,
	HTTPHandler: GetPlatformServiceDetailHTTP,
}

// GetPlatformServiceDetailSRPC is sRPC handler of GetPlatformServiceDetail service.
func GetPlatformServiceDetailSRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	var req = &getPlatformServiceDetailReq{}
	st.ReqRes.Err = req.syllabDecoder(st.Payload[4:])
	if st.ReqRes.Err != nil {
		return
	}

	var res *getPlatformServiceDetailRes
	res, st.ReqRes.Err = getPlatformServiceDetail(st, req)
	// Check if any error occur in bussiness logic
	if st.ReqRes.Err != nil {
		return
	}

	st.ReqRes.Payload = res.Syllab
}

// GetPlatformServiceDetailHTTP is HTTP handler of GetPlatformServiceDetail service.
func GetPlatformServiceDetailHTTP(s *achaemenid.Server, st *achaemenid.Stream, httpReq *http.Request, httpRes *http.Response) {
	var req = &getPlatformServiceDetailReq{}
	st.ReqRes.Err = req.jsonDecoder(httpReq.Body)
	if st.ReqRes.Err != nil {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		return
	}

	var res *getPlatformServiceDetailRes
	res, st.ReqRes.Err = getPlatformServiceDetail(st, req)
	// Check if any error occur in bussiness logic
	if st.ReqRes.Err != nil {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		return
	}

	httpRes.SetStatus(http.StatusOKCode, http.StatusOKPhrase)
	httpRes.Header.Set(http.HeaderKeyContentType, "application/json")
	httpRes.Body = res.JSON
}

type getPlatformServiceDetailReq struct {
	ServiceID  uint32
	ServiceURI string
}

type getPlatformServiceDetailRes struct {
	JSON   []byte
	Syllab []byte
}

func getPlatformServiceDetail(st *achaemenid.Stream, req *getPlatformServiceDetailReq) (res *getPlatformServiceDetailRes, err error) {
	var service *achaemenid.Service
	if req.ServiceID != 0 {
		service = server.Services.GetServiceHandlerByID(req.ServiceID)
	} else {
		service = server.Services.GetServiceHandlerByURI(req.ServiceURI)
	}

	if service == nil {
		err = srpc.ErrSRPCServiceNotFound
	} else {
		res = &getPlatformServiceDetailRes{
			JSON:   service.JSON,
			Syllab: service.Syllab,
		}
	}
	return
}

func (req *getPlatformServiceDetailReq) syllabDecoder(buf []byte) (err error) {
	// TODO::: Use syllab generator to have better performance!
	err = syllab.UnMarshal(buf, req)
	return
}

func (req *getPlatformServiceDetailReq) jsonDecoder(buf []byte) (err error) {
	// TODO::: Use json generator to have better performance!
	err = json.UnMarshal(buf, req)
	return
}
