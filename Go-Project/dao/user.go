package dao

//查询用户名是否存在
func QueryName(username string) bool {
	//使用gorm查询用户名是否存在

	return true
}

//将用户信息存入数据库
func AddUserInfo(id int64, username string, password string, token string) error {
	//先查询数据库中token列是否存在，不存在，创建token列
	// 检查 token 列是否存在
	//检查库中的表列名
	// result := db.Exec(`SELECT * FROM user LIMIT 1`)
	// if result.Error != nil {
	// 	return result.Error
	// }
	// //获取查询结果的列名
	// columns, err := result.Rows().Columns()
	// if err != nil {
	// 	return err
	// }
	// //历遍来检查是否存在 "token" 列
	// var tokenColumnExists bool
	// for _, column := range columns {
	// 	if column == "token" {
	// 		tokenColumnExists = true
	// 		break
	// 	}
	// }
	// // 创建 token 列
	// if !tokenColumnExists {
	// 	db.Exec(`ALTER TABLE user ADD COLUMN token VARCHAR(255)`)
	// 	if result.Error != nil {
	// 		return result.Error
	// 	}
	// }
	//将id username password token存入数据库中
	return nil
}

//获取所有数据
func GetAllUsers() ([]User, error) {
	//数据库连接
	var users []User
	// result := db.Find(&users)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }

	return users, nil

}
