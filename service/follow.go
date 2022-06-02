package service

import "github.com/warthecatalyst/douyin/dao"

const (
	followAct   = "follow"
	unfollowAct = "unfollow"
)

// TODO 并发：1.隔离级别，2.3处修改并发提高效率
func followImpl(userId, toUserId int64, actTyp string) error {
	tx := dao.GetTx()
	if actTyp == followAct {
		if err := dao.AddFollow(tx, userId, toUserId); err != nil {
			tx.Rollback()
			return err
		}
	} else {
		if err := dao.DeleteFollow(tx, userId, toUserId); err != nil {
			tx.Rollback()
			return err
		}
	}
	user, err := dao.GetUserById(userId)
	if err != nil {
		tx.Rollback()
		return err
	}
	toUser, err := dao.GetUserById(toUserId)
	if err != nil {
		tx.Rollback()
		return err
	}
	if actTyp == followAct {
		user.FollowCount++
		toUser.FollowerCount++
	} else {
		user.FollowCount--
		toUser.FollowerCount--
	}
	if err := dao.UpdateUser(user); err != nil {
		tx.Rollback()
		return err
	}
	if err := dao.UpdateUser(toUser); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func Follow(userId, toUserId int64) error {
	return followImpl(userId, toUserId, followAct)
}

func UnFollow(userId, toUserId int64) error {
	return followImpl(userId, toUserId, unfollowAct)
}

func GetFollowList(userId int64) ([]*User, error) {
	var users []*User
	relationList, err := dao.GetFollowList(userId)
	if err != nil {
		return users, err
	}
	for _, relation := range relationList {
		user, err := dao.GetUserById(relation.ToUserID)
		if err != nil {
			return users, err
		}
		users = append(users, &User{
			Id:            user.UserID,
			Name:          user.UserName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      true,
		})
	}

	return users, nil
}

func GetFollowerList(userId int64) ([]*User, error) {
	var users []*User
	relationListTmp, err := dao.GetFollowList(userId)
	if err != nil {
		return users, err
	}
	followUserIdList := make(map[int64]bool)
	for _, relation := range relationListTmp {
		followUserIdList[relation.ToUserID] = true
	}
	relationList, err := dao.GetFollowerList(userId)
	if err != nil {
		return users, err
	}
	for _, relation := range relationList {
		user, err := dao.GetUserById(relation.FromUserID)
		if err != nil {
			return users, err
		}
		users = append(users, &User{
			Id:            user.UserID,
			Name:          user.UserName,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      followUserIdList[user.UserID],
		})
	}

	return users, nil
}
