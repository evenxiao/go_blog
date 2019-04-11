package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go_blog/models"
	_ "go_blog/routers"
)

func init()  {
	//fmt.Println("wernwerr.........")
	logs.SetLogger("console")
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"test.log"}`)
    models.InitConn()
}
func main() {
	beego.Run()

	log := logs.NewLogger(100000)
	log.Emergency("Emergency")
	log.Alert("Alert")
	log.Critical("Critical")
	log.Error("Error")
	log.Warning("Warning")
	log.Notice("Notice")
	log.Informational("Informational")
	log.Debug("Debug")

	log.Flush() // 将日志从缓冲区读出，写入到文件
	log.Close()
}

