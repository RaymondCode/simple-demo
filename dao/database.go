package dao

import (
	"fmt"
	"github.com/warthecatalyst/douyin/config"
	"github.com/warthecatalyst/douyin/logx"
	"github.com/warthecatalyst/douyin/model"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

func GetTx() *gorm.DB {
	return db.Begin()
}

func InitDB() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DbUser,
		config.DbPassWord,
		config.DbHost,
		config.DbPort,
		config.DbName)

	var err error
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, //关闭外键！！！
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false,       //默认在表的后面加s
			TablePrefix:   "t_douyin_", // 表名前缀
		},
		SkipDefaultTransaction: true, // 禁用默认事务
	})

	if err != nil {
		logx.DyLogger.Panicf("数据库连接失败，错误：%s", err)
	}

	err = db.AutoMigrate(&model.Video{}, &model.User{}, &model.Follow{}, &model.Comment{}, &model.Favourite{}) //TODO 数据库自动迁移
	if err != nil {
		logx.DyLogger.Panicf("数据库自动迁移失败，错误：%s", err)
	}
	sqlDb, _ := db.DB()

	// TODO 这方面后期再说吧， 参数到底设为多少
	sqlDb.SetMaxIdleConns(50)                  // 连接池中的最大闲置连接数
	sqlDb.SetMaxOpenConns(100)                 // 数据库的最大连接数量
	sqlDb.SetConnMaxLifetime(10 * time.Second) // 连接的最大可复用时间
}
