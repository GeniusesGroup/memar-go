/* For license and copyright information please see LEGAL file in repository */

package http

import (
	"reflect"
	"testing"
)

func Test_getContentType(t *testing.T) {
	type args struct {
		contentType string
	}
	tests := []struct {
		name  string
		args  args
		wantC ContentType
	}{
		{
			name: "test1",
			args: args{
				contentType: "text/html; charset=utf-8",
			},
			wantC: ContentType{
				Type:    1,
				SubType: 1,
				Charset: "utf-8",
			},
		},
		{
			name: "test2",
			args: args{
				contentType: "application/json; charset=utf-8",
			},
			wantC: ContentType{
				Type:    2,
				SubType: 2,
				Charset: "utf-8",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := getContentType(tt.args.contentType); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("getContentType() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}
