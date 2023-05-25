package dao

import (
	"SongSyncProgram/model"
	"SongSyncProgram/util"
	"strconv"
)

func DownloadLrc(doremiLogin model.DoremiLogin, sourceId, duetMode int) model.HttpResponseResult {
	lrcUrl := "https://api.karadoremi.net/api/download/download_lrc.php"
	lrcParam := "device_id=999&platform=8&user_id=" + doremiLogin.UserId + "&session_id=" + doremiLogin.SessionId +
		"&timestamp=" + doremiLogin.Timestamp + "&song_id=" + strconv.Itoa(sourceId) + "&lrc_version=1.0&client_version=0.0.0" +
		"&service=3&duet_mode=" + strconv.Itoa(duetMode)
	return util.HttpsPost(lrcUrl, []byte(lrcParam))

}

func DownloadStructLrc(doremiLogin model.DoremiLogin, sourceId, duetMode int) model.HttpResponseResult {
	lrcUrl := "https://api.karadoremi.net/api/download/download_lrc.php"
	lrcParam := "device_id=999&platform=8&user_id=" + doremiLogin.UserId + "&session_id=" + doremiLogin.SessionId +
		"&timestamp=" + doremiLogin.Timestamp + "&song_id=" + strconv.Itoa(sourceId) + "&lrc_version=1.4&client_version=0.0.0" +
		"&service=3&duet_mode=" + strconv.Itoa(duetMode)
	resp := util.HttpsPost(lrcUrl, []byte(lrcParam))
	return resp
}
