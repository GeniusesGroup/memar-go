/* For license and copyright information please see the LEGAL file in the code repository */

package integer

import (
	"fmt"
	"testing"

	errs "memar/math/errors"
	"memar/protocol"
)

func Test_U64_FromString_Base10(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantNum U64
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
			name:    "4294967295",
			args:    args{str: "4294967295"},
			wantNum: 4294967295,
		},
		{
			name:    "18446744073709551615",
			args:    args{str: "18446744073709551615"},
			wantNum: 18446744073709551615,
		},
		{
			name:    "184467440737095516151",
			args:    args{str: "184467440737095516151"},
			wantErr: &errs.ErrValueOutOfRange,
		},
		{
			name:    "154565a",
			args:    args{str: "154565a"},
			wantErr: &errs.ErrBadValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var u64 U64
			gotErr := u64.FromString_Base10(tt.args.str)
			if u64 != tt.wantNum {
				t.Errorf("Test_U64_FromString_Base10() u64 = %v, want %v", u64, tt.wantNum)
			}
			if gotErr != tt.wantErr {
				t.Errorf("Test_U64_FromString_Base10() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			fmt.Println("Test_U64_FromString_Base10 - Test number:", tt.args.str, u64)
		})
	}
}
