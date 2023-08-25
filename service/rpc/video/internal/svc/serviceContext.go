package svc

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"tiktok_startup/common/model"
	"tiktok_startup/service/rpc/video/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Mysql  *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Mysql:  initMysql(c),
	}
}
func initMysql(c config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Mysql.Username,
		c.Mysql.Password,
		c.Mysql.Address,
		c.Mysql.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.Mysql.TablePrefix, // 表名前缀
			SingularTable: true,                // 使用单数表名
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&model.Video{}, &model.Favorite{}, &model.Comment{})
	if err != nil {
		panic(err)
	}
	return db
}
