/* For license and copyright information please see LEGAL file in repository */

package tcp

import "testing"

func Test_gc_check(t *testing.T) {
	type args struct {
		now int64
	}
	tests := []struct {
		name string
		gc   *gc
		args args
	}{
		{
			name: "test1",
			gc: &gc{
				keepAlive_Interval_next: 0,
			},
			args: args{
				now: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.gc.check(tt.args.now)
		})
	}
}
