/* For license and copyright information please see the LEGAL file in the code repository */

package unix

import (
	"fmt"
	"memar/time/duration"
	"testing"
	"time"
)

func Test_ShowNow(t *testing.T) {
	var ti Time = Now()
	fmt.Println(ti)
	fmt.Println(ti.NanoElapsed())
	fmt.Println(ti.MicroElapsed())
	fmt.Println(ti.MilliElapsed())
	fmt.Println(ti.SecElapsed())

	var timeNow = time.Now()
	fmt.Println(timeNow.Unix())
	fmt.Println(timeNow.UnixMilli())
	fmt.Println(timeNow.UnixMicro())
	fmt.Println(timeNow.UnixNano())
}

func TestTime_ElapsedByDuration(t *testing.T) {
	type args struct {
		d duration.NanoSecond
	}
	tests := []struct {
		name       string
		tr         Time
		args       args
		wantPeriod int64
	}{
		{
			name: "test1",
			tr: Time{
				sec:  1,
				nsec: 0,
			},
			args: args{
				d: 10,
			},
			wantPeriod: 1 * int64(duration.OneSecond) / 10,
		},
		{
			name: "test2",
			tr: Time{
				sec:  1,
				nsec: 100,
			},
			args: args{
				d: 10,
			},
			wantPeriod: (1*int64(duration.OneSecond) + 100) / 10,
		},
		{
			name: "test3",
			tr: Time{
				sec:  10,
				nsec: 100,
			},
			args: args{
				d: (1 * duration.OneSecond) + 10,
			},
			wantPeriod: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPeriod := tt.tr.ElapsedByDuration(tt.args.d); gotPeriod != tt.wantPeriod {
				t.Errorf("Time.ElapsedByDuration() = %v, want %v", gotPeriod, tt.wantPeriod)
			}
		})
	}
}
