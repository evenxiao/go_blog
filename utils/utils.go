package utils

import (
	"crypto/md5"
	"flag"
	"fmt"
)
var (
	addr = flag.String("addr", "localhost:8972", "server address")
)
/**
* md5加密
 */
func Md5Encode(s string) string{
	data := []byte(s)

	has := md5.Sum(data)

	return fmt.Sprintf("%x", has)
}

//func GetRpcUser(id int) string{
//	data := rpcx_service.UserRes{}
//
//	d := client.NewPeer2PeerDiscovery("tcp@" +*addr, "")
//	xclient := client.NewXClient("UserRes", client.Failtry, client.RandomSelect, d, client.DefaultOption)
//	err := xclient.Call(context.Background(), "GetById",id, &data)
//
//	if err != nil {
//		log.Fatalf("failed to call :%v", err)
//	}
//
//	log.Printf("查询用户id的数据========== 》》》 用户名: %s, 手机号：%s", data.User_name, data.User_phone)
//	contents := " 你好用户 " + data.User_name + ", 手机：" + data.User_phone
//
//	contents += "\n\r ^___^"
//	return contents
//}
