/* For license and copyright information please see LEGAL file in repository */

package convert

import (
	"fmt"
	"testing"

	"../protocol"
)

func TestStringToUint8Base10(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantNum uint8
		wantErr protocol.Error
	}{
		{
			name:    "empty",
			args:    args{str: ""},
			wantErr: ErrEmptyValue,
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
			wantErr: ErrValueOutOfRange,
		},
		{
			name:    "15a",
			args:    args{str: "15a"},
			wantErr: ErrBadValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNum, gotErr := StringToUint8Base10(tt.args.str)
			if gotNum != tt.wantNum {
				t.Errorf("StringToUint8Base10() gotNum = %v, want %v", gotNum, tt.wantNum)
			}
			if gotErr != tt.wantErr {
				t.Errorf("StringToUint8Base10() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			fmt.Println("Test number:", tt.args.str, gotNum)
		})
	}
}

func TestStringToUint32Base10(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantNum uint32
		wantErr protocol.Error
	}{
		{
			name:    "empty",
			args:    args{str: ""},
			wantNum: 0,
			wantErr: ErrEmptyValue,
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
			wantErr: ErrValueOutOfRange,
		},
		{
			name:    "1545a",
			args:    args{str: "1545a"},
			wantErr: ErrBadValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNum, gotErr := StringToUint32Base10(tt.args.str)
			if gotNum != tt.wantNum {
				t.Errorf("StringToUint32Base10() gotNum = %v, want %v", gotNum, tt.wantNum)
			}
			if gotErr != tt.wantErr {
				t.Errorf("StringToUint32Base10() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			fmt.Println("Test number:", tt.args.str, gotNum)
		})
	}
}

func TestStringToUint64Base10(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantNum uint64
		wantErr protocol.Error
	}{
		{
			name:    "empty",
			args:    args{str: ""},
			wantNum: 0,
			wantErr: ErrEmptyValue,
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
			wantErr: ErrValueOutOfRange,
		},
		{
			name:    "154565a",
			args:    args{str: "154565a"},
			wantErr: ErrBadValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNum, gotErr := StringToUint64Base10(tt.args.str)
			if gotNum != tt.wantNum {
				t.Errorf("StringToUint64Base10() gotNum = %v, want %v", gotNum, tt.wantNum)
			}
			if gotErr != tt.wantErr {
				t.Errorf("StringToUint64Base10() gotErr = %v, want %v", gotErr, tt.wantErr)
			}
			fmt.Println("Test number:", tt.args.str, gotNum)
		})
	}
}
