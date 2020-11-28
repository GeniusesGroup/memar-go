/* For license and copyright information please see LEGAL file in repository */

package sep

import (
	"fmt"
	"reflect"
	"testing"

	er "../../error"
)

func Test_idnGetTokenRes_jsonDecoder(t *testing.T) {
	type args struct {
		buf []byte
	}
	tests := []struct {
		name    string
		res     *idnGetTokenRes
		args    args
		wantErr *er.Error
	}{
		{
			name: "test",
			res:  &idnGetTokenRes{},
			args: args{
				buf: []byte(`{"access_token":"ey",					"expires_in": 3600,					"token_type": "Bearer"}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := tt.res.jsonDecoder(tt.args.buf)
			fmt.Println(*tt.res)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("idnGetTokenRes.jsonDecoder() = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
