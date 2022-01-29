/* For license and copyright information please see LEGAL file in repository */

package time

import (
	"testing"

	"../protocol"
)

func TestCheckWeekdays(t *testing.T) {
	type args struct {
		weekdays protocol.Weekdays
		day      protocol.Weekdays
	}
	tests := []struct {
		name      string
		args      args
		wantExist bool
	}{
		{
			name: "test1",
			args: args{
				weekdays: protocol.WeekdaysSaturday | protocol.WeekdaysMonday,
				day:      protocol.WeekdaysSaturday,
			},
			wantExist: true,
		}, {
			name: "reverse not work",
			args: args{
				weekdays: protocol.WeekdaysSaturday,
				day:      protocol.WeekdaysSaturday | protocol.WeekdaysMonday,
			},
			wantExist: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotExist := CheckWeekdays(tt.args.weekdays, tt.args.day); gotExist != tt.wantExist {
				t.Errorf("CheckWeekdays() = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}

func TestCheckWeekdaysReverse(t *testing.T) {
	type args struct {
		day      protocol.Weekdays
		weekdays protocol.Weekdays
	}
	tests := []struct {
		name      string
		args      args
		wantExist bool
	}{
		{
			name: "test1",
			args: args{
				day:      protocol.WeekdaysSaturday,
				weekdays: protocol.WeekdaysSaturday | protocol.WeekdaysMonday,
			},
			wantExist: true,
		}, {
			name: "reverse not work",
			args: args{
				day:      protocol.WeekdaysSaturday | protocol.WeekdaysMonday,
				weekdays: protocol.WeekdaysSaturday,
			},
			wantExist: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotExist := CheckWeekdaysReverse(tt.args.day, tt.args.weekdays); gotExist != tt.wantExist {
				t.Errorf("CheckWeekdaysReverse() = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}
