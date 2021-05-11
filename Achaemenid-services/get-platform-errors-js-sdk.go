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

var getPlatfromErrorsSDKService = achaemenid.Service{
	URN:                "urn:giti:achaemenid.libgo:service:get-platform-errors-js-sdk",
	Domain:             DomainName,
	ID:                 5727209077706595565,
	IssueDate:          1620305281,
	ExpiryDate:         0,
	ExpireInFavorOfURN: "",
	ExpireInFavorOfID:  0,
	Status:             achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.LanguageEnglish: "Get Platform Errors JS SDK",
	},
	Description: map[lang.Language]string{
		lang.LanguageEnglish: "Returns all platform errors in JS language",
	},
	TAGS: []string{
		"",
	},

	SRPCHandler: GetPlatfromErrorsSDKSRPC,
	HTTPHandler: GetPlatfromErrorsSDKHTTP,
}

// GetPlatfromErrorsSDKSRPC is sRPC handler of GetPlatfromErrorsSDK service.
func GetPlatfromErrorsSDKSRPC(st *achaemenid.Stream) {
	var req = &getPlatfromErrorsSDKReq{}
	st.Err = req.syllabDecoder(srpc.GetPayload(st.IncomePayload))
	if st.Err != nil {
		return
	}

	var res *getPlatfromErrorsSDKRes
	res, st.Err = getPlatfromErrorsSDK(st, req)
	// Check if any error occur in bussiness logic
	if st.Err != nil {
		return
	}

	st.OutcomePayload = res.SDK
}

// GetPlatfromErrorsSDKHTTP is HTTP handler of GetPlatfromErrorsSDK service.
func GetPlatfromErrorsSDKHTTP(st *achaemenid.Stream, httpReq *http.Request, httpRes *http.Response) {
	var req = &getPlatfromErrorsSDKReq{}
	st.Err = req.jsonDecoder(httpReq.Body)
	if st.Err != nil {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		return
	}

	var res *getPlatfromErrorsSDKRes
	res, st.Err = getPlatfromErrorsSDK(st, req)
	// Check if any error occur in bussiness logic
	if st.Err != nil {
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		return
	}

	httpRes.SetStatus(http.StatusOKCode, http.StatusOKPhrase)
	httpRes.Header.Set(http.HeaderKeyContentType, "application/javascript")
	httpRes.Body = res.SDK
}

type getPlatfromErrorsSDKReq struct {
	Language lang.Language
}

type getPlatfromErrorsSDKRes struct {
	SDK []byte
}

func getPlatfromErrorsSDK(st *achaemenid.Stream, req *getPlatfromErrorsSDKReq) (res *getPlatfromErrorsSDKRes, err *er.Error) {
	res = &getPlatfromErrorsSDKRes{
		SDK: er.Errors.GetErrorsInJsFormat(req.Language),
	}
	return
}

/*
	Request Encoders & Decoders
*/

func (req *getPlatfromErrorsSDKReq) syllabDecoder(buf []byte) (err *er.Error) {
	if uint32(len(buf)) < req.syllabStackLen() {
		err = syllab.ErrShortSliceDecoded
		return
	}

	req.Language = lang.Language(syllab.GetUInt32(buf, 0))
	return
}

func (req *getPlatfromErrorsSDKReq) syllabEncoder(buf []byte) {
	syllab.SetUInt32(buf, 0, uint32(req.Language))
	return
}

func (req *getPlatfromErrorsSDKReq) syllabStackLen() (ln uint32) {
	return 4
}

func (req *getPlatfromErrorsSDKReq) syllabHeapLen() (ln uint32) {
	return
}

func (req *getPlatfromErrorsSDKReq) syllabLen() (ln int) {
	return int(req.syllabStackLen() + req.syllabHeapLen())
}

func (req *getPlatfromErrorsSDKReq) jsonDecoder(buf []byte) (err *er.Error) {
	var decoder = json.DecoderUnsafeMinifed{
		Buf: buf,
	}
	for err == nil {
		var keyName = decoder.DecodeKey()
		switch keyName {
		case "Language":
			var num uint32
			num, err = decoder.DecodeUInt32()
			req.Language = lang.Language(num)
		default:
			err = decoder.NotFoundKeyStrict()
		}

		if len(decoder.Buf) < 3 {
			return
		}
	}
	return
}

func (req *getPlatfromErrorsSDKReq) jsonEncoder() (buf []byte) {
	var encoder = json.Encoder{
		Buf: make([]byte, 0, req.jsonLen()),
	}

	encoder.EncodeString(`{"Language":`)
	encoder.EncodeUInt32(uint32(req.Language))

	encoder.EncodeByte('}')
	return encoder.Buf
}

func (req *getPlatfromErrorsSDKReq) jsonLen() (ln int) {
	ln = 23
	return
}
