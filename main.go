package main

import (
	"github.com/lqf1215/kitex-demo/global"
	"github.com/lqf1215/kitex-demo/kitex_gen/service/userservice"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DSN = "root:lilishop@tcp(1.12.69.129:3306)/kitex_demo?charset=utf8&parseTime=True&loc=Local"
var MaxIdleConns = 5
var MaxOpenConns = 20

func GormMysql() *gorm.DB {
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Println("server gorm error:", err)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(MaxIdleConns)
		sqlDB.SetMaxOpenConns(MaxOpenConns)
		sqlDB.SetConnMaxLifetime(1000)
		return db
	}
}
func init() {
	db := GormMysql()
	if db == nil {
		panic("未连接数据库,请检查库或用户或密码等")
	}
	global.DB = db
}

func main() {

	svr := userservice.NewServer(new(UserServiceImpl))
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

	log.Println("server ...")
}
