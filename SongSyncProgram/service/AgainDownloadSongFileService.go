package service

import (
	"SongSyncProgram/dao"
	"SongSyncProgram/model"
	"SongSyncProgram/util"
	"fmt"
	"os"
	"strconv"
)

func AgainDownLoadMp3(sourceIdList []int, mp3Dir string) {
	mp3DirExists, _ := PathExists(mp3Dir)
	if !mp3DirExists {
		err := os.Mkdir(mp3Dir, 0766)
		if err != nil {
			fmt.Println("创建mp3DirFile路径异常")
			return
		}
	}
	for _, sourceId := range sourceIdList {
		doremiLogin := dao.GetDoremiLogin2()
		mp3Exist := util.ExistS3SongFile("youtube_tv_mp3/" + strconv.Itoa(sourceId) + ".mp3")

		if !mp3Exist { //如果歌曲不存在
			result := dao.DownloadMp3(doremiLogin, sourceId)
			if result.Bytes == nil || result.Code == 404 {
				fmt.Println("mp3 == null || mp3.length == 0：跳过-", sourceId)
				continue
			}

			if result.Code == 403 {
				doremiLogin = dao.GetDoremiLogin()
				result = dao.DownloadMp3(doremiLogin, sourceId)
			}
			mp3 := result.Bytes
			if mp3 == nil || len(mp3) == 0 {
				//		Update(sourceId, lrcSize1, lrcSize2, 0, 1)
				fmt.Println("mp3 == null || mp3.length == 0：跳过-", sourceId)
				mp3Log := GetMp3LogBySourceId(sourceId)
				fmt.Println("mp3Log1", mp3Log)
				if mp3Log.Id == 0 { // 表中没有该歌曲记录时添加
					SaveMp3Log(sourceId)
				}
				continue
			}
			mp3Bytes := util.DownLoadFromUrl(string(mp3))
			mp3fileUrl := mp3Dir + strconv.Itoa(sourceId) + ".mp3"
			UploadS3(mp3Bytes, mp3fileUrl)
			//上传成功更新
			UpdateMp3Log(sourceId)
			beiYuanSongSizeMap[sourceId] = len(mp3Bytes.Bytes)
			//从最后时间表里面查询出updateTime
			beiYuanUpdateTimeMap[sourceId] = GetLasUpdateTimeBySourceId(sourceId).String()
		}
	}
	fmt.Println("删除MP3文件", os.RemoveAll(mp3Dir))

}

func AgainDownLoadLyrics(LrcLogList []model.LrcLog, lrcDir string) {

	//判断目录是否存在不存在创建目录
	lrcDirexist, _ := PathExists(lrcDir)
	if !lrcDirexist {
		err := os.Mkdir(lrcDir, 0766)
		if err != nil {
			fmt.Println("创建mp3DirFile路径异常")
			return
		}
	}
	var num int
	for _, LrcLog := range LrcLogList {
		var lrcSize1 int
		var lrcSize2 int
		num++
		fmt.Println("开始第", num, "首：", LrcLog.SourceId)
		doremiLogin := dao.GetDoremiLogin2()
		if LrcLog.Duet == 2 {
			LrcFileName := lrcDir + strconv.Itoa(LrcLog.SourceId) + ".lrc"
			LrcFileName2 := lrcDir + strconv.Itoa(LrcLog.SourceId) + "_2.lrc"
			/*
				220829添加歌曲歌词文件是否存在如果存在就跳过
			*/
			//下载
			fileExist := util.ExistS3SongFile("youtube_tv_lrc/" + strconv.Itoa(LrcLog.SourceId) + ".lrc")

			fileExist2 := util.ExistS3SongFile("youtube_tv_lrc/" + strconv.Itoa(LrcLog.SourceId) + "_2.lrc")

			if !fileExist {
				lrc1 := dao.DownloadLrc(doremiLogin, LrcLog.SourceId, 1)
				if lrc1.Bytes != nil && len(lrc1.Bytes) > 0 {
					//lrcSize1 = len(lrc1.Bytes)
					//先将文件写入服务器（创建lrc临时文件）
					lrcSize1 = UploadS3(lrc1, LrcFileName)
					beiYuanLrc1SizeMap[LrcLog.SourceId] = lrcSize1
					//如果下载成功就更改记录表状态
					UpdateLrc1Log(LrcLog.SourceId, 2)
				} else {
					fmt.Println("歌词文件下载异常", LrcLog.SourceId)
				}
			}
			if !fileExist2 {
				lrc2 := dao.DownloadLrc(doremiLogin, LrcLog.SourceId, 2)
				if lrc2.Bytes != nil && len(lrc2.Bytes) > 0 {
					//lrcSize2 = len(lrc2.Bytes)
					lrcSize2 = UploadS3(lrc2, LrcFileName2)
					beiYuanLrc2SizeMap[LrcLog.SourceId] = lrcSize2
					UpdateLrc2Log(LrcLog.SourceId, 2)
				} else {
					fmt.Println("歌词文件下载异常", LrcLog.SourceId)
				}
			}

		} else {
			//	fileExist := util.ExistS3SongFile("youtube_tv_lrc_test/" + strconv.Itoa(LrcLog.SourceId) + ".lrc")
			fileExist := util.ExistS3SongFile("youtube_tv_lrc/" + strconv.Itoa(LrcLog.SourceId) + ".lrc")
			if !fileExist {
				lrc := dao.DownloadLrc(doremiLogin, LrcLog.SourceId, 0)
				if lrc.Bytes != nil && len(lrc.Bytes) > 0 {
					FileName := lrcDir + strconv.Itoa(LrcLog.SourceId) + ".lrc"
					lrcSize1 = UploadS3(lrc, FileName)
					beiYuanLrc1SizeMap[LrcLog.SourceId] = lrcSize1
					UpdateLrc1Log(LrcLog.SourceId, 2)
				} else {
					fmt.Println("歌词文件下载异常", LrcLog.SourceId)
				}
			}
		}

	}
	//最后删除歌曲文件
	fmt.Println("删除LRC文件", os.RemoveAll(lrcDir))
}
