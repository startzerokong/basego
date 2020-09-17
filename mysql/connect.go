package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/startzerokong/basego/util"
)

type ConnectPool struct {
	WriteDB *gorm.DB
	ReadDB *gorm.DB
}

var ConnPool *ConnectPool

func GetConnect() *ConnectPool {
	if ConnPool == nil {
		InitMysql()
	}
	return ConnPool
}

func InitMysql()  {
	ConnPool = &ConnectPool{}

	read, err1 := ReadDb()
	if err1 != nil {

	}
	ConnPool.ReadDB = read

	write, err2 := WriteDb()
	if err2 != nil {

	}
	ConnPool.WriteDB = write
}

func ReadDb() (*gorm.DB, error) {
	conf := util.GetDbConfig()
	connect := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Password, conf.Hostname, conf.Dbname)
	return gorm.Open("mysql",connect)
}

func WriteDb() (*gorm.DB, error) {
	conf := util.GetDbConfig()
	connect := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Password, conf.Hostname, conf.Dbname)
	return gorm.Open("mysql",connect)
}
