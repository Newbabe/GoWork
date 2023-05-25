package dao

import (
	"SongSyncProgram/model"
	"SongSyncProgram/util"
	"github.com/tidwall/gjson"
)

func GetDoremiLogin() model.DoremiLogin {
	var doremiLogin model.DoremiLogin
	loginUrl := "https://api.karadoremi.net/api/member/secure/login.php"
	loginParam := "device_id=999&platform=8&client_version=0.0.0&username=service@goodmusic-corp.com&password=7FD1C2FE6474&service=3"
	result := util.HttpsPost(loginUrl, []byte(loginParam))
	jsonStr := string(result.Bytes)
	status := gjson.Get(jsonStr, "status").String()
	if status == "success" {
		doremiLogin.SessionId = gjson.Get(jsonStr, "session_id").String()
		doremiLogin.Timestamp = gjson.Get(jsonStr, "timestamp").String()
		doremiLogin.UserId = gjson.Get(jsonStr, "user_id").String()
	}
	return doremiLogin
}

func GetDoremiLogin2() model.DoremiLogin {
	var doremiLogin model.DoremiLogin
	loginUrl := "https://api.karadoremi.net/api/member/secure/login.php"
	loginParam := "device_id=999&platform=8&client_version=0.0.0&username=service@goodmusic-corp.com&password=7FD1C2FE6474&service=3"
	result := util.HttpsPost(loginUrl, []byte(loginParam))
	bytes := string(result.Bytes)
	status := gjson.Get(bytes, "status").String()
	if status == "success" {
		doremiLogin.SessionId = gjson.Get(bytes, "session_id").String()
		doremiLogin.Timestamp = gjson.Get(bytes, "timestamp").String()
		doremiLogin.UserId = gjson.Get(bytes, "user_id").String()
		return doremiLogin
	}
	return doremiLogin
}
