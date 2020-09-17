package mysql

import (
	"github.com/jinzhu/gorm"
)

func getDBPool() *ConnectPool {
	return GetConnect()
}

func getReadConn() *gorm.DB {
	return getDBPool().ReadDB
}

func getWriteConn() *gorm.DB {
	return getDBPool().WriteDB
}
