package action

import (
	"SongSyncProgram/malysiaService"
	"fmt"
	"time"
)

func StartYoutubeMalysia() {
	nowDateTime := time.Now().Format("2006-01-02 15:04:05")
	//
	TimeStamp1 := time.Now().Unix()
	//先统计平差值表的信息
	malysiaService.SaveAutomaticUpdateLog()
	fmt.Println("start:" + nowDateTime)
	malysiaService.CopyTable()
	fmt.Println("第一步完成copyTable")
	malysiaService.AddSongInfoMergeTmp()
	//把数据库内容改为-1AddSongInfoMergeTmp
	//malysiaService.UpdateMergeSongDb() //【将旧版本的歌曲状态改为下架】
	TimeStamp2 := time.Now().Unix()
	fmt.Println("第二步完成 updateMergeSongDb :", TimeStamp2-TimeStamp1)
	malysiaService.AddYoutubeSongDb()
	TimeStamp3 := time.Now().Unix()
	fmt.Println("第三步完成 updateMergeSongDb :", TimeStamp3-TimeStamp2)
	malysiaService.EndTable()
	TimeStamp4 := time.Now().Unix()
	fmt.Println("第四步完成 updateMergeSongDb :", TimeStamp4-TimeStamp3)
	fmt.Println("总用时", TimeStamp4-TimeStamp1)
}
