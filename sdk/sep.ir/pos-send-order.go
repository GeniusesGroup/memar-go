/* For license and copyright information please see LEGAL file in repository */

package sep

import (
	"../../achaemenid"
	er "../../error"
	"../../http"
	"../../json"
	"../../log"
)

// POSSendOrderReq is request structure of POSSendOrder()
type POSSendOrderReq struct {
	TerminalID           string
	Amount               string
	AccountType          posAccountType
	TimeOut              string `json:",optional"` // in minutes but not implement yet by SEP so don't send it now!
	ResNum               string // e.g. InvoiceID that can be check later if needed.
	Identifier           string
	UserNotifiable       POSUserNotifiable
	TransactionType      posTransactionType
	BillID               string `json:",optional"`
	PayID                string `json:",optional"`
	PurchaseID           string `json:",optional"`             // set just if TransactionType==POSTransactionPurchaseID
	ReferenceData        string `json:"RefrenceData,optional"` // store by switch for analytics purpose
	AmountIBAN           string `json:"Amount_Iban,optional"`
	TotalAmount          string `json:"totalAmount,optional"`
	PurchaseIDAmountIBAN string `json:"PurchaseID_Amount_Iban,optional"`
	// SwitchID             string
}

// POSSendOrderRes is response structure of POSSendOrder()
type POSSendOrderRes struct {
	IsSuccess        bool
	ErrorCode        int64
	ErrorDescription string

	Identifier          string
	TerminalID          string
	TransactionType     string // name of request transaction type!! really WHY???
	AccountType         string // name of request account type!! really WHY???
	CreateOn            string
	CreateBy            string
	ResponseCode        string
	ResponseDescription string
	TraceNumber         string
	ResNum              string
	RRN                 string // Shaparak.ir Universal ID
	State               int64
	StateDescription    string
	Amount              string
	AffectiveAmount     string // May be different from request Amount just & only if you have contract with SEP.ir to have some discount debit cards!
	PosAppVersion       string
	PosProtocolVersion  string
	CardHashNumber      string
	CardMaskNumber      string
}

// POSSendOrder use to send order to pos
func (pos *POS) POSSendOrder(req *POSSendOrderReq) (res *POSSendOrderRes, err *er.Error) {
	err = pos.checkTerminalID(req.TerminalID)
	if err != nil {
		return
	}

	err = pos.IDN.checkAccessToken(pos.Username, pos.Password)
	if err != nil {
		return
	}

	req.Identifier, err = pos.posGetIdentifier()
	if err != nil {
		return
	}

	if req.UserNotifiable.PrintItems == nil {
		req.UserNotifiable.PrintItems = []POSPrintItem{
			{
				Item:        "",
				Value:       "",
				Alignment:   POSPrintAlignmentCenter,
				ReceiptType: POSPrintReceiptTypeBoth,
			},
		}
	}

	var httpReq = http.MakeNewRequest()
	httpReq.Method = http.MethodPOST
	httpReq.URI.Path = posPathSendOrder
	httpReq.Version = http.VersionHTTP11

	httpReq.Header.Set(http.HeaderKeyHost, posDomain)
	httpReq.Header.Set(http.HeaderKeyAcceptContent, "application/json")
	httpReq.Header.Set(http.HeaderKeyContentType, "application/json")
	httpReq.Header.Set(http.HeaderKeyAuthorization, pos.IDN.TokenType+" "+pos.IDN.AccessToken)

	httpReq.Body = req.jsonEncoder()

	// Set some other header data
	httpReq.SetContentLength()
	httpReq.Header.Set(http.HeaderKeyServer, http.DefaultServer)

	var httpRes *http.Response
	httpRes, err = pos.doHTTP(httpReq)
	if err != nil {
		if err.Equal(achaemenid.ErrReceiveResponse) {
			// TODO::: check order
		}
		return
	}

	switch httpRes.StatusCode {
	case http.StatusOKCode:
		// TODO::: anything???
	case http.StatusBadRequestCode:
		err = achaemenid.ErrBadRequest
		return
	case http.StatusUnauthorizedCode:
		log.Warn("Authorization failed with sep services! Double check account username & password")
	case http.StatusInternalServerErrorCode:
		err = achaemenid.ErrInternalError
		return
	}

	switch httpRes.Header.GetContentType().SubType {
	case http.ContentTypeMimeSubTypeHTML:
		err = ErrPOSInternalError
		return
	case http.ContentTypeMimeSubTypeJSON:
		res = &POSSendOrderRes{}
		err = res.jsonDecoder(httpRes.Body)
		if err != nil {
			err = achaemenid.ErrBadResponse
			return
		}
		if !res.IsSuccess && res.ErrorCode == 0 {
			err = ErrPOSInternalError
			return
		}
		err = getErrorByResCode(res.ErrorCode)
	}
	return
}

func (req *POSSendOrderReq) jsonEncoder() (buf []byte) {
	var encoder = json.Encoder{
		Buf: make([]byte, 0, req.jsonLen()),
	}

	encoder.EncodeString(`{"TerminalID":"`)
	encoder.EncodeString(req.TerminalID)

	encoder.EncodeString(`","Amount":"`)
	encoder.EncodeString(req.Amount)

	encoder.EncodeString(`","AccountType":`)
	encoder.EncodeUInt8(uint8(req.AccountType))

	if req.TimeOut != "" {
		encoder.EncodeString(`,"TimeOut":"`)
		encoder.EncodeString(req.TimeOut)
		encoder.EncodeByte('"')
	}

	encoder.EncodeString(`,"ResNum":"`)
	encoder.EncodeString(req.ResNum)

	encoder.EncodeString(`","Identifier":"`)
	encoder.EncodeString(req.Identifier)

	encoder.EncodeString(`","UserNotifiable":`)
	req.UserNotifiable.jsonEncoder(&encoder)

	encoder.EncodeString(`,"TransactionType":`)
	encoder.EncodeUInt8(uint8(req.TransactionType))

	if req.BillID != "" {
		encoder.EncodeString(`,"BillID":"`)
		encoder.EncodeString(req.BillID)
		encoder.EncodeByte('"')
	}

	if req.PayID != "" {
		encoder.EncodeString(`,"PayID":"`)
		encoder.EncodeString(req.PayID)
		encoder.EncodeByte('"')
	}

	if req.PurchaseID != "" {
		encoder.EncodeString(`,"PurchaseID":"`)
		encoder.EncodeString(req.PurchaseID)
		encoder.EncodeByte('"')
	}

	if req.ReferenceData != "" {
		encoder.EncodeString(`,"ReferenceData":"`)
		encoder.EncodeString(req.ReferenceData)
		encoder.EncodeByte('"')
	}

	if req.AmountIBAN != "" {
		encoder.EncodeString(`","Amount_Iban":"`)
		encoder.EncodeString(req.AmountIBAN)
		encoder.EncodeByte('"')
	}

	if req.TotalAmount != "" {
		encoder.EncodeString(`","totalAmount":"`)
		encoder.EncodeString(req.TotalAmount)
		encoder.EncodeByte('"')
	}

	if req.PurchaseIDAmountIBAN != "" {
		encoder.EncodeString(`","PurchaseID_Amount_Iban":"`)
		encoder.EncodeString(req.PurchaseIDAmountIBAN)
		encoder.EncodeByte('"')
	}

	encoder.EncodeByte('}')
	return encoder.Buf
}

func (req *POSSendOrderReq) jsonLen() (ln int) {
	ln = len(req.TerminalID) + len(req.Amount) + len(req.TimeOut) + len(req.ResNum) + len(req.Identifier) + len(req.BillID) + len(req.PayID) + len(req.PurchaseID) + len(req.ReferenceData) + len(req.AmountIBAN) + len(req.TotalAmount) + len(req.PurchaseIDAmountIBAN)
	ln += req.UserNotifiable.jsonLen()
	ln += 260
	return
}

func (res *POSSendOrderRes) jsonDecoder(buf []byte) (err *er.Error) {
	var decoder = json.DecoderUnsafe{
		Buf: buf,
	}
	for err == nil {
		var keyName = decoder.DecodeKey()
		switch keyName {
		case "":
			return
		case "IsSuccess":
			res.IsSuccess, err = decoder.DecodeBool()
		case "ErrorCode":
			res.ErrorCode, err = decoder.DecodeInt64()
		case "ErrorDescription":
			res.ErrorDescription, err = decoder.DecodeString()

		case "Data":
			continue
		case "Identifier":
			res.Identifier, err = decoder.DecodeString()
		case "TerminalID":
			res.TerminalID, err = decoder.DecodeString()
		case "TransactionType":
			res.TransactionType, err = decoder.DecodeString()
		case "AccountType":
			res.AccountType, err = decoder.DecodeString()
		case "CreateOn":
			res.CreateOn, err = decoder.DecodeString()
		case "CreateBy":
			res.CreateBy, err = decoder.DecodeString()
		case "ResponseCode":
			res.ResponseCode, err = decoder.DecodeString()
		case "ResponseDescription":
			res.ResponseDescription, err = decoder.DecodeString()
		case "TraceNumber":
			res.TraceNumber, err = decoder.DecodeString()
		case "ResNum":
			res.ResNum, err = decoder.DecodeString()
		case "RRN":
			res.RRN, err = decoder.DecodeString()
		case "State":
			res.State, err = decoder.DecodeInt64()
		case "StateDescription":
			res.StateDescription, err = decoder.DecodeString()
		case "Amount":
			res.Amount, err = decoder.DecodeString()
		case "AffectiveAmount":
			res.AffectiveAmount, err = decoder.DecodeString()
		case "PosAppVersion":
			res.PosAppVersion, err = decoder.DecodeString()
		case "PosProtocolVersion":
			res.PosProtocolVersion, err = decoder.DecodeString()
		case "CardHashNumber":
			res.CardHashNumber, err = decoder.DecodeString()
		case "CardMaskNumber":
			res.CardMaskNumber, err = decoder.DecodeString()
		default:
			err = decoder.NotFoundKey()
		}

		if len(decoder.Buf) < 3 {
			return
		}
	}
	return
}

// POSUserNotifiable indicate data to print on User receiption paper
type POSUserNotifiable struct {
	FooterMessage string
	PrintItems    []POSPrintItem
}

// JSONEncoder encode given POSUserNotifiable to json format.
func (pos *POSUserNotifiable) jsonEncoder(encoder *json.Encoder) {
	encoder.EncodeString(`{"FooterMessage":"`)
	encoder.EncodeString(pos.FooterMessage)

	encoder.EncodeString(`","PrintItems":[`)
	for i := 0; i < len(pos.PrintItems); i++ {
		pos.PrintItems[i].jsonEncoder(encoder)
		encoder.EncodeByte(',')
	}
	encoder.RemoveTrailingComma()

	encoder.EncodeString(`]}`)
}

// JSONLen return json needed len to encode!
func (pos *POSUserNotifiable) jsonLen() (ln int) {
	ln = len(pos.FooterMessage)
	for i := 0; i < len(pos.PrintItems); i++ {
		ln += pos.PrintItems[i].jsonLen()
	}
	ln += 36
	return
}

// POSPrintItem indicate each data to print
type POSPrintItem struct {
	Item        string
	Value       string
	Alignment   posPrintAlignment
	ReceiptType posPrintReceiptType
}

// JSONEncoder encode given POSPrintItem to json format.
func (pos *POSPrintItem) jsonEncoder(encoder *json.Encoder) {
	encoder.EncodeString(`{"Item":"`)
	encoder.EncodeString(pos.Item)

	encoder.EncodeString(`","Value":"`)
	encoder.EncodeString(pos.Value)

	encoder.EncodeString(`","Alignment":`)
	encoder.EncodeUInt8(uint8(pos.Alignment))

	encoder.EncodeString(`,"ReceiptType":`)
	encoder.EncodeUInt8(uint8(pos.ReceiptType))

	encoder.EncodeString("},")
}

// JSONLen return json needed len to encode!
func (pos *POSPrintItem) jsonLen() (ln int) {
	ln = len(pos.Item) + len(pos.Value)
	ln += 90
	return
}

type posPrintAlignment uint8

// POSPrintAlignment items
const (
	POSPrintAlignmentRight  posPrintAlignment = 0
	POSPrintAlignmentLeft   posPrintAlignment = 1
	POSPrintAlignmentCenter posPrintAlignment = 2
)

type posPrintReceiptType uint8

// POSPrintReceiptType items
const (
	POSPrintReceiptTypeCustomer posPrintReceiptType = 0
	POSPrintReceiptTypeMerchant posPrintReceiptType = 1
	POSPrintReceiptTypeBoth     posPrintReceiptType = 2
)

type posAccountType uint8

// Account Types
const (
	POSAccountTypeSingle            posAccountType = 0
	POSAccountTypeShare             posAccountType = 1
	POSAccountTypeShareByAmountIBAN posAccountType = 2
)

type posTransactionType uint8

// POS Transaction Types
const (
	POSTransactionPurchase   posTransactionType = 0
	POSTransactionBill       posTransactionType = 1
	POSTransactionPurchaseID posTransactionType = 2
)
