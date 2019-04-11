package routers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"go_blog/controllers"
)

 func FilterUser(ctx *context.Context)  {
 	fmt.Println("用户session uid:", ctx.Input.Session("uid"))
	//uid := ctx.Input.Session("uid")
	uid := ctx.Input.Session("uid")
	uid_v := 0
	if uid != nil {
		uid_v = uid.(int)
	}

	logs.Info("用户id :", uid)
	//if uid_v <= 0 && ctx.Request.RequestURI != "/login" {
	if uid_v <= 0 && ctx.Request.RequestURI != "/login" {
		ctx.Redirect(302,"/login")
	}
}

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/api", &controllers.ApiController{})
    beego.Router("/welcome", &controllers.MainController{}, "get,post:Welcome")
    beego.Router("/login", &controllers.MainController{}, "get,post:Login")
    beego.Router("/loginOut", &controllers.MainController{}, "get,post:LoginOut")
    beego.Router("/test", &controllers.MainController{}, "get,post:Test")

    beego.InsertFilter("/*", beego.BeforeRouter,FilterUser)
}
