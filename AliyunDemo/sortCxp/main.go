package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	conn, err := rpc.DialHTTPPath("tcp", "47.74.221.96:1099", "/OkeLiveRMI")
	if err != nil {
		fmt.Println(err)
		return
	}
	var str string
	err = conn.Call("OkeLiveRMIInterface.getChatRoom", false, &str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("str", str)
}
