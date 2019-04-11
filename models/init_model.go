package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func InitConn()  {
	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", beego.AppConfig.String("mysqluser"),beego.AppConfig.String("mysqlpass"),beego.AppConfig.String("mysqlurls"),beego.AppConfig.String("mysqldb"))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysqlDsn)
	orm.Debug = true
}

func GetOrm() orm.Ormer{
	return orm.NewOrm()

}
