/* For license and copyright information please see the LEGAL file in the code repository */

package integer

import (
	"fmt"
	"testing"

	errs "memar/math/errors"
	"memar/protocol"
)

func Test_U16_FromString_Base10(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantNum U16
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
			name:    "65535",
			args:    args{str: "65535"},
			wantNum: 65535,
		},
		{
			name:    "655355",
			args:    args{str: "655355"},
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
			var u16 U16
			gotErr := u16.FromString_Base10(tt.args.str)
			if u16 != tt.wantNum {
				t.Errorf("Test_U16_FromString_Base10() gotNum = %v, want %v", u16, tt.wantNum)
			}
			if gotErr != tt.wantErr {
				t.Errorf("Test_U16_FromString_Base10() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			fmt.Println("Test_U16_FromString_Base10 - Test number:", tt.args.str, u16)
		})
	}
}
