package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	User_id    int64  `orm:"pk;"`
	Headimgurl string `orm:"size(255);"`
	Uname      string `orm:"size(255);"`
}

func init() {
	orm.RegisterModel(new(Users))
}

func (u *Users) TableName() string {
	return "hhs_users"
}
/**
按照where里面的in条件查询多条数据
 */
func GetUserInfo(user_id []int64) (users []Users) {
	user := new(Users)
	o := orm.NewOrm()
	qs := o.QueryTable(user)

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.From("hhs_users").In()
	var userList []Users
	//qs.Filter("hhs_users__user_id__in", user_id).All(&userList)
	qs.Filter("user_id__in", user_id).All(&userList)
	return userList
}
/**
按照where条件查询一条数据
 */
func GetUserInfoOne(user_id int64) (users Users) {
	o := orm.NewOrm()
	qs := o.QueryTable("hhs_users")
	qs.Filter("user_id", user_id).One(&users)
	return users
}