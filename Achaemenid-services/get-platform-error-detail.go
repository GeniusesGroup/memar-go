/* For license and copyright information please see LEGAL file in repository */

package ss

import (
	"../achaemenid"
	errorr "../error"
	"../http"
	"../json"
	lang "../language"
	"../syllab"
)

var getPlatformErrorDetailService = achaemenid.Service{
	ID:                2240633614,
	IssueDate:         1601109379,
	ExpiryDate:        0,
	ExpireInFavorOf:   "", // English name of favor service just to show off!
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.EnglishLanguage: "GetPlatformErrorDetail",
	},
	Description: map[lang.Language]string{
		lang.EnglishLanguage: "Returns error all details by given ID",
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: GetPlatformErrorDetailSRPC,
	HTTPHandler: GetPlatformErrorDetailHTTP,
}

// GetPlatformErrorDetailSRPC is sRPC handler of GetPlatformErrorDetail service.
func GetPlatformErrorDetailSRPC(s *achaemenid.Server, st *achaemenid.Stream) {
	var req = &getPlatformErrorDetailReq{}
	st.ReqRes.Err = req.syllabDecoder(st.Payload[4:])
	if st.ReqRes.Err != nil {
		return
	}

	var res *getPlatformErrorDetailRes
	res, st.ReqRes.Err = getPlatformErrorDetail(st, req)
	// Check if any error occur in bussiness logic
	if st.ReqRes.Err != nil {
		return
	}

	st.ReqRes.Payload = res.Syllab
}

// GetPlatformErrorDetailHTTP is HTTP handler of GetPlatformErrorDetail service.
func GetPlatformErrorDetailHTTP(s *achaemenid.Server, st *achaemenid.Stream, httpReq *http.Request, httpRes *http.Response) {
	var req = &getPlatformErrorDetailReq{}
	st.ReqRes.Err = req.jsonDecoder(httpReq.Body)
	if st.ReqRes.Err != nil {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		return
	}

	var res *getPlatformErrorDetailRes
	res, st.ReqRes.Err = getPlatformErrorDetail(st, req)
	// Check if any error occur in bussiness logic
	if st.ReqRes.Err != nil {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		return
	}

	httpRes.SetStatus(http.StatusOKCode, http.StatusOKPhrase)
	httpRes.Header.Set(http.HeaderKeyContentType, "application/json")
	httpRes.Body = res.JSON
}

type getPlatformErrorDetailReq struct {
	ErrorID uint32
}

type getPlatformErrorDetailRes struct {
	JSON   []byte
	Syllab []byte
}

func getPlatformErrorDetail(st *achaemenid.Stream, req *getPlatformErrorDetailReq) (res *getPlatformErrorDetailRes, err error) {
	var er = errorr.GetErrByCode(req.ErrorID)
	res = &getPlatformErrorDetailRes{
		JSON:   er.JSON,
		Syllab: er.Syllab,
	}
	return
}

func (req *getPlatformErrorDetailReq) syllabDecoder(buf []byte) (err error) {
	// TODO::: Use syllab generator to have better performance!
	err = syllab.UnMarshal(buf, req)
	return
}

func (req *getPlatformErrorDetailReq) jsonDecoder(buf []byte) (err error) {
	// TODO::: Use json generator to have better performance!
	err = json.UnMarshal(buf, req)
	return
}
