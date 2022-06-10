package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/BaiZe1998/douyin-simple-demo/dto"

	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"gorm.io/gorm"
)

func IsFollow(ctx context.Context, userID int64, followedUser int64) (bool, bool) {
	var relation model.Follow

	if err := model.DB.Table("follow").WithContext(ctx).Where("user_id=? and followed_user=? ", userID, followedUser).First(&relation).Error; err == nil {
		if relation.Status == 1 {
			return true, true
		} else {
			return true, false
		}
	} else {
		return false, false
	}
}

func FollowCountAction(ctx context.Context, userID int64, followedUser int64, action_type int) error {
	if action_type == 1 {
		// 关注操作
		// user_id 关注数+1， to_user_id 被关注数+1
		model.DB.Table("user").WithContext(ctx).Where("id = ?", userID).Update("follow_count", gorm.Expr("follow_count+?", 1))
		model.DB.Table("user").WithContext(ctx).Where("id = ?", followedUser).Update("follower_count", gorm.Expr("follower_count+?", 1))
	} else {
		// 取关操作
		model.DB.Table("user").WithContext(ctx).Where("id = ?", userID).Update("follow_count", gorm.Expr("follow_count-?", 1))
		model.DB.Table("user").WithContext(ctx).Where("id = ?", followedUser).Update("follower_count", gorm.Expr("follower_count-?", 1))
	}

	return nil
}

func FollowAction(ctx context.Context, user_id int64, to_user_id int64, action_type int) error {
	is_exist, is_follow := IsFollow(ctx, user_id, to_user_id)

	if !is_exist {
		// 不存在的关系直接创建
		new_follow_relation := model.Follow{
			UserId:       int64(user_id),
			FollowedUser: int64(to_user_id),
			Status:       int(action_type),
		}

		err := model.CreateFollow(ctx, &new_follow_relation)

		if err == nil {
			FollowCountAction(ctx, user_id, to_user_id, action_type)
		} else {
			fmt.Println(err)
		}
	} else {
		// 存在的关系进行更新
		if (action_type == 1 && !is_follow) || (action_type == 2 && is_follow) {
			err := model.UpdateFollow(ctx, user_id, to_user_id, &action_type)

			if err == nil {
				FollowCountAction(ctx, user_id, to_user_id, action_type)
			} else {
				fmt.Println(err)
			}
		}
	}

	return nil
}

func GetFollowList(ctx context.Context, userId int64, actionType uint) ([]dto.User, error) {
	var followList []dto.User

	if actionType == 1 {
		// 操作类型为获取关注列表
		if err := model.DB.Table("user").WithContext(ctx).Joins("left join follow on user.id = follow.followed_user").
			Select("user.id", "user.name", "user.follow_count", "user.follower_count").
			Where("follow.user_id = ?", userId).Scan(&followList).Error; err != nil {
			return followList, nil
		} else {
			fmt.Println(err)
		}
	} else if actionType == 2 {
		// 操作类型为获取粉丝列表
		if err := model.DB.Table("user").WithContext(ctx).Joins("left join follow on user.id = follow.user_id").
			Select("user.id", "user.name", "user.follow_count", "user.follower_count").
			Where("follow.followed_user = ?", userId).Scan(&followList).Error; err != nil {
			return followList, nil
		}
	} else {
		return followList, errors.New("ambiguous actionType")
	}

	for i, n := range followList {
		isExist, isFollow := IsFollow(ctx, userId, n.Id)
		if isExist && isFollow {
			followList[i].IsFollow = true
		} else {
			followList[i].IsFollow = false
		}
	}

	return followList, nil
}
