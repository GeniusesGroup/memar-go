/* For license and copyright information please see LEGAL file in repository */

package user

import (
	"testing"

	"../protocol"
)

func Test_CheckUserType(t *testing.T) {
	type args struct {
		baseUsers protocol.UserType
		user      protocol.UserType
	}
	tests := []struct {
		name     string
		args     args
		WantBool bool
	}{
		{
			name: "test1",
			args: args{
				baseUsers: protocol.UserType_All,
				user:      protocol.UserType_App,
			},
			WantBool: true,
		}, {
			name: "test2",
			args: args{
				baseUsers: protocol.UserType_Person,
				user:      protocol.UserType_Person,
			},
			WantBool: true,
		}, {
			name: "test3",
			args: args{
				baseUsers: protocol.UserType_Person,
				user:      protocol.UserType_App,
			},
			WantBool: false,
		}, {
			name: "test4",
			args: args{
				baseUsers: protocol.UserType_Unset,
				user:      protocol.UserType_Person,
			},
			WantBool: false,
		}, {
			name: "test5",
			args: args{
				baseUsers: protocol.UserType_All ^ protocol.UserType_Guest,
				user:      protocol.UserType_Guest,
			},
			WantBool: false,
		}, {
			name: "test6",
			args: args{
				baseUsers: ^protocol.UserType_Guest,
				user:      protocol.UserType_Guest,
			},
			WantBool: false,
		}, {
			name: "test7",
			args: args{
				baseUsers: protocol.UserType_Person | protocol.UserType_Thing,
				user:      protocol.UserType_Person,
			},
			WantBool: true,
		}, 
		// TODO:::
		// {
			// name: "test8",
			// args: args{
				// baseUsers: ^protocol.UserType_Guest, 
				// user:      protocol.UserType_Unset,
			// },
			// WantBool: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var exist = CheckUserType(tt.args.baseUsers, tt.args.user)
			if exist != tt.WantBool {
				t.Errorf("CheckUserType() Base = %b, Checked by = %b", tt.args.baseUsers, tt.args.user)
				t.Errorf("CheckUserType() on %v return %v, want %v", tt.name, exist, tt.WantBool)
			}
		})
	}
}
