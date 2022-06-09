package dao

import (
	"simple-demo/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB
var UserDao *userDao

func InitMysql() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:admin#123456@tcp(127.0.0.1:3306)/simple_demo?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)           //连接池中最大的空闲连接数
	sqlDB.SetMaxOpenConns(100)          //连接池最多容纳连接数
	sqlDB.SetConnMaxLifetime(time.Hour) //连接池中连接的最大可复用时间
	//db.AutoMigrate(&User{})

	if err != nil {
		fmt.Println("gorm open failed:", err.Error())
	} else {
		fmt.Println("gorm open succeed!")
	}
	fmt.Println(db, err)
	return db
}

func GetDB() *gorm.DB {
	db = InitMysql()
	return db
}

type userDao struct {
	gm *gorm.DB
}

func (dao *userDao) Create(tx *gorm.DB, data *model.User) (rowsAffected int64, err error) {
	db := tx.Create(data)
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *userDao) Update(tx *gorm.DB, id uint, data map[string]interface{}) (rowsAffected int64, err error) {
	db := tx.Model(&model.User{}).Where("id = ?", id).Updates(data)
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

func (dao *userDao) Delete(tx *gorm.DB, data []int) (rowsAffected int64, err error) {
	db := tx.Where("id in (?)", data).Delete(&model.User{})
	if err = db.Error; db.Error != nil {
		return
	}
	rowsAffected = db.RowsAffected
	return
}

// FindAll 查询全部
func (dao *userDao) FindAll() (list []model.UserFind, err error) {
	db := dao.gm.Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *userDao) FindAllWhere(query interface{}, args ...interface{}) (list []model.UserFind, err error) {
	db := dao.gm.Where(query, args...).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *userDao) FindOneWhere(query interface{}, args ...interface{}) (record model.User, err error) {
	db := dao.gm.Where(query, args...).First(&record)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *userDao) FindCountWhere(query interface{}, args ...interface{}) (count int64, err error) {
	db := dao.gm.Model(&model.User{}).Where(query, args...).Count(&count)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *userDao) FindCount() (count int64, err error) {
	db := dao.gm.Model(&model.User{}).Count(&count)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *userDao) Raw(sqlStr string, params ...interface{}) (list []model.UserFind, err error) {
	db := dao.gm.Debug().Raw(sqlStr, params...).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

// WhereQuery 按条件查询
func (dao *userDao) WhereQuery(query interface{}, args ...interface{}) *userDao {
	return &userDao{
		dao.gm.Where(query, args...),
	}

}

func (dao *userDao) WhereUserNameLike(username string) *userDao {
	return &userDao{
		dao.gm.Where("username like ?", "%"+username+"%"),
	}
}

func (dao *userDao) WhereDisabled(isDisabled int) *userDao {
	return &userDao{
		dao.gm.Where("is_disabled = ?", isDisabled),
	}
}

// Paginate 分页查询
func (dao *userDao) Paginate(offset, limit int) (count int64, list []model.UserFind, err error) {
	db := dao.gm.Model(&model.UserFind{}).Count(&count).Offset(offset).Limit(limit).Find(&list)
	if err = db.Error; db.Error != nil {
		return
	}
	return
}

func (dao *userDao) Debug() *userDao {
	return &userDao{
		dao.gm.Debug(),
	}
}

func (dao *userDao) Offset(offset int) *userDao {
	return &userDao{
		dao.gm.Offset(offset),
	}
}

func (dao *userDao) Limit(limit int) *userDao {
	return &userDao{
		dao.gm.Limit(limit),
	}
}

func (dao *userDao) OrderBy(sortFlag, sortOrder string) *userDao {
	return &userDao{
		dao.gm.Order(sortFlag + " " + sortOrder),
	}
}

// Joins 关联查询
func (dao *userDao) Joins(query string, args ...interface{}) *userDao {
	return &userDao{
		dao.gm.Joins(query, args),
	}
}

// Preloads 预加载
func (dao *userDao) Preloads(query string) *userDao {
	return &userDao{
		dao.gm.Preload(query),
	}
}