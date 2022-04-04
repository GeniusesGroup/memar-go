/* For license and copyright information please see LEGAL file in repository */

package utc

import (
	"testing"
)

func TestWeekdays_Check(t *testing.T) {
	type args struct {
		weekdays Weekdays
		day      Weekdays
	}
	tests := []struct {
		name      string
		args      args
		wantExist bool
	}{
		{
			name: "test1",
			args: args{
				weekdays: Weekdays_Saturday | Weekdays_Monday,
				day:      Weekdays_Saturday,
			},
			wantExist: true,
		}, {
			name: "reverse not work",
			args: args{
				weekdays: Weekdays_Saturday | Weekdays_Monday,
				day:      Weekdays_Friday,
			},
			wantExist: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var gotExist = tt.args.weekdays.Check(tt.args.day)
			if gotExist != tt.wantExist {
				t.Errorf("Weekdays_Check() = %v, want %v", gotExist, tt.wantExist)
			}
		})
	}
}
