package initialize

import (
	"github.com/life-studied/douyin-simple/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Mysql() {

	mysqlConfig := mysql.Config{
		DSN:                       global.CONFIG.Mysql.Dsn(), // DSN data source name
		SkipInitializeWithVersion: false,                     // 根据版本自动配置
	}
	var err error
	global.DB, err = gorm.Open(mysql.New(mysqlConfig),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			Logger:                 logger.Default.LogMode(logger.Info),
		},
	) //启动MySQL，启用预编译SQL语句，不会自动开启事务等功能
	if err != nil {
		panic(err)
	} else {
		sqlDB, _ := global.DB.DB()
		sqlDB.SetMaxIdleConns(global.CONFIG.Mysql.MaxIdleConns) // 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxOpenConns(global.CONFIG.Mysql.MaxOpenConns) // 设置打开数据库连接的最大数量
	}
}
