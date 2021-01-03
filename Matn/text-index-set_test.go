/* For license and copyright information please see LEGAL file in repository */

package matn

import (
	"reflect"
	"testing"

	er "../error"
)

func TestTextIndexSet(t *testing.T) {
	type args struct {
		req *TextIndexSetReq
	}
	tests := []struct {
		name    string
		args    args
		wantErr *er.Error
	}{
		{
			name: "test1",
			args: args{
				req: &TextIndexSetReq{
					Text: "test is best",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotErr := TextIndexSet(tt.args.req); !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("TextIndexSet() = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
