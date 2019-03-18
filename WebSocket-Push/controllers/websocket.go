package controllers

import (
	"github.com/astaxie/beego"
	"golang.org/x/net/websocket"
	"github.com/astaxie/beego/cache"
	m "git.culiu.org/pintuan-go/models"
	"time"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"sync"
)

// WebSocketController handles WebSocket requests.
type WebSocketController struct {
	beego.Controller
}

type OnlineUser struct {
	connection *websocket.Conn
	clientIP   string
	time       string
	GoodsId    int
}
var  CN_SockServer = 0;
var  CN_SockServerDeteail = 0;
const TIME_FORMAT = "2006-01-02 03:04:05 PM"

var (
	cached cache.Cache
	ActiveUsers = make(map[string]*OnlineUser)
	lock sync.Mutex
)

func init() {
	cached, _ = cache.NewCache("memory", `{"interval":300}`)
}

//var TASK_CHANNEL = make(chan *OnlineUser)
func SockServer(ws *websocket.Conn) {
	tm := time.Unix(time.Now().Unix(), 0)
	onlineUser := &OnlineUser{
		connection:ws,
		clientIP:ws.Request().RemoteAddr,
		time:tm.Format(TIME_FORMAT),
	}
	lock.Lock()
	hash := md5.New()
	hash.Write([]byte(onlineUser.clientIP))
	hashByte := hash.Sum(nil)
	hashStr := hex.EncodeToString(hashByte)
	ActiveUsers[hashStr] = onlineUser
	lock.Unlock()
	var orders []m.Order
	if cached.IsExist("last_order_limit") {
		orders = cached.Get("last_order_limit").([]m.Order)
	} else {
		orders = m.GetOrderLimit(1)
		cached.Put("last_order_limit", orders, 60 * time.Second)
	}
	if len(orders) <=  0{
		orders = m.GetOrderLimit(1)
	}
	onlineUser.pushToClient(orders)
	onlineUser.closeUser()
}

func SockServerDeteail(ws *websocket.Conn) {
	tm := time.Unix(time.Now().Unix(), 0)
	onlineUser := &OnlineUser{
		connection:ws,
		clientIP:ws.Request().RemoteAddr,
		time:tm.Format(TIME_FORMAT),
	}
	lock.Lock()
	hash := md5.New()
	hash.Write([]byte(onlineUser.clientIP))
	hashByte := hash.Sum(nil)
	hashStr := hex.EncodeToString(hashByte)
	ActiveUsers[hashStr] = onlineUser
	lock.Unlock()

	ws.Request().ParseForm()
	Goods :=ws.Request().FormValue("goodsid")
	cacheName := "last_order_limit_detail_"+Goods;
	GoodsId,error:= strconv.Atoi(Goods)
	if error != nil{
		fmt.Println("goodsid atoi false")
	}
	var orders []m.Order
	if cached.IsExist(cacheName) {
		orders = cached.Get(cacheName).([]m.Order)
		fmt.Println("text")
	} else {
		orders = m.GetOrderLimit(GoodsId)
		cached.Put(cacheName, orders, 60 * time.Second)
	}
	if len(orders) <=  0{
		orders = m.GetOrderLimit(GoodsId)
	}
	onlineUser.pushToClient(orders)
	onlineUser.closeUser()
}
func (user *OnlineUser) pushToClient(rows []m.Order) {
	/*var list = make(map[int]int)
	var list_id []int*/
	type UserOrderInfo struct {
		User_id    int64  `json:"user_id"`
		Headimgurl string `json:"headimgurl"`
		Uname      string `json:"uname"`
		Add_time   int    `json:"add_time"`
		Goods_id   int    `json:"goods_id"`
		Team_sign  string `json:"team_sign"`
		Order_type  int   `json:"order_type"`
	}

	userOrderInfo := []UserOrderInfo{}

	for _, order := range rows {
		user_id := int64(order.User_id)
		//mysql查询出来的user_id不会重复，感觉这一部不用写   如果非得需要可以直接mysql去重查询
		/*if list[user_id] != 1 {
			list[user_id] = 1
			list_id = append(list_id, user_id)
		}*/
		userOne := m.GetUserInfoOne(user_id)
		userInfo := UserOrderInfo{
			User_id:    user_id,
			Headimgurl: userOne.Headimgurl,
			Uname:      userOne.Uname,
			Add_time:   order.Add_time,
			Goods_id:   order.Goods_id,
			Team_sign:   order.Team_sign,
			Order_type:   order.Order_type,
		}
		userOrderInfo = append(userOrderInfo, userInfo)
	}
	err := websocket.JSON.Send(user.connection, userOrderInfo)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func (user *OnlineUser) closeUser() {
	lock.Lock()
	user.connection.Close()
	delete(ActiveUsers, user.clientIP)
	lock.Unlock()
}
