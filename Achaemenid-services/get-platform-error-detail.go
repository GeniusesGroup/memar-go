/* For license and copyright information please see LEGAL file in repository */

package ss

import (
	"../achaemenid"
	er "../error"
	"../http"
	"../json"
	lang "../language"
	"../srpc"
	"../syllab"
)

var getPlatformErrorDetailService = achaemenid.Service{
	URN:                "urn:giti:achaemenid.libgo:service:get-platform-error-detail",
	Domain:             DomainName,
	ID:                 8746018193450136356,
	IssueDate:          1601109379,
	ExpiryDate:         0,
	ExpireInFavorOfURN: "",
	ExpireInFavorOfID:  0,
	Status:             achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.LanguageEnglish: "Get Platform Error Detail",
	},
	Description: map[lang.Language]string{
		lang.LanguageEnglish: "Returns error all details by given ID",
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: GetPlatformErrorDetailSRPC,
	HTTPHandler: GetPlatformErrorDetailHTTP,
}

// GetPlatformErrorDetailSRPC is sRPC handler of GetPlatformErrorDetail service.
func GetPlatformErrorDetailSRPC(st *achaemenid.Stream) {
	var req = &getPlatformErrorDetailReq{}
	st.Err = req.syllabDecoder(srpc.GetPayload(st.IncomePayload))
	if st.Err != nil {
		return
	}

	var res *getPlatformErrorDetailRes
	res, st.Err = getPlatformErrorDetail(st, req)
	// Check if any error occur in bussiness logic
	if st.Err != nil {
		return
	}

	st.OutcomePayload = res.Syllab
}

// GetPlatformErrorDetailHTTP is HTTP handler of GetPlatformErrorDetail service.
func GetPlatformErrorDetailHTTP(st *achaemenid.Stream, httpReq *http.Request, httpRes *http.Response) {
	var req = &getPlatformErrorDetailReq{}
	st.Err = req.jsonDecoder(httpReq.Body)
	if st.Err != nil {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		return
	}

	var res *getPlatformErrorDetailRes
	res, st.Err = getPlatformErrorDetail(st, req)
	// Check if any error occur in bussiness logic
	if st.Err != nil {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		return
	}

	httpRes.SetStatus(http.StatusOKCode, http.StatusOKPhrase)
	httpRes.Header.Set(http.HeaderKeyContentType, "application/json")
	httpRes.Body = res.JSON
}

type getPlatformErrorDetailReq struct {
	ErrorID uint64
}

type getPlatformErrorDetailRes struct {
	JSON   []byte
	Syllab []byte
}

func getPlatformErrorDetail(st *achaemenid.Stream, req *getPlatformErrorDetailReq) (res *getPlatformErrorDetailRes, err *er.Error) {
	var er = er.Errors.GetErrorByCode(req.ErrorID)
	res = &getPlatformErrorDetailRes{
		JSON:   er.JSON,
		Syllab: er.Syllab,
	}
	return
}

/*
	Request Encoders & Decoders
*/

func (req *getPlatformErrorDetailReq) syllabDecoder(buf []byte) (err *er.Error) {
	if uint32(len(buf)) < req.syllabStackLen() {
		err = syllab.ErrShortSliceDecoded
		return
	}

	req.ErrorID = syllab.GetUInt64(buf, 0)
	return
}

func (req *getPlatformErrorDetailReq) syllabEncoder(buf []byte) {
	syllab.SetUInt64(buf, 0, req.ErrorID)
	return
}

func (req *getPlatformErrorDetailReq) syllabStackLen() (ln uint32) {
	return 8
}

func (req *getPlatformErrorDetailReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *getPlatformErrorDetailReq) syllabLen() (ln int) {
	return int(req.syllabStackLen() + req.syllabHeapLen())
}

func (req *getPlatformErrorDetailReq) jsonDecoder(buf []byte) (err *er.Error) {
	var decoder = json.DecoderUnsafeMinifed{
		Buf: buf,
	}
	for err == nil {
		var keyName = decoder.DecodeKey()
		switch keyName {
		case "ErrorID":
			req.ErrorID, err = decoder.DecodeUInt64()
		default:
			err = decoder.NotFoundKeyStrict()
		}

		if len(decoder.Buf) < 3 {
			return
		}
	}
	return
}

func (req *getPlatformErrorDetailReq) jsonEncoder() (buf []byte) {
	var encoder = json.Encoder{
		Buf: make([]byte, 0, req.jsonLen()),
	}

	encoder.EncodeString(`{"ErrorID":`)
	encoder.EncodeUInt64(req.ErrorID)

	encoder.EncodeByte('}')
	return encoder.Buf
}

func (req *getPlatformErrorDetailReq) jsonLen() (ln int) {
	ln = 42
	return
}
