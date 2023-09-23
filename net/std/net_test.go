/* For license and copyright information please see the LEGAL file in the code repository */

package std

import (
	"testing"
)

func TestSocket_Read(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		s       *Socket
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name: "buffer full and return",
			s: &Socket{
				recv: recv{
					// buf: buffer.Queue{},
				},
			},
			args: args{
				b: make([]byte, 120),
			},
			wantN:   120,
			wantErr: false,
		},
		{
			name: "buffer not full and timer based",
			s: &Socket{
				recv: recv{
					// readTimer: timer.Timer{},
					// buf: buffer.Queue{},
				},
			},
			args: args{
				b: make([]byte, 120),
			},
			wantN:   120,
			wantErr: false,
		},
		{
			name: "buffer not full and signal based",
			s: &Socket{
				recv: recv{
					// readTimer: timer.Timer{},
					// buf: buffer.Queue{},
				},
			},
			args: args{
				b: make([]byte, 120),
			},
			wantN:   120,
			wantErr: false,
		},
	}

	tests[0].s.recv.buf.Init(60)
	tests[0].s.recv.buf.Write(make([]byte, 60))

	tests[1].s.recv.buf.Init(1024)
	tests[1].s.recv.readTimer.Init()
	tests[1].s.recv.readTimer.Start(10000)
	tests[1].s.recv.buf.Write(make([]byte, 160))

	tests[2].s.recv.buf.Init(1024)
	tests[2].s.recv.buf.Write(make([]byte, 200))
	tests[2].s.recv.flag <- flag_PSH

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := tt.s.Read(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Socket.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("Socket.Read() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
