package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type Dealer struct {
	Id int `pk`
	Dealer_name string
	Dealer_email string
	Dealer_username string
	Dealer_pwd string
	Create_time string
}
var DealerM  orm.Ormer
func init()  {
	//orm.RegisterModel(new(Dealer))
	orm.RegisterModelWithPrefix(beego.AppConfig.String("mysqldbprefix"), new(Dealer))
}
func TableName() string {

	return "oil_dealer"
}

//func AddUser(u *User) error{
//
//	o := orm.NewOrm()
//	id, err := o.Insert(u)
//
//	if err != nil {
//		logs.Error("生成用户失败", err)
//	}else{
//		logs.Info("生成用户id:", id)
//	}
//	return nil
//}
//
func FindById(u *Dealer) (*Dealer,error) {
	o := orm.NewOrm()
	err := o.Read(u)
	//err := DealerM.Read(u)

	if err != nil {
		logs.Error("查询用户失败", err)
	}
	return u, nil
}
func (d *Dealer) FindByUsernameAndPwd(username string, pwd string) bool {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("id", "dealer_name", "dealer_pwd", "dealer_username").
		From(TableName()).
		Where("dealer_username = ? and dealer_pwd = ?").
		Limit(1)
	sql := qb.String()
	dealer_u := Dealer{}
	o.Raw(sql, username,pwd).QueryRow(&dealer_u)
	if dealer_u.Id > 0 {
		fmt.Println("查询到用户 id=", dealer_u.Id)
		return true
	}else{
		fmt.Println("未查询到用户")
		return false
	}
	//c,err := o.QueryTable(new(Dealer)).Filter("dealer_username",username).Filter("dealer_pwd", pwd).Count()
	//if err != nil {
	//	logs.Error("登陆用记名或密码不正确")
	//	return false
	//}
	//if c > 0 {
	//	return true
	//}else{
	//	return false
	//}

}
//
//func FindOneUser(id int) (User, error) {
//	var u_data User
//	o := orm.NewOrm()
//	err := o.QueryTable(new(User)).Filter("id", id).One(&u_data)
//
//	if err != nil {
//		logs.Error("查询用户失败", err)
//	}
//	return u_data, nil
//}