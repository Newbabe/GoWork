package action

import (
	"SongSyncProgram/service"
)

const Path = "/home/ec2-user/SongSyncProgram/"

func AgainDownloadSongFile() {
	//查询下载失败的MP3歌曲
	SourceIdList := service.QueryAllYoutubeMp3Log()
	mp3Dir := Path + "youtube_tv_mp3/"

	service.AgainDownLoadMp3(SourceIdList, mp3Dir)
	//查询下载失败的歌起文件
	LrcLogList := service.QueryAllYoutubeLrcLog()
	lrcDir := Path + "youtube_tv_lrc/"

	service.AgainDownLoadLyrics(LrcLogList, lrcDir)
}
