package services

import (
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/client"
	"golang.org/x/net/context"
	"log"
	"study/day1/rpcx_service"
	"time"
)
var (
	addr = flag.String("addr", "localhost:8972", "server address")
)
func HandleUser(id int) rpcx_service.UserRes{
	start_time := time.Now()
	data := rpcx_service.UserRes{}

	//id := r.FormValue("id")
	////id := ps.ByName("id")
	//id_v, _ := strconv.Atoi(id)
	id_v := int(id)
	d := client.NewPeer2PeerDiscovery("tcp@" +*addr, "")
	xclient := client.NewXClient("UserRes", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	err := xclient.Call(context.Background(), "GetById",id_v, &data)
	fmt.Println("rpc调用完后时间：", time.Since(start_time))
	if err != nil {
		log.Fatalf("failed to call :%v", err)
	}

	log.Printf("查询用户id的数据========== 》》》 用户名: %s, 手机号：%s", data.User_name, data.User_phone)

	contents := " 你好用户 " + data.User_name + ", 手机：" + data.User_phone

	contents += "\n\r ^___^"
	return data
}

func HandleHello(s string) string{
	start_time := time.Now()
	//data := rpcx_service.UserRes{}

	//id := r.FormValue("id")
	////id := ps.ByName("id")
	//id_v, _ := strconv.Atoi(id)
	reply_s := ""
	d := client.NewPeer2PeerDiscovery("tcp@" +*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	if s == "" {
		s = "阿文 你好，RPCX"
	}
	err := xclient.Call(context.Background(), "Say", s, &reply_s)
	fmt.Println("rpc调用完后时间：", time.Since(start_time))
	if err != nil {
		log.Fatalf("failed to call :%v", err)
	}

	return reply_s
}