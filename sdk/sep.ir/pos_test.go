/* For license and copyright information please see LEGAL file in repository */

package sep

import (
	"fmt"
	"testing"

	er "../../error"
)

// to run test fill {{}} fields!
func Test_pos(t *testing.T) {
	var pos POS
	pos.Init([]byte(`{
		"Username": "{{Username}}",
		"Password": "{{Password}}",
		"TerminalIDs": {
			"{{PosID}}": true
		}
	}`))

	fmt.Println(pos)

	var posSendOrderReq = POSSendOrderReq{
		TerminalID:      "{{PosID}}",
		Amount:          "10000",
		AccountType:     AccountTypeSingle,
		TransactionType: OrderTransactionPurchase,
		// TimeOut:         "9",
	}
	var posSendOrderRes *POSSendOrderRes
	var err *er.Error
	posSendOrderRes, err = pos.POSSendOrder(&posSendOrderReq)
	fmt.Println(*err)
	fmt.Println(*posSendOrderRes)
}
