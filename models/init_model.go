package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)
//var DORM interface{}
func init(){

}
func InitConn()  {
	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", beego.AppConfig.String("mysqluser"),beego.AppConfig.String("mysqlpass"),beego.AppConfig.String("mysqlurls"),beego.AppConfig.String("mysqldb"))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", mysqlDsn)
	orm.Debug = true

	f, _ := os.Create("sql_test.log")
	orm.DebugLog = orm.NewLog(f)

	//sql_log , err := ioutil.ReadAll(os.Stdout)
	//if err != nil {
	//	logs.Error(err)
	//}
	//ioutil.WriteFile("sql_test.log", sql_log, 0644)
}
