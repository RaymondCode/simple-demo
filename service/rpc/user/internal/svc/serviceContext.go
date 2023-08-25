package svc

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"user/common"
	"user/internal/model"
)

type ServiceContext struct {
	Config common.Config
	DB     *gorm.DB
}

func NewServiceContext(c common.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(getDSN(&c)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "user_", // 表明前缀，可不设置
			SingularTable: true,    // 使用单数表名，即不会在表名后添加复数s
		},
	})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}

func getDSN(c *common.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		c.MYSQL.User,
		c.MYSQL.Password,
		c.MYSQL.Host,
		c.MYSQL.Port,
		c.MYSQL.Database,
	)
}
