package service

import (
	"SongSyncProgram/dao"
	"SongSyncProgram/model"
	"SongSyncProgram/util"
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

var beiYuanSongSizeMap map[int]int

var beiYuanLrc1SizeMap map[int]int

var beiYuanLrc2SizeMap map[int]int

var beiYuanUpdateTimeMap map[int]string

const Path = "/home/ec2-user/SongSyncProgram/"

func UpdateDoremiUpdateTimeAndFile(doremiNewUpdateTimeMap map[int]model.DoremiSong) {
	//初始化map
	beiYuanSongSizeMap = make(map[int]int)
	beiYuanLrc1SizeMap = make(map[int]int)
	beiYuanLrc2SizeMap = make(map[int]int)
	beiYuanUpdateTimeMap = make(map[int]string)

	mp3Dir := Path + "youtube_tv_mp3/"
	lrcDir := Path + "youtube_tv_lrc/"
	//判断目录是否存在不存在创建目录
	mp3Direxists, _ := PathExists(mp3Dir)
	if !mp3Direxists {
		err := os.Mkdir(mp3Dir, 0766)
		if err != nil {
			fmt.Println("创建mp3DirFile路径异常")
			return
		}
	}
	//判断目录是否存在不存在创建目录
	lrcDirexist, _ := PathExists(lrcDir)
	if !lrcDirexist {
		err := os.Mkdir(lrcDir, 0766)
		if err != nil {
			fmt.Println("创建mp3DirFile路径异常")
			return
		}
	}
	//遍历map
	var num int
	for sourceId, doremiSong := range doremiNewUpdateTimeMap {
		var lrcSize1 int
		var lrcSize2 int
		num++
		fmt.Println("开始第", num, "首：", sourceId)
		doremiLogin := dao.GetDoremiLogin2()
		if doremiSong.Duet {
			LrcFileName := lrcDir + strconv.Itoa(sourceId) + ".lrc"
			LrcFileName2 := lrcDir + strconv.Itoa(sourceId) + "_2.lrc"
			/*
				220829添加歌曲歌词文件是否存在如果存在就跳过
			*/
			//下载
			fileExist := util.ExistS3SongFile("youtube_tv_lrc/" + strconv.Itoa(sourceId) + ".lrc")
			fileExist2 := util.ExistS3SongFile("youtube_tv_lrc/" + strconv.Itoa(sourceId) + "_2.lrc")
			//fileExist := util.ExistS3SongFile("youtube_tv_lrc_test/" + strconv.Itoa(sourceId) + ".lrc")
			//fileExist2 := util.ExistS3SongFile("youtube_tv_lrc_test/" + strconv.Itoa(sourceId) + "_2.lrc")
			if !fileExist {
				lrc1 := dao.DownloadLrc(doremiLogin, sourceId, 1)
				if lrc1.Bytes != nil && len(lrc1.Bytes) > 0 {
					//lrcSize1 = len(lrc1.Bytes)
					//先将文件写入服务器（创建lrc临时文件）
					lrcSize1 = UploadS3(lrc1, LrcFileName)
					util.RecordLogUtil([]byte(strconv.Itoa(sourceId) + "-----lrcSize1:" + strconv.Itoa(lrcSize1) + "---上传成功"))
					beiYuanLrc1SizeMap[sourceId] = lrcSize1
				} else { //歌词文件1下载失败
					lrcLog := QueryLrcLogBySourceId(sourceId)
					if lrcLog.Id == 0 {
						SaveLrcLog(sourceId, 1, 0, 2)
					}
				}
			} /*else {
				lrc1 := dao.DownloadLrc(doremiLogin, sourceId, 1)
				lrcSize1 = len(lrc1.Bytes)
				beiYuanLrc1SizeMap[sourceId] = lrcSize1
			}*/
			if !fileExist2 {
				lrc2 := dao.DownloadLrc(doremiLogin, sourceId, 2)
				if lrc2.Bytes != nil && len(lrc2.Bytes) > 0 {
					//lrcSize2 = len(lrc2.Bytes)
					lrcSize2 = UploadS3(lrc2, LrcFileName2)
					util.RecordLogUtil([]byte(strconv.Itoa(sourceId) + "-----lrcSize1:" + strconv.Itoa(lrcSize2) + "---上传成功"))
					beiYuanLrc2SizeMap[sourceId] = lrcSize2
				} else { //歌词文件2下载失败
					//判断改歌曲记录表中是否存在
					lrcLog := QueryLrcLogBySourceId(sourceId)
					if lrcLog.Id == 0 { //不存在说明lrc1文件下载成功
						SaveLrcLog(sourceId, 2, 1, 2)
					} else { //存在说明lrc1文件也下载失败了
						UpdateLrc2Log(sourceId, 1)
					}
				}
			} /*else {
				lrc2 := dao.DownloadLrc(doremiLogin, sourceId, 2)
				lrcSize2 = len(lrc2.Bytes)
				beiYuanLrc2SizeMap[sourceId] = lrcSize2
			}*/

		} else {
			fileExist := util.ExistS3SongFile("youtube_tv_lrc/" + strconv.Itoa(sourceId) + ".lrc")
			//fileExist := util.ExistS3SongFile("youtube_tv_lrc_test/" + strconv.Itoa(sourceId) + ".lrc")
			if !fileExist {
				lrc := dao.DownloadLrc(doremiLogin, sourceId, 0)
				if lrc.Bytes != nil && len(lrc.Bytes) > 0 {
					//lrcSize1 = len(lrc.Bytes)
					FileName := lrcDir + strconv.Itoa(sourceId) + ".lrc"
					lrcSize1 = UploadS3(lrc, FileName)
					util.RecordLogUtil([]byte(strconv.Itoa(sourceId) + "-----lrcSize1:" + strconv.Itoa(lrcSize1) + "---上传成功"))
					beiYuanLrc1SizeMap[sourceId] = lrcSize1

				} else { //如果下载失败就存入记录表里面[因为要重复运行所以要先判断表中是否存改歌曲数据]
					lrcLog := QueryLrcLogBySourceId(sourceId)
					if lrcLog.Id == 0 { //[单重奏歌曲  lrc2的歌词文件状态为0// ]
						SaveLrcLog(sourceId, 1, 0, 1)
					}
				}
			} /*else {
				lrc1 := dao.DownloadLrc(doremiLogin, sourceId, 1)
				lrcSize1 = len(lrc1.Bytes)
				beiYuanLrc1SizeMap[sourceId] = lrcSize1
			}*/
		}
		//下载mp3
		//mp3Exist := util.ExistS3SongFile("youtube_tv_mp3/" + strconv.Itoa(sourceId) + ".mp3")
		//mp3Exist := util.ExistS3SongFile("youtube_tv_mp3_test/" + strconv.Itoa(sourceId) + ".mp3")

		result := dao.DownloadMp3(doremiLogin, sourceId)
		if result.Bytes == nil || result.Code == 404 {
			Update(sourceId, lrcSize1, lrcSize2, 0, 1)
			fmt.Println("mp3 == null || mp3.length == 0：跳过-", sourceId)
			//如果歌曲下载失败就存入统计表中【先判断表中是否有改歌曲的数据】
			mp3Log := GetMp3LogBySourceId(sourceId)
			if mp3Log.Id == 0 { // 表中没有该歌曲记录时添加
				SaveMp3Log(sourceId)
			}
			continue
		}
		if result.Code == 403 {
			doremiLogin = dao.GetDoremiLogin()
			result = dao.DownloadMp3(doremiLogin, sourceId)
		}
		//fmt.Println("result.Code", result.Code)
		mp3 := result.Bytes
		if mp3 == nil || len(mp3) == 0 {
			Update(sourceId, lrcSize1, lrcSize2, 0, 1)
			fmt.Println("mp3 == null || mp3.length == 0：跳过-", sourceId)
			mp3Log := GetMp3LogBySourceId(sourceId)
			if mp3Log.Id == 0 { // 表中没有该歌曲记录时添加
				SaveMp3Log(sourceId)
			}
			continue
		}
		mp3Bytes := util.DownLoadFromUrl(string(mp3))
		mp3fileUrl := mp3Dir + strconv.Itoa(sourceId) + ".mp3"
		UploadS3(mp3Bytes, mp3fileUrl)
		util.RecordLogUtil([]byte(strconv.Itoa(sourceId) + "-----MP3---上传成功--songSize" + strconv.Itoa(len(mp3Bytes.Bytes))))
		beiYuanSongSizeMap[sourceId] = len(mp3Bytes.Bytes)
		beiYuanUpdateTimeMap[sourceId] = doremiSong.UpdateTime.String()
		//记录备源歌曲更新mp3文件  该文件的伴奏类型,目前只要记录source=3, source=4的暂时不用----2021-09-15 葛嘉豪[翻译cxp]
		//原因:手机版备源mp3的音档类型和song_info_merge表里lrc_channel字段记录不一致,导致安卓和ios播放同一首歌听到的伴奏不一样
		//解决办法:ios修改播放模式改为根据传的类型播放youTube歌曲,目前song_info_merge表里的lrc_channel
		//只要是备源歌曲这个字段代表的是影片的类型 真正备源音档的歌曲类型在youtube_backup_source_type表中
		SongMap := GetSongIdSngVersionLrcVersionBySourceId(sourceId)
		fmt.Println("SongMap", SongMap)
		if SongMap["id"] != "" {
			songIdStr := SongMap["id"]
			songVersionStr := SongMap["songVersion"]
			lrcVersionStr := SongMap["lrcVersion"]
			songId, _ := strconv.Atoi(songIdStr)
			songVersion, _ := strconv.Atoi(songVersionStr)
			lrcVersion, _ := strconv.Atoi(lrcVersionStr)
			dao.Save(songId, sourceId, doremiSong.YoutubeSongType, songVersion, lrcVersion, doremiSong.LrcChannel)
		} else {
			fmt.Println("songId等于null:", sourceId)
		}

	}

	fmt.Println("结束备援歌曲===需要更新的数量:", len(beiYuanSongSizeMap))
	fmt.Println("删除MP3文件", os.RemoveAll(mp3Dir))
	fmt.Println("删除LRC文件", os.RemoveAll(lrcDir))
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil

	}
	return false, err

}

func UploadS3(obj model.HttpResponseResult, fileUrl string) int {
	lrcSize1 := len(obj.Bytes)
	//先将文件写入服务器（创建lrc临时文件）
	lrc1file, _ := os.Create(fileUrl)
	defer lrc1file.Close()
	writer := bufio.NewWriter(lrc1file)
	writer.Write(obj.Bytes)
	writer.Flush()
	//读取服务器的文件写入亚马逊数据库
	open, _ := os.Open(fileUrl)
	defer open.Close()
	fileName := path.Base(open.Name())
	var uploadPath string
	if strings.Contains(fileName, ".mp3") {
		uploadPath = "youtube_tv_mp3/" + fileName
	} else {
		uploadPath = "youtube_tv_lrc/" + fileName
	}
	//测试目录
	//fmt.Println("uploadPath", uploadPath)
	//上传部分先注掉
	err := util.UploadS3(uploadPath, open)
	if err != nil {
		fmt.Println("上传文件异常:", err)
		return 0
	}
	//删除服务器文件
	os.Remove(fileUrl)
	return lrcSize1
}
func S3YoutubeStructLrcFile(obj model.HttpResponseResult, fileUrl string) {
	//先将文件写入服务器（创建lrc临时文件）
	lrc1file, err3 := os.Create(fileUrl)
	if err3 != nil {
		fmt.Println("打开文件异常", err3)
	}
	//fmt.Println("fileUrl", fileUrl)
	defer lrc1file.Close()
	writer := bufio.NewWriter(lrc1file)
	_, err2 := writer.Write(obj.Bytes)
	if err2 != nil {
		fmt.Println("写入文件异常", err2)
		return
	}
	writer.Flush()
	//读取服务器的文件写入亚马逊数据库
	open, err := os.Open(fileUrl)
	if err != nil {
		fmt.Println("打开文件异常", err)
	}
	defer open.Close()
	fileName := path.Base(open.Name())

	var uploadPath string
	//测试目录
	//uploadPath = "YTsong_test/" + fileName
	uploadPath = "youtube_struct_lyrics/" + fileName
	//fmt.Println("uploadPath", uploadPath)
	//上传部分先注掉
	err = util.UploadS3(uploadPath, open)
	if err != nil {
		fmt.Println("上传文件异常:", err)
	}
	//删除服务器文件
	os.Remove(fileUrl)
}
func updateDoreminNewLrcAndFile(doremiNewLrcMap map[int]int) {
	lrcDir := Path + "youtube_struct_lyrics/"
	//判断路径是否存在
	pathExists, _ := PathExists(lrcDir)
	if !pathExists { //如果不存在就创建目录
		err := os.Mkdir(lrcDir, 0766)
		if err != nil {
			fmt.Println("创建mp3DirFile路径异常")
			return
		}
	}
	//遍历map
	var count int
	for sourceId, flag := range doremiNewLrcMap {
		doremiLogin := dao.GetDoremiLogin2()
		lrcCount := flag
		hasStruct := false
		var clip1 string
		var clip2 string
		if lrcCount == 1 {
			lrc1Obj := dao.DownloadStructLrc(doremiLogin, sourceId, 0)
			if string(lrc1Obj.Bytes) != "" && len(lrc1Obj.Bytes) > 0 {
				hasStruct = true
				fileName := lrcDir + strconv.Itoa(sourceId) + ".lrc"
				//判断lrc文件是否存在如果存在就跳过[此路径椒亚马逊服务器文件路径]
				fileExit := util.ExistS3SongFile("youtube_struct_lyrics/" + strconv.Itoa(sourceId) + ".lrc")
				if fileExit {
					continue
				}
				clip1 = util.ReadSongStructure(lrc1Obj.Bytes)
				S3YoutubeStructLrcFile(lrc1Obj, fileName)
			} else {
				fmt.Println("结构歌词下载失败:", sourceId)
				dao.SaveUploadLrcFailRecord(sourceId)
			}
		} else if lrcCount == 2 {
			hasStruct = true
			lrc1 := dao.DownloadStructLrc(doremiLogin, sourceId, 1)
			lrc2 := dao.DownloadStructLrc(doremiLogin, sourceId, 2)
			if string(lrc1.Bytes) != "" && len(string(lrc1.Bytes)) > 0 {
				fileName := lrcDir + strconv.Itoa(sourceId) + ".lrc"
				fileExit := util.ExistS3SongFile("youtube_struct_lyrics/" + strconv.Itoa(sourceId) + ".lrc")
				if !fileExit {
					S3YoutubeStructLrcFile(lrc1, fileName)
					clip1 = util.ReadSongStructure(lrc1.Bytes)
				}

			} else {
				fmt.Println("结构歌词下载失败:", sourceId)
				dao.SaveUploadLrcFailRecord(sourceId)
			}

			if string(lrc2.Bytes) != "" && len(string(lrc2.Bytes)) > 0 {
				fileName := lrcDir + strconv.Itoa(sourceId) + "_2.lrc"
				fileExit := util.ExistS3SongFile("youtube_struct_lyrics/" + strconv.Itoa(sourceId) + "_2.lrc")
				if !fileExit {
					S3YoutubeStructLrcFile(lrc2, fileName)
					clip2 = util.ReadSongStructure(lrc2.Bytes)
				}
			} else {
				fmt.Println("结构歌词下载失败")
				dao.SaveUploadLrcFailRecord(sourceId)
			}
		}
		if hasStruct {
			curSongInfo := GetSongNameAndSingerBySongInfoMergeSourceId(sourceId)
			songClipId := GetSongClipInfoDaoBySongId(curSongInfo.Id)
			if songClipId == 0 {
				if clip1 != "" || clip2 != "" {
					fmt.Println("保存结构歌词：", curSongInfo.Id, curSongInfo.SongName)
					SaveSongClipInfo(curSongInfo.Id, curSongInfo.SongName, clip1, clip2)
					count++
				}
			} else {
				UpdateSongClipInfo(curSongInfo.Id, curSongInfo.SongName, clip1, clip2)
				count++
			}
		}

	}
	fmt.Println("结束开始下载结构歌词===需要更新的数量:", count)
	fmt.Println("删除LRC文件=", os.RemoveAll(lrcDir))
}

func UpdateBeiYuanFileSize() {
	fmt.Println("beiYuanSongSizeMap_长度", len(beiYuanSongSizeMap))
	fmt.Println("beiYuanUpdateTimeMap_长度", len(beiYuanUpdateTimeMap))

	for sourceId, songSize := range beiYuanSongSizeMap {
		var lrcSize1 int
		var lrcSize2 int
		result1 := containsKey(beiYuanLrc1SizeMap, sourceId)
		if result1 {
			lrcSize1 = beiYuanLrc1SizeMap[sourceId]
		}
		result2 := containsKey(beiYuanLrc2SizeMap, sourceId)
		if result2 {
			lrcSize2 = beiYuanLrc2SizeMap[sourceId]
		}
		//更新song_info_merge_tv
		fmt.Println("sourceId:", sourceId, ",songSize:", songSize, ",lrcSize1:", lrcSize1, ",lrcSize2:", lrcSize2)
		Update(sourceId, lrcSize1, lrcSize2, songSize, 2)
		UpdatePhone(sourceId, lrcSize1, lrcSize2, songSize, 2)
		updateTime := beiYuanUpdateTimeMap[sourceId]
		dao.UpdatelastUpdate(sourceId, updateTime)
		util.RecordLogUtil([]byte("更新上传成功歌曲Doremi时间" + strconv.Itoa(sourceId)))
		fmt.Println("==========新增加同步代码=============")
		curSongInfo := GetSongNameAndSingerBySongInfoMergeSourceId(sourceId)
		//fmt.Println("curSongInfo", curSongInfo)
		if curSongInfo.Id != 0 {
			needFilterSongList := GetSongInfoMergeBySongNameAndSinger(curSongInfo.SongName, curSongInfo.Singer)
			//	fmt.Println("needFilterSongList", needFilterSongList)
			if needFilterSongList != nil && len(needFilterSongList) > 0 {
				for i := 0; i < len(needFilterSongList); i++ {
					id := GetYoutubeSongFilterBySongId(needFilterSongList[i])
					fmt.Println("id", id)
					if id == 0 {
						fmt.Println("YoutubeSongFilterId", id)
						SaveYoutubeSongFilter(needFilterSongList[i])
					}
				}
			}
		}
	}
}

//判断map中是否存在该键
func containsKey(Map map[int]int, key int) bool {
	for k, _ := range Map {
		if key == k {
			return true
		}
	}
	return false
}
