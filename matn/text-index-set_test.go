/* For license and copyright information please see LEGAL file in repository */

package matn

import (
	"fmt"
	"testing"
)

func TestTextWordTokenization(t *testing.T) {
	type args struct {
		req *TextIndexSetReq
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				req: &TextIndexSetReq{
					Text: "test is best",
				},
			},
		},
		{
			name: "test2",
			args: args{
				req: &TextIndexSetReq{
					Text: "مجتمع سبز - گلستان شیراز",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("WordTokenization() for", tt.name)
			indexes := WordTokenization(tt.args.req)
			for _, index := range indexes {
				fmt.Println(index.Word)
			}
		})
	}
}
