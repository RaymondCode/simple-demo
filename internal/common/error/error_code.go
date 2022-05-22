package errorcode

/**
 * @create 2020-03-24
 * @description 错误码
 */
type ErrorCode int

const (
	// 成功
	Success ErrorCode = iota
	// 未知错误
	UnknownError
	// 数据库错误
	DatabaseError
	// 数据库操作错误
	DatabaseOperationError
	// 数据库查询错误
	DatabaseQueryError
	// 数据库更新错误
	DatabaseUpdateError
	// 数据库删除错误
	DatabaseDeleteError
	// 数据库插入错误
	DatabaseInsertError
	// 数据库查询结果为空
	DatabaseQueryResultEmpty
	// 数据库查询结果不唯一
	DatabaseQueryResultNotUnique
	// 数据库查询结果不唯一
	DatabaseQueryResultNotExist
	// 数据库查询结果不唯一
	DatabaseQueryResultNotExistOrNotUnique
	// 数据库查询结果不唯一
	DatabaseQueryResultNotExistOrNotUniqueOrEmpty
	// 数据库查询结果不唯一
	DatabaseQueryResultNotExistOrNotUniqueOrEmptyOrNotUnique
	// 数据库查询结果不唯一
	DatabaseQueryResultNotExistOrNotUniqueOrEmptyOrNotUniqueOrNotExist
	// 数据库查询结果不唯一
	DatabaseQueryResultNotExistOrNotUniqueOrEmptyOrNotUniqueOrNotExistOrNotUnique

	//video
	//create video 外键不统一
	VideoCreateForeignKeyNotUnified
	//create video 外键不存在
	VideoCreateForeignKeyNotExist
)

// 错误码对应的错误信息
func (e ErrorCode) Message() string {
	switch e {
	case Success:
		return "success"
	case UnknownError:
		return "unknown error"
	case DatabaseError:
		return "database error"
	case DatabaseOperationError:
		return "database operation error"
	case DatabaseQueryError:
		return "database query error"
	case DatabaseUpdateError:
		return "database update error"
	case DatabaseDeleteError:
		return "database delete error"
	case DatabaseInsertError:
		return "database insert error"
	case DatabaseQueryResultEmpty:
		return "database query result empty"
	case DatabaseQueryResultNotUnique:
		return "database query result not unique"
	case DatabaseQueryResultNotExist:
		return "database query result not exist"
	case DatabaseQueryResultNotExistOrNotUnique:
		return "database query result not exist or not unique"
	case DatabaseQueryResultNotExistOrNotUniqueOrEmpty:
		return "database query result not exist or not unique or empty"
	case DatabaseQueryResultNotExistOrNotUniqueOrEmptyOrNotUnique:
		return "database query result not exist or not unique or empty or not unique"
	case DatabaseQueryResultNotExistOrNotUniqueOrEmptyOrNotUniqueOrNotExist:
		return "database query result not exist or not unique or empty or not unique or not exist"
	case DatabaseQueryResultNotExistOrNotUniqueOrEmptyOrNotUniqueOrNotExistOrNotUnique:
		return "database query result not exist or not unique or empty or not unique or not exist or not unique"
	case VideoCreateForeignKeyNotUnified:
		return "video create foreign key not unified"
	case VideoCreateForeignKeyNotExist:
		return "video create foreign key not exist"
	default:
		return "unknown error"
	}
}
