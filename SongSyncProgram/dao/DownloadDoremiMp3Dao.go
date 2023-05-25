package dao

import (
	"SongSyncProgram/model"
	"SongSyncProgram/util"
	"strconv"
)

func DownloadMp3(doremiLogin model.DoremiLogin, sourceId int) model.HttpResponseResult {

	lrcUrl := "https://api.karadoremi.net/api/song/get_original_music_url.php"
	lrcParam := "device_id=999&platform=8&user_id=" + doremiLogin.UserId +
		"&session_id=" + doremiLogin.SessionId + "&timestamp=" + doremiLogin.Timestamp +
		"&song_id=" + strconv.Itoa(sourceId) + "&client_version=0.0.0&service=2"
	return util.HttpsPost(lrcUrl, []byte(lrcParam))

}
