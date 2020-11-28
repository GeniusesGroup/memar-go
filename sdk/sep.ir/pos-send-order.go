/* For license and copyright information please see LEGAL file in repository */

package sep

import (
	er "../../error"
	"../../http"
	"../../json"
	"../../log"
)

// POSSendOrderReq is request structure of POSSendOrder()
type POSSendOrderReq struct {
	TerminalID           string
	Amount               string
	AccountType          accountType
	Additional           string
	TimeOut              string // in minutes
	ResNum               string // e.g. InvoiceID that can be check later if needed
	Identifier           string
	UserNotifiable       POSUserNotifiable
	TransactionType      transactionType
	BillID               string
	PayID                string
	PurchaseID           string
	ReferenceData        string
	AmountIBAN           string
	TotalAmount          string `json:"totalAmount"`
	PurchaseIDAmountIBAN string `json:"PurchaseId_Amont_Iban"`
	// SwitchID             string
}

// POSSendOrderRes is response structure of POSSendOrder()
type POSSendOrderRes struct {
	IsSuccess        bool
	ErrorCode        int64
	ErrorDescription string

	Identifier          string
	TerminalID          string
	CreatedOn           string
	ResponseCode        string
	ResponseDescription string
	TraceNumber         string
	ResNum              string
	RRN                 string // Shaparak.ir Universal ID
	State               int64
	StateDescription    string
	HashData            string
	Amount              string // May be different from request Amount just & only if you have contract with SEP.ir to have some discount debit cards!
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

	var serverReq []byte = httpReq.Marshal()
	var serverRes []byte
	serverRes, err = pos.sendHTTPRequest(serverReq)
	if err != nil {
		return
	}

	if log.DevMode {
		log.Debug("sep.ir - Send msg to /v1/PcPosTransaction/StartPayment ::\n", string(serverReq))
		log.Debug("sep.ir - Received msg from /v1/PcPosTransaction/StartPayment::\n", string(serverRes))
	}

	var httpRes = http.MakeNewResponse()
	err = httpRes.UnMarshal(serverRes)
	if err != nil {
		return
	}

	switch httpRes.StatusCode {
	case http.StatusBadRequestCode:
		err = ErrBadRequest
		return
	case http.StatusUnauthorizedCode:
		log.Warn("Authorization failed with sep services! Double check account username & password")
	case http.StatusInternalServerErrorCode:
		err = ErrInternalError
		return
	}

	res = &POSSendOrderRes{}
	err = res.jsonDecoder(httpRes.Body)
	if err != nil {
		err = ErrBadResponse
		return
	}
	err = getErrorByResCode(res.ErrorCode)
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

	encoder.EncodeString(`,"Additional":"`)
	encoder.EncodeString(req.Additional)

	encoder.EncodeString(`","TimeOut":"`)
	encoder.EncodeString(req.TimeOut)

	encoder.EncodeString(`","ResNum":"`)
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

	encoder.EncodeString(`,"ReferenceData":"`)
	encoder.EncodeString(req.ReferenceData)

	encoder.EncodeString(`","AmountIBAN":"`)
	encoder.EncodeString(req.AmountIBAN)

	if req.TotalAmount != "" {
		encoder.EncodeString(`","totalAmount":"`)
		encoder.EncodeString(req.TotalAmount)
	}

	encoder.EncodeString(`","PurchaseId_Amont_Iban":"`)
	encoder.EncodeString(req.PurchaseIDAmountIBAN)

	encoder.EncodeString("\"}")
	return encoder.Buf
}

func (req *POSSendOrderReq) jsonLen() (ln int) {
	ln = len(req.TerminalID) + len(req.Amount) + len(req.Additional) + len(req.TimeOut) + len(req.ResNum) + len(req.Identifier) + len(req.BillID) + len(req.PayID) + len(req.PurchaseID) + len(req.ReferenceData) + len(req.AmountIBAN) + len(req.TotalAmount) + len(req.PurchaseIDAmountIBAN)
	ln += req.UserNotifiable.jsonLen()
	ln += 275
	return
}

func (res *POSSendOrderRes) jsonDecoder(buf []byte) (err *er.Error) {
	var decoder = json.DecoderUnsafe{
		Buf: buf,
	}
	var keyName string
	for len(decoder.Buf) > 2 {
		keyName = decoder.DecodeKey()
		switch keyName {
		case "":
			return
		case "IsSuccess":
			res.IsSuccess, err = decoder.DecodeBool()
			if err != nil {
				return
			}
		case "ErrorCode":
			res.ErrorCode, err = decoder.DecodeInt64()
			if err != nil {
				return
			}
		case "ErrorDescription":
			res.ErrorDescription, err = decoder.DecodeString()
			if err != nil {
				return
			}

		case "Data":
			continue
		case "Identifier":
			res.Identifier, err = decoder.DecodeString()
			if err != nil {
				return
			}
		case "TerminalID":
			res.TerminalID, err = decoder.DecodeString()
			if err != nil {
				return
			}
		case "CreatedOn":
			res.CreatedOn, err = decoder.DecodeString()
			if err != nil {
				return
			}
		case "ResponseCode":
			res.ResponseCode, err = decoder.DecodeString()
			if err != nil {
				return
			}
		case "ResponseDescription":
			res.ResponseDescription, err = decoder.DecodeString()
			if err != nil {
				return
			}
		case "TraceNumber":
			res.TraceNumber, err = decoder.DecodeString()
			if err != nil {
				return
			}
		case "ResNum":
			res.ResNum, err = decoder.DecodeString()
			if err != nil {
				return
			}
		case "RRN":
			res.RRN, err = decoder.DecodeString()
			if err != nil {
				return
			}
		case "State":
			res.State, err = decoder.DecodeInt64()
			if err != nil {
				return
			}
		case "StateDescription":
			res.StateDescription, err = decoder.DecodeString()
			if err != nil {
				return
			}
		case "HashData":
			res.HashData, err = decoder.DecodeString()
			if err != nil {
				return
			}
		case "Amount":
			res.Amount, err = decoder.DecodeString()
			if err != nil {
				return
			}
		default:
			err = decoder.NotFoundKey()
			if err != nil {
				return
			}
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
	Alignment   int
	ReceiptType int
}

// JSONEncoder encode given POSPrintItem to json format.
func (pos *POSPrintItem) jsonEncoder(encoder *json.Encoder) {
	encoder.EncodeString(`{"Item":"`)
	encoder.EncodeString(pos.Item)

	encoder.EncodeString(`","Value":"`)
	encoder.EncodeString(pos.Value)

	encoder.EncodeString(`","Alignment":`)
	encoder.EncodeInt64(int64(pos.Alignment))

	encoder.EncodeString(`,"ReceiptType":`)
	encoder.EncodeInt64(int64(pos.ReceiptType))

	encoder.EncodeString("},")
}

// JSONLen return json needed len to encode!
func (pos *POSPrintItem) jsonLen() (ln int) {
	ln = len(pos.Item) + len(pos.Value)
	ln += 90
	return
}

type accountType uint8

// Account Types
const (
	AccountTypeSingle            accountType = 0
	AccountTypeShare             accountType = 1
	AccountTypeShareByAmountIBAN accountType = 2
)

type transactionType uint8

// Order Transaction Types
const (
	OrderTransactionPurchase   transactionType = 0
	OrderTransactionBill       transactionType = 1
	OrderTransactionPurchaseID transactionType = 2
)
