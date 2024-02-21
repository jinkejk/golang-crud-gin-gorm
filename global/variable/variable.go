package variable

import (
	"gorm.io/gorm"
)

var (
	// gorm 数据库客户端
	GormDbMysql *gorm.DB // 全局gorm的客户端连接
)
