package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {
	//c.Data["Website"] = "beego.me"
	myStruct := make(map[string]string)
	//myStruct :=

	myStruct["name"] = "awen"
	myStruct["age"] = "27"
	myStruct["sex"] = "man"
	c.Data["json"] = &myStruct
	c.ServeJSON()
}

type UserData struct {
	Username string
	Age string
	Sex string
	// string
}
