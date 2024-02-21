/* For license and copyright information please see the LEGAL file in the code repository */

package integer

import (
	"fmt"
	"testing"

	errs "memar/math/errors"
	"memar/protocol"
)

func Test_U32_FromString_Base10(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantNum U32
		wantErr protocol.Error
	}{
		{
			name:    "empty",
			args:    args{str: ""},
			wantNum: 0,
			wantErr: &errs.ErrEmptyValue,
		},
		{
			name:    "0",
			args:    args{str: "0"},
			wantNum: 0,
		},
		{
			name:    "158564",
			args:    args{str: "158564"},
			wantNum: 158564,
		},
		{
			name:    "4294967295",
			args:    args{str: "4294967295"},
			wantNum: 4294967295,
		},
		{
			name:    "429496729546",
			args:    args{str: "429496729546"},
			wantNum: 0,
			wantErr: &errs.ErrValueOutOfRange,
		},
		{
			name:    "1545a",
			args:    args{str: "1545a"},
			wantErr: &errs.ErrBadValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var u32 U32
			gotErr := u32.FromString_Base10(tt.args.str)
			if u32 != tt.wantNum {
				t.Errorf("Test_U32_FromString_Base10() u32 = %v, want %v", u32, tt.wantNum)
			}
			if gotErr != tt.wantErr {
				t.Errorf("Test_U32_FromString_Base10() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			fmt.Println("Test_U32_FromString_Base10 - Test number:", tt.args.str, u32)
		})
	}
}
