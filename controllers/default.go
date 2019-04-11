package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils/captcha"
	"go_blog/models"
	"go_blog/utils"
	"time"
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
	var u  models.User
	u.Reg_time = time.Now().Format("2006-01-02 15:04:05")
	u.User_name = "肖文问"
	u.User_phone = "13761463100"
	u.Pwd = utils.Md5Encode("123456")
	models.AddUser(&u)
	logs.Info("user last id:", u.Id)

	//设置session数据
	user := make(map[string]string)
	user["name"]  = "xiaowen"
	user["phone"]  = "13761463130"
	c.SetSession("user", user)

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	webInfo := make(map[string]string)
	webInfo["title"] = "牛人网站后台管理系统"
	webInfo["webName"] = "牛人网络"
	c.Data["WebInfo"] = &webInfo
	c.Data["s_user"] = &user
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
	c.Data["s_user"] = c.GetSession("user")

	c.TplName = "welcome.html"
}
func (c *MainController) Login(){

	if c.Ctx.Input.IsPost() {
		fmt.Println("1111")
		logs.Info("1111")
		c.SetSession("uid", 2)
		c.Data["s_user"] = c.GetSession("user")
		logs.Info("session uid :" , c.GetSession("uid"))
		//var myStruct map[string]string
		//myStruct = make(map[string]string)
		//myStruct["name"] = "awen"
		//myStruct["age"] = "27"
		//myStruct["sex"] = "man"
		//c.Data["json"] = &myStruct
		//c.ServeJSON()
		c.Redirect("/", 302)
	}
	c.TplName = "login.html"
}

func (c *MainController) LoginOut(){

	c.SetSession("uid", 0)

	c.Redirect("/login", 302)

}
func (c *MainController) Test(){
    var u models.User
    u.Id = 1
	models.FindById(&u)

    //fmt.Println(u)

	c.Ctx.WriteString("hello world")
	c.Ctx.WriteString("hello world " + u.User_phone+ "---" +u.User_name)

    //var u2 models.User
    u2 , err:= models.FindOneUser(2)
	if err != nil {
		 panic(err)
	}
    c.Ctx.WriteString("\n\r" + u2.User_phone+ "---" +u2.User_name)
}