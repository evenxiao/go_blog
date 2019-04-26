package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils/captcha"
	"go_blog/models"
	"go_blog/services"
	"os"
	"strconv"
	"strings"
)
var JsonV JsonContent
type MainController struct {
	beego.Controller
}
var cpt *captcha.Captcha
func init()  {

	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store)
	cpt.StdHeight = 50
	cpt.StdWidth = 120
	cpt.ChallengeNums = 4
}
func (c *MainController) Get() {
	//var u  models.User
	//u.Reg_time = time.Now().Format("2006-01-02 15:04:05")
	//u.User_name = "肖文问"
	//u.User_phone = "13761463100"
	//u.Pwd = utils.Md5Encode("123456")
	//models.AddUser(&u)
	//logs.Info("user last id:", u.Id)

	//设置session数据
	user := make(map[string]string)
	user["name"]  = "xiaowen"
	user["phone"]  = "13761463130"
	c.SetSession("user", user)

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	webInfo := make(map[string]string)
	webInfo["title"] = "公司后台管理系统"
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
		c.Data["s_user"] = c.GetSession("user")
		userName := strings.TrimSpace(c.GetString("username"))
		pwd := strings.TrimSpace(c.GetString("password"))
		//captcha := strings.TrimSpace(c.GetString("captcha"))
		Is_pass_cpt := cpt.VerifyReq(c.Ctx.Request)
		fmt.Println("验证码是否通过：", Is_pass_cpt)
		//if !cpt.VerifyReq(c.Ctx.Request){
		if !Is_pass_cpt{
			JsonV.EchoJson(0,"验证码错误", "")
			c.Data["json"] = &JsonV
			c.ServeJSON()
			return
		}else if userName == "" && pwd == "" {
			JsonV.EchoJson(0,"用户名、密码必填", "")
			c.Data["json"] = &JsonV
			c.ServeJSON()
			return
		}
		DealerM := models.Dealer{}
		user_have := DealerM.FindByUsernameAndPwd(userName,pwd)
		if user_have {
			c.SetSession("uid", 2)
			logs.Info("session uid :" , c.GetSession("uid"))
			JsonV.EchoJson(1,"登陆成功", "")
			c.Data["json"] = &JsonV
			c.ServeJSON()
			return
		}else{
			JsonV.EchoJson(0,"用户名或密码错误", "")
			c.Data["json"] = &JsonV
			c.ServeJSON()
			return
		}
		os.Exit(-1)
	}

	c.Data["s_user"] = c.GetSession("user")
	c.TplName = "login.html"

}

func (c *MainController) LoginOut(){

	c.SetSession("uid", 0)

	c.Redirect("/login", 302)

}
func (c *MainController) Test(){
    //var u models.User
    //u.Id = 1
	//models.FindById(&u)
	//
    ////fmt.Println(u)

	c.Ctx.WriteString("hello world")
	//c.Ctx.WriteString("hello world " + u.User_phone+ "---" +u.User_name)

    //var u2 models.User
    //u2 , err:= models.FindOneUser(2)
	//if err != nil {
	//	 panic(err)
	//}
    //c.Ctx.WriteString("\n\r" + u2.User_phone+ "---" +u2.User_name)
}
func (c *MainController) GetUserInfo()  {
	id := c.Input().Get("id")

	id_v, _:= strconv.Atoi(id)
	content := services.HandleUser(id_v)

	//c.Ctx.WriteString("user id by url :" + id + "\n\r 用户信息：" + content)
	JsonV.EchoJson(1,"数据返回成功", content)
	c.Data["json"] = &JsonV
	c.ServeJSON()
}

func (c *MainController) Say()  {
	s := c.Input().Get("hello")
	str_v := services.HandleHello(s)
	//c.Ctx.WriteString("user id by url :" + id + "\n\r 用户信息：" + content)
	c.Ctx.WriteString("rpcx 返回结果：" + str_v)
}
func (c *MainController) GetDealerInfo()  {
	id := c.Input().Get("id")
	content := models.Dealer{}
	id_v, _:= strconv.Atoi(id)
	content.Id = id_v
	content_v, _ := models.FindById(&content)

	//c.Ctx.WriteString("user id by url :" + id + "\n\r 用户信息：" + content)

	//JsonContent.Status = 1
	//JsonContent.Msg = "数据返回成功"
	//JsonContent.Data = content_v
	JsonV.EchoJson(1,"数据返回成功", content_v)
	c.Data["json"] = &JsonV
	c.ServeJSON()
}

type JsonContent struct {
	Status int `json:"status"`
	Msg string `json:"message"`
	Data interface{} `json:"data"`
}

func (j *JsonContent) EchoJson(status int, msg string, data interface{}){
	j.Status = status
	j.Msg = msg
	j.Data = data
}
