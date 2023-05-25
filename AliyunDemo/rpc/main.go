package main

import (
	"HeroOkWebGo/src/model"
	"HeroOkWebGo/src/rpc"
	"HeroOkWebGo/src/service"
	"context"
	"encoding/json"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"net"
	"os"
	"runtime/debug"
	"strconv"
	"time"
)

func main() {
	//list := getLastRoomIdList()
	// 当前所有房间和对应活跃值
	/*66666
	2432
	2431
	1888
	*/
	// 1971 1025 1028
	service.IsSpecialYoutubeRoom(1888)

}

func getLastRoomIdList() []int {
	var list = make([]int, 0, 5)
	list = append(list, 69846)
	list = append(list, 69921)
	list = append(list, 69988)
	list = append(list, 65450)
	list = append(list, 66957)
	return list
}
func GetInactiveRoomList() []int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()

		}
	}()
	var list = make([]int, 0)
	data := getRpc2RoomServerData("getChatRoom", map[string]string{"clear": "false"})
	if data[0] == "0" && data[1] == "成功" {
		expireTime := 120000 //2分钟内没说话
		var roomMap = make(map[string]int64)
		json.Unmarshal([]byte(data[2]), &roomMap)
		now := time.Now().UnixNano() / 1e6
		for roomId, t := range roomMap {
			if now-t > int64(expireTime) {
				atoi, _ := strconv.Atoi(roomId)
				list = append(list, atoi)
			}
		}
	}
	return list
}

func getRpc2RoomServerData(serviceMethod string, sendContent map[string]string) []string {
	client, ctx, socket := getConnection()
	defer socket.Close()
	str, err := client.FunCall(ctx, 0, serviceMethod, sendContent)
	if err != nil {
		fmt.Println("getRpc2RoomServerData", err)
		debug.PrintStack()
	}
	return str
}

func getConnection() (*rpc.RpcServiceClient, context.Context, *thrift.TSocket) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()
		}
	}()
	port := "19094"
	//if testRoom {
	//	port = "19090"
	//}
	host := "172.21.79.6"
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(host, port))

	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
	}

	useTransport, err := transportFactory.GetTransport(transport)
	client := rpc.NewRpcServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 47.74.229.230:"+port, " ", err)
	}

	ctx, _ := context.WithCancel(context.Background())

	return client, ctx, transport
}

func GetRoomSongs(roomId int) []model.OkeLiveSongInfo {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()

		}
	}()
	var songList = make([]model.OkeLiveSongInfo, 0)
	data := getRpc2RoomServerData("getRoomSongs", map[string]string{"roomId": strconv.Itoa(roomId)})
	if data[0] == "0" && data[1] == "成功" {
		json.Unmarshal([]byte(data[2]), &songList)
		return songList
	}
	return songList
}

type Data struct {
	Iid          string `json:"iid"`
	UserID       string `json:"userId"`
	SongInfoData string `json:"songInfoData"`
}
type OkeLiveSongInfo struct {
	Iid      string `json:"iid"`
	UserId   string `json:"userId"`
	SongData string `json:"songInfoData"`
	// 尊享用户
	VipUsers string `json:"vipUsers,omitempty"`
	// 请求合唱用户
	ReqUsers string `json:"reqUsers,omitempty"`
	// 合唱用户
	ChorusUsers string `json:"chorusUsers,omitempty"`
	ZxLen       int    `json:"zxLen,omitempty"`
}

func GetRoomUserList(roomId int) []model.OkeLiveUserInfo {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			debug.PrintStack()

		}
	}()
	var userList = make([]model.OkeLiveUserInfo, 0)
	data := getRpc2RoomServerData("getRoomUsers", map[string]string{"roomId": strconv.Itoa(roomId)})
	if len(data) < 3 {
		return userList
	}
	if data[0] == "0" && data[1] == "成功" {
		fmt.Println("data[2]", data[2])
		json.Unmarshal([]byte(data[2]), &userList)
		return userList
	}
	return userList
}
