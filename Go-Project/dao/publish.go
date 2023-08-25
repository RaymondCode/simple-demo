// Package dao -----------------------------
// @file      : publish.go
// @author    : Yunyin
// @contact   : yunyin_jayyi@qq.com
// @time      : 2023/8/24 17:21
// -------------------------------------------
package dao

import "github.com/life-studied/douyin-simple/global"

func SaveVideoToMysql(newVideo Video) error {
	err := global.DB.Create(&newVideo).Error
	if err != nil {
		return err
	}
	return nil
}
