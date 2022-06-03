package service

import (
	"errors"
	"github.com/warthecatalyst/douyin/api"
	"github.com/warthecatalyst/douyin/dao"
)

func FollowAction(userId, toUserId int64, actTyp int) error {
	return NewFollowActionFlow(userId, toUserId, actTyp).Do()
}

func NewFollowActionFlow(userId, toUserId int64, actTyp int) *FollowActionFlow {
	return &FollowActionFlow{
		FromUserId: userId,
		ToUserId:   toUserId,
		ActionType: actTyp,
	}
}

type FollowActionFlow struct {
	FromUserId int64
	ToUserId   int64
	ActionType int
}

func (f *FollowActionFlow) Do() error {
	if err := f.checkParam(); err != nil {
		return err
	}
	return f.followImpl()
}

func (f *FollowActionFlow) checkParam() error {
	if f.ActionType != api.FollowAction && f.ActionType != api.UnfollowAction {
		return errors.New("未知关注相关操作类型！")
	}

	return nil
}

// TODO 并发：1.隔离级别，2.3处修改并发提高效率
func (f *FollowActionFlow) followImpl() error {
	tx := dao.GetTx()
	if f.ActionType == api.FollowAction {
		if err := dao.NewFollowDaoInstance().AddFollow(tx, f.FromUserId, f.ToUserId); err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if err := dao.NewFollowDaoInstance().DeleteFollow(tx, f.FromUserId, f.ToUserId); err != nil {
			tx.Rollback()
			return err
		}
	}
	user, err := dao.NewUserDaoInstance().GetUserById(f.FromUserId)
	if err != nil {
		tx.Rollback()
		return err
	}
	toUser, err := dao.NewUserDaoInstance().GetUserById(f.ToUserId)
	if err != nil {
		tx.Rollback()
		return err
	}
	if f.ActionType == api.FollowAction {
		user.FollowCount++
		toUser.FollowerCount++
	} else {
		user.FollowCount--
		toUser.FollowerCount--
	}
	if err := dao.NewUserDaoInstance().UpdateUser(user); err != nil {
		tx.Rollback()
		return err
	}
	if err := dao.NewUserDaoInstance().UpdateUser(toUser); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func GetFollowList(userId int64) ([]*api.User, error) {
	var users []*api.User
	relationList, err := dao.NewFollowDaoInstance().GetFollowList(userId)
	if err != nil {
		return users, err
	}
	for _, relation := range relationList {
		user, err := dao.NewUserDaoInstance().GetUserById(relation.ToUserID)
		if err != nil {
			return users, err
		}
		users = append(users, &api.User{
			Id:            user.UserID,
			Name:          user.UserName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      true,
		})
	}

	return users, nil
}

func GetFollowerList(userId int64) ([]*api.User, error) {
	var users []*api.User
	relationListTmp, err := dao.NewFollowDaoInstance().GetFollowList(userId)
	if err != nil {
		return users, err
	}
	followUserIdList := make(map[int64]bool)
	for _, relation := range relationListTmp {
		followUserIdList[relation.ToUserID] = true
	}
	relationList, err := dao.NewFollowDaoInstance().GetFollowerList(userId)
	if err != nil {
		return users, err
	}
	for _, relation := range relationList {
		user, err := dao.NewUserDaoInstance().GetUserById(relation.FromUserID)
		if err != nil {
			return users, err
		}
		users = append(users, &api.User{
			Id:            user.UserID,
			Name:          user.UserName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      followUserIdList[user.UserID],
		})
	}

	return users, nil
}
