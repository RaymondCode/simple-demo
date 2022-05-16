package main_test

import (
	"reflect"
	"testing"

	"github.com/fitenne/youthcampus-dousheng/internal/repository"
	"github.com/fitenne/youthcampus-dousheng/pkg/model"
)

func Test_userCtl_QueryUserByID(t *testing.T) {
	type args struct {
		id uint
	}
	tests := []struct {
		name string
		args args
		want *model.User
	}{
		{
			name: "alice",
			args: args{
				id: 1,
			},
			want: &model.User{
				ID:            1,
				UserName:      "alice",
				FollowCount:   0,
				FollowerCount: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctl := repository.GetUserCtl()
			if got := ctl.QueryUserByID(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userCtl.QueryUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
