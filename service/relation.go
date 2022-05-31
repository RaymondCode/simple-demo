package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/repository"
	"gorm.io/gorm"
)

type RelationService struct {
}

//根据前台传来的信息开始关系操作
func (rs *RelationService) RelationAction(userId, toUserId int64, actionType string) error {
	switch actionType {
	case "1":
		err := global.DB.Transaction(func(tx *gorm.DB) error {
			//关注
			if err := repository.GroupApp.RelationRepository.InsertUserId(userId, toUserId); err != nil {
				return err
			}
			//增加己身关注数量
			if err := repository.GroupApp.RelationRepository.IncreaseFollowCount(userId, 1); err != nil {
				return err
			}
			//被关注者增加粉丝数量
			if err := repository.GroupApp.RelationRepository.IncreaseFollowerCount(toUserId, 1); err != nil {
				return err
			}
			return nil
		})
		return err
	case "2":
		err := global.DB.Transaction(func(tx *gorm.DB) error {
			//取消关注
			if err := repository.GroupApp.RelationRepository.DeleteUserId(userId, toUserId); err != nil {
				return err
			}
			if err := repository.GroupApp.RelationRepository.DecreaseFollowCount(userId, 1); err != nil {
				return err
			}
			if err := repository.GroupApp.RelationRepository.DecreaseFollowerCount(toUserId, 1); err != nil {
				return err
			}
			return nil
		})
		return err
	default:
		return errors.New("action_type is wrong")
	}
}

//查找关注列表
func (rr *RelationService) FollowList(userId int64) ([]*model.User, error) {
	userList, err := repository.GroupApp.RelationRepository.GetFollowListByUserId(userId)
	if err != nil {
		return nil, err
	}
	//
	for _, u := range userList {
		u.IsFollow = true
	}
	return userList, nil
}

//查找粉丝列表
func (rr *RelationService) FollowerList(userId int64) ([]*model.User, error) {
	userList, err := repository.GroupApp.RelationRepository.GetFollowerListByToUserId(userId)
	if err != nil {
		return nil, err
	}
	for _, u := range userList {
		f, _ := repository.GroupApp.RelationRepository.GetFollowerByUserIdAndToUserId(userId, u.Id)
		u.IsFollow = f != nil
	}
	return userList, nil
}
