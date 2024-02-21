/* For license and copyright information please see the LEGAL file in the code repository */

package integer

import (
	"fmt"
	"testing"

	errs "memar/math/errors"
	"memar/protocol"
)

func Test_U8_FromString_Base10(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantNum U8
		wantErr protocol.Error
	}{
		{
			name:    "empty",
			args:    args{str: ""},
			wantErr: &errs.ErrEmptyValue,
		},
		{
			name:    "0",
			args:    args{str: "0"},
			wantNum: 0,
		},
		{
			name:    "158",
			args:    args{str: "158"},
			wantNum: 158,
		},
		{
			name:    "255",
			args:    args{str: "255"},
			wantNum: 255,
		},
		{
			name:    "1582",
			args:    args{str: "1582"},
			wantErr: &errs.ErrValueOutOfRange,
		},
		{
			name:    "15a",
			args:    args{str: "15a"},
			wantErr: &errs.ErrBadValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var u8 U8
			gotErr := u8.FromString_Base10(tt.args.str)
			if u8 != tt.wantNum {
				t.Errorf("Test_U8_FromString_Base10() gotNum = %v, want %v", u8, tt.wantNum)
			}
			if gotErr != tt.wantErr {
				t.Errorf("Test_U8_FromString_Base10() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			fmt.Println("Test_U8_FromString_Base10 - Test number:", tt.args.str, u8)
		})
	}
}
