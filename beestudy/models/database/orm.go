package database

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", getDbUrl())
}

func getDbUrl() string {
	dbUser := beego.AppConfig.String("beestudy.dbuser")
	dbPwd := beego.AppConfig.String("beestudy.dbpwd")
	dbHost := beego.AppConfig.String("beestudy.dbhost")
	dbPort := beego.AppConfig.String("beestudy.dbport")
	dbName := beego.AppConfig.String("beestudy.dbname")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", dbUser, dbPwd, dbHost, dbPort, dbName)
}
