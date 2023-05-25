package main

import (
	"SongSyncProgram/action"
	"fmt"
	"net/http"
	"time"
)

func main() {

	//歌曲同步程序
	action.Crontab()
	server := http.Server{
		Addr:         ":9034", //测试端口号
		ReadTimeout:  500 * time.Second,
		WriteTimeout: 500 * time.Second,
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("ListenAndServe err", err)
	}
}

/*func main() {
	action.StartYoutubePhone()
	action.StartYoutubeTV()

}*/

/*func main() {
	//todo  记得更换 文件服务器地址 /home/ec2-user/GOTest/MalaysiaTEST
	action.StartYoutubeMalysia()
}*/
