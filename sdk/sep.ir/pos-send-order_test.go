/* For license and copyright information please see LEGAL file in repository */

package sep

import (
	"fmt"
	"reflect"
	"testing"

	er "../../error"
)

func Test_POSSendOrderRes_jsonDecoder(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name    string
		res     *POSSendOrderRes
		args    args
		wantErr *er.Error
	}{
		{
			name: "test",
			res:  &POSSendOrderRes{},
			args: args{
				buf: []byte(`{
					"IsSuccess":true,
					"ErrorCode":0,
					"ErrorDescription":"transaction done",
					"Data":{
						"Identifier":"d0523d75",
						"TerminalID":"123",
						"TransactionType":"Purchase",
						"AccountType":"Single",
						"CreateOn":"2020-11-29T17:44:23.1290997+03:30",
						"CreateBy":"ba78a22a",
						"ResponseCode":"00",
						"ResponseDescription":"عملیات موفق",
						"TraceNumber":"94",
						"ResNum":"testInvoice",
						"RRN":"814",
						"State":0,
						"StateDescription":"transaction done",
						"Amount":"10000",
						"AffectiveAmount":"10000",
						"PosAppVersion":"10.050.00PO",
						"PosProtocolVersion":"V1.0.0.0",
						"CardHashNumber":"7B8B",
						"CardMaskNumber":"6037-99##-####-####"}}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := tt.res.jsonDecoder(tt.args.buf)
			fmt.Println(*tt.res)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("POSSendOrderRes.jsonDecoder() = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
