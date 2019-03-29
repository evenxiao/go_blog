package routers

import (
	"go_blog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/api", &controllers.ApiController{})
    beego.Router("/welcome", &controllers.MainController{}, "get,post:Welcome")
    beego.Router("/login", &controllers.MainController{}, "get,post:Login")
}
