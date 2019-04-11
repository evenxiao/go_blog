package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int `pk`
	User_phone string
	User_name string
	Pwd string
	Reg_time string
}

func init()  {
	//orm.RegisterModel(new(User))
	orm.RegisterModelWithPrefix(beego.AppConfig.String("mysqldbprefix"), new(User))
	//orm.RegisterModelWithPrefix("oil_", new(User))
}
func (u *User) TableName() string {

	return "user"
}

func AddUser(u *User) error{

	o := orm.NewOrm()
	id, err := o.Insert(u)

	if err != nil {
		logs.Error("生成用户失败", err)
	}else{
		logs.Info("生成用户id:", id)
	}
	return nil
}

func FindById(u *User)  {
	o := orm.NewOrm()
	err := o.Read(u)

	if err != nil {
		logs.Error("查询用户失败", err)
	}
}

func FindOneUser(id int) (User, error) {
	var u_data User
	o := orm.NewOrm()
	err := o.QueryTable(new(User)).Filter("id", id).One(&u_data)

	if err != nil {
		logs.Error("查询用户失败", err)
	}
	return u_data, nil
}