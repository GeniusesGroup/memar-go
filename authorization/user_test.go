/* For license and copyright information please see LEGAL file in repository */

package authorization

import (
	"testing"

	er "../error"
)

func TestUserType_Check(t *testing.T) {
	type args struct {
		user UserType
	}
	tests := []struct {
		name    string
		ut      UserType
		args    args
		wantErr *er.Error
	}{
		{
			name: "test1",
			ut:   UserTypeAll,
			args: args{
				user: UserTypeApp,
			},
			wantErr: nil,
		}, {
			name: "test2",
			ut:   UserTypePerson,
			args: args{
				user: UserTypePerson,
			},
			wantErr: nil,
		}, {
			name: "test3",
			ut:   UserTypePerson,
			args: args{
				user: UserTypeApp,
			},
			wantErr: ErrUserNotAllow,
		}, {
			name: "test4",
			ut:   UserTypeUnset,
			args: args{
				user: UserTypePerson,
			},
			wantErr: ErrUserNotAllow,
		}, {
			name: "test5",
			ut:   UserTypeAll ^ UserTypeGuest,
			args: args{
				user: UserTypeGuest,
			},
			wantErr: ErrUserNotAllow,
		}, {
			name: "test6",
			ut:   ^UserTypeGuest,
			args: args{
				user: UserTypeGuest,
			},
			wantErr: ErrUserNotAllow,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := tt.ut.Check(tt.args.user)
			if !gotErr.Equal(tt.wantErr) {
				t.Errorf("UserType.Check() = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}

func TestUserType_CheckReverse(t *testing.T) {
	type args struct {
		users UserType
	}
	tests := []struct {
		name    string
		ut      UserType
		args    args
		wantErr *er.Error
	}{
		{
			name: "test1",
			ut:   UserTypePerson,
			args: args{
				users: UserTypePerson | UserTypeAI,
			},
			wantErr: nil,
		}, {
			name: "test2",
			ut:   UserTypeGuest,
			args: args{
				users: ^UserTypeGuest,
			},
			wantErr: ErrUserNotAllow,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := tt.ut.CheckReverse(tt.args.users)
			if !gotErr.Equal(tt.wantErr) {
				t.Errorf("UserType.CheckReverse() Base = %b, Checked by = %b", tt.ut, tt.args.users)
				t.Errorf("UserType.CheckReverse() = %v, want %v", gotErr, tt.wantErr)
			}
		})
	}
}
