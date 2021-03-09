/* For license and copyright information please see LEGAL file in repository */

package asanak

import (
	"../../achaemenid"
	"../../convert"
	er "../../error"
	"../../http"
	"../../json"
	lang "../../language"
	"../../srpc"
)

var rec Asanak

var getSmsByAsanakService = achaemenid.Service{
	ID:                3261053697,
	URI:               "/apis/get-sms-by-asanak",
	IssueDate:         1593342176,
	ExpiryDate:        0,
	ExpireInFavorOf:   "",
	ExpireInFavorOfID: 0,
	Status:            achaemenid.ServiceStatePreAlpha,

	Name: map[lang.Language]string{
		lang.LanguageEnglish: "GetSmsByAsanak",
	},
	Description: map[lang.Language]string{
		lang.LanguageEnglish: "register and receive income SMS from Asanak",
	},
	TAGS: []string{
		"Asanak", "SMS", "Iran-SMS",
	},

	SRPCHandler: GetSmsByAsanakSRPC,
	HTTPHandler: rec.GetSmsByAsanakHTTP,
}

// GetSmsByAsanakSRPC is sRPC handler of GetSmsByAsanak service.
func GetSmsByAsanakSRPC(st *achaemenid.Stream) {
	var req = &GetSmsByAsanak{}
	st.Err = req.syllabDecoder(srpc.GetPayload(st.IncomePayload))
	if st.Err != nil {
		st.Connection.FailedPacketsReceived++
		// Attack??
		return
	}
}

// GetSmsByAsanakHTTP is HTTP handler of GetSmsByAsanak service.
// {{URL}}?Destination=$Destination&Source=$Source&ReceiveTime=$ReceiveTime&MsgBody=$MsgBody
func (a *Asanak) GetSmsByAsanakHTTP(st *achaemenid.Stream, httpReq *http.Request, httpRes *http.Response) {
	var req = &GetSmsByAsanak{}
	st.Err = req.jsonDecoder(httpReq.Body)
	if st.Err != nil {
		st.Connection.FailedPacketsReceived++
		httpRes.SetStatus(http.StatusBadRequestCode, http.StatusBadRequestPhrase)
		httpRes.Body = convert.UnsafeStringToByteSlice(st.Err.Error())
		// Attack??
		return
	}

	// TODO::: decode data to GetSmsByAsanak and send to a.RecChan

	httpRes.SetStatus(http.StatusOKCode, http.StatusOKPhrase)
	httpRes.Header.Set(http.HeaderKeyContentType, "application/json")
}

// GetSmsByAsanak is data that receive by asanak and pass to chanel to handle by desire logic!
type GetSmsByAsanak struct {
	Destination string
	Source      string
	ReceiveTime string
	MsgBody     string
}

func (req *GetSmsByAsanak) syllabDecoder(buf []byte) (err *er.Error) {
	return
}

func (req *GetSmsByAsanak) jsonDecoder(buf []byte) (err *er.Error) {
	err = json.UnMarshal(buf, req)
	return
}
