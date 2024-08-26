/* For license and copyright information please see the LEGAL file in the code repository */

package duration

import (
	"testing"
)

func Test_nsecToSec(t *testing.T) {
	type args struct {
		d NanoSecond
	}
	tests := []struct {
		name     string
		args     args
		wantSec  Second
		wantNsec NanoInSecond
	}{
		{
			name: "test1",
			args: args{
				d: 1,
			},
			wantSec:  0,
			wantNsec: 1,
		},
		{
			name: "test2",
			args: args{
				d: 1*OneSecond + 1,
			},
			wantSec:  1,
			wantNsec: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSec, gotNsec := tt.args.d.ToSecAndNano()
			if gotSec != tt.wantSec {
				t.Errorf("nsecToSec() gotSec = %v, want %v", gotSec, tt.wantSec)
			}
			if gotNsec != tt.wantNsec {
				t.Errorf("nsecToSec() gotNsec = %v, want %v", gotNsec, tt.wantNsec)
			}
		})
	}
}
