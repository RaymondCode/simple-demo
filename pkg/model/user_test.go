package model_test

import (
	"reflect"
	"testing"

	"github.com/fitenne/youthcampus-dousheng/internal/common/settings"
	"github.com/fitenne/youthcampus-dousheng/internal/repository"
	"github.com/fitenne/youthcampus-dousheng/pkg/model"
	"github.com/spf13/viper"
)

func init() {
	if err := settings.Init("../../config.yaml"); err != nil {
		panic(err.Error())
	}

	repository.Init(repository.DBConfig{
		Driver:   viper.GetString("db.driver"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.pass"),
		DBname:   viper.GetString("db.database"),
	})
}

func TestQueryUserByID(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name string
		args args
		want *model.User
	}{
		{
			name: "TestUser",
			args: args{
				id: 1,
			},
			want: &model.User{
				ID:            1,
				Name:          "TestUser",
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
