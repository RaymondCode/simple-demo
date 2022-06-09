package service

import (
	"context"

	"github.com/BaiZe1998/douyin-simple-demo/db/model"
	"gorm.io/gorm"
)

func IsFollow(ctx context.Context, userID uint, followedUser uint) (bool, bool) {
	var relation model.Follow

	if err := model.DB.Table("follow").WithContext(ctx).Where("user_id = ? and followed_user = ? ", userID, followedUser).First(relation, 1).Error; err == nil {
		if relation.Status == 1 {
			return true, true
		} else {
			return true, false
		}
	} else {
		return false, false
	}
}

func FollowCountAction(ctx context.Context, userID uint, followedUser uint, action_type int) error {
	if action_type == 1 {
		// 关注操作
		// user_id 关注数+1， to_user_id 被关注数+1
		model.DB.Table("user").WithContext(ctx).Where("id = ?", userID).Update("FollowCount", gorm.Expr("FollowCount+?", 1))
		model.DB.Table("user").WithContext(ctx).Where("id = ?", followedUser).Update("FollowerCount", gorm.Expr("FollowerCount+?", 1))
	} else {
		// 取关操作
		model.DB.Table("user").WithContext(ctx).Where("id = ?", userID).Update("FollowCount", gorm.Expr("FollowCount+?", -1))
		model.DB.Table("user").WithContext(ctx).Where("id = ?", followedUser).Update("FollowerCount", gorm.Expr("FollowerCount+?", -1))
	}

	return nil
}

func FollowAction(ctx context.Context, user_id uint, to_user_id uint, action_type int) error {
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
			FollowAction(ctx, user_id, to_user_id, action_type)
		} else {
			// todo: 写日志
		}
	} else {
		// 存在的关系进行更新
		if (action_type == 1 && !is_follow) || (action_type == 2 && is_follow) {
			err := model.UpdateFollow(ctx, user_id, to_user_id, &action_type)

			if err == nil {
				FollowCountAction(ctx, user_id, to_user_id, action_type)
			} else {
				// todo: 写日志
			}
		}
	}

	return nil
}
