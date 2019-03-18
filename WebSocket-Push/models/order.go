package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//订单表
type Order struct {
	Order_id    int    `orm:"pk;"`
	/*Order_sn string `orm:"unique;size(30);"`*/
	User_id     int64    `orm:"int"`
	Goods_id    int    `orm:"int"`
	Add_time    int    `orm:"int"`
	Team_sign   string `orm:"unique;size(30);"`
	Team_status int `orm:"int"`
	Order_type  int `orm:"int"`;
	Extension_id    int    `orm:"int"`
	//	order_status    string `orm:"-" `
	//	shipping_status string `orm:"unique;size(32)"`
	//	pay_status      string `orm:"size(32)"`
	//	consignee       string `orm:"null;size(200)"`
	//	country         int    `orm:"default(2)"`
	//	city            string `orm:"null;type(datetime)"`
	//	district        string `orm:"type(datetime);" `
	//	address         string `orm:"rel(m2m)"`
}

func init() {
	orm.RegisterModel(new(Order))
}

func (u *Order) TableName() string {
	return "hhs_order_info"
}

/*func GetOrderLimit() (orders []orm.Params) {
	order := new(Order)
	o := orm.NewOrm()
	qs := o.QueryTable(order)
	//err := o.Read(&order, "order_sn")
	qs.Limit(5).OrderBy("order_id").Values(&orders)
	//if err == orm.ErrNoRows {
	//	fmt.Println("查询不到")
	//} else if err == orm.ErrMissPK {
	//	fmt.Println("找不到主键")
	//} else if err != nil {
	//	fmt.Println(err)
	//}
	return orders
}*/
//
func GetOrderLimit(team_strtus int) (orders []Order) {
	o := orm.NewOrm()
	qs := o.QueryTable("hhs_order_info")
	if team_strtus == 1 {
		qs.Filter("team_status", 1).Limit(5).OrderBy("-order_id").All(&orders)
	}else if team_strtus > 1{
		qs.Filter("extension_id",team_strtus ).Limit(5).OrderBy("-order_id").All(&orders)
	}else{
		qs.Limit(5).OrderBy("-order_id").All(&orders)
	}
	return orders
}
