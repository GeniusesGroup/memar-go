/* For license and copyright information please see LEGAL file in repository */

package sep

import (
	"fmt"
	"reflect"
	"testing"

	er "../../error"
)

func Test_posGetIdentifierRes_jsonDecoder(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name    string
		res     *posGetIdentifierRes
		args    args
		wantErr *er.Error
	}{
		{
			name: "test",
			res:  &posGetIdentifierRes{},
			args: args{
				buf: []byte(`{"IsSuccess":true,"ErrorCode":0,"ErrorDescription":null,"Data":{"Identifier":"f80a098f-aa49-4967-987b"}}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := tt.res.jsonDecoder(tt.args.buf)
			fmt.Println(*tt.res)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("posGetIdentifierRes.jsonDecoder() = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
