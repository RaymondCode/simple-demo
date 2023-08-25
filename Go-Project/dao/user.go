package dao

import (
	"errors"

	"github.com/life-studied/douyin-simple/global"
	"gorm.io/gorm"
)

//查询Token是否存在
func QueryToken（Token string）bool{
	var tokens []Token
	result := global.DB.Where("token=?", token).First(&tokens)
	
	// 检查查询结果和错误
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			if result.RowsAffected == 0 {
				return true
			}
		} else {
			panic(result.Error)
		}
	}
	
	return false
}