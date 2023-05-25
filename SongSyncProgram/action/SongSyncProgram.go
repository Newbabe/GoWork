package action

import (
	"fmt"
	"github.com/robfig/cron"
)

// 定时任务
func Crontab() {

	crontab := cron.New() //精确到秒
	str := "0 7 * * * "   //每日7时0分执行同步程序
	_, err := crontab.AddFunc(str, SongSyncProgram)
	if err != nil {
		fmt.Println("定时任务异常关闭定时任务:", err)
		crontab.Stop()
		return
	}
	crontab.Start()

}
func SongSyncProgram() {
	StartYoutubePhone()
	StartYoutubeTV()
	//更新新进歌曲到solr服务器
	UpdateNewSongSolr()
}
