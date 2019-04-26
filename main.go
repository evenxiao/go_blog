package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"go_blog/models"
	_ "go_blog/routers"
	_ "net/http/pprof"
	"runtime"

)

func init()  {
	//fmt.Println("wernwerr.........")
	logs.SetLogger("console")
	logs.SetLogger(logs.AdapterMultiFile, `{"filename":"test.log"}`)
	logs.SetLevel(beego.LevelDebug)
    models.InitConn()
}
func main() {
	fmt.Println("CPU 核数：",runtime.NumCPU())
	beego.Run()

}

