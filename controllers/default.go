package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/utils/captcha"
)

type MainController struct {
	beego.Controller
}
var cpt *captcha.Captcha
func init()  {
	store := cache.NewMemoryCache()

	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.StdHeight = 60
	cpt.StdWidth = 120

}
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	webInfo := make(map[string]string)
	webInfo["title"] = "牛人网站后台管理系统"
	webInfo["webName"] = "牛人网络"
	c.Data["WebInfo"] = &webInfo
	c.TplName = "index.html"
}

func (c *MainController) Post() {
	c.TplName = "index.html"
	if cpt.VerifyReq(c.Ctx.Request) {
		c.Data["Success"] = "验证成功"
	}else{
		c.Data["Success"] = "验证失败"
	}

}

func (c *MainController) Welcome(){
	c.TplName = "welcome.html"
}
func (c *MainController) Login(){
	c.TplName = "login.html"
}