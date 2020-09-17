package mysql

import (
	"github.com/jinzhu/gorm"
)

func GetDBPool() *ConnectPool {
	return GetConnect()
}

func GetReadConn() *gorm.DB {
	return GetDBPool().ReadDB
}

func GetWriteConn() *gorm.DB {
	return GetDBPool().WriteDB
}
