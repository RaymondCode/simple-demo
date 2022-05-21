package repository

import (
	"errors"
	"fmt"
	"github.com/fitenne/youthcampus-dousheng/pkg/model"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Driver         string
	Host, Port     string
	User, Password string
	DBname         string
}

type MysqlProdiver struct {
	db          *gorm.DB
	connectOnce sync.Once
}

type DBProvider interface {
	Connect(c DBConfig) error
	GetDB() *gorm.DB
}

var dbProvider DBProvider
var initOnce sync.Once

// 连接到 DBConfig 制定的数据库，忽略 DBConfig 中的 Driver 字段
func (p *MysqlProdiver) Connect(c DBConfig) error {
	err := errors.New("already connected")
	p.connectOnce.Do(func() {
		template := "%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local"
		dsn := fmt.Sprintf(template, c.User, c.Password, c.Host, c.Port, c.DBname)
		dialector := mysql.New(mysql.Config{
			DriverName: "mysql",
			DSN:        dsn,
		})
		p.db, err = gorm.Open(dialector)
	})
	return err
}

func (p *MysqlProdiver) GetDB() *gorm.DB {
	return p.db
}

// 初始化数据库，只有第一次调用有效
func Init(c DBConfig) error {
	err := errors.New("Init called twice")
	initOnce.Do(func() {
		switch c.Driver {
		case "mysql":
			dbProvider = &MysqlProdiver{}
		default:
			err = errors.New("db driver not supported")
		}
		err = dbProvider.Connect(c)
		if err != nil {
			return
		}
		//创建表video
		var db = dbProvider.GetDB()
		if !db.Migrator().HasTable(&model.Video{}) {
			if err := db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(model.Video{}).Error; err != nil {
				panic(err)
			}
		}
		err = nil
	})
	return err
}
