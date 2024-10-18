/* For license and copyright information please see the LEGAL file in the code repository */

package duration

import (
	"reflect"
	"testing"
)

func TestNanoSecond_ToSecAndNano(t *testing.T) {
	tests := []struct {
		name     string
		d        NanoSecond
		wantSec  Second
		wantNsec NanoInSecond
	}{
		{
			name:     "test1",
			d:        1,
			wantSec:  0,
			wantNsec: 1,
		},
		{
			name:     "test2",
			d:        1*NanoSecondInSecond + 1,
			wantSec:  1,
			wantNsec: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSec, gotNsec := tt.d.ToSecAndNano()
			if !reflect.DeepEqual(gotSec, tt.wantSec) {
				t.Errorf("NanoSecond.ToSecAndNano() gotSec = %v, want %v", gotSec, tt.wantSec)
			}
			if !reflect.DeepEqual(gotNsec, tt.wantNsec) {
				t.Errorf("NanoSecond.ToSecAndNano() gotNsec = %v, want %v", gotNsec, tt.wantNsec)
			}
		})
	}
}
