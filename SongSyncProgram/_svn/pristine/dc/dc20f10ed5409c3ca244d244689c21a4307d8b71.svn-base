package util

import (
	"SongSyncProgram/model"
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) string {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}} //如果需要测试自签名的证书 这里需要设置跳过证书检测 否则编译报错
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)

	if err != nil {
		fmt.Println("error:", err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}
func HttpsPost(url string, data []byte) model.HttpResponseResult {
	var HttpResponseResult model.HttpResponseResult
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	response, err := client.Post(url, "application/x-www-form-urlencoded", bytes.NewBuffer(data))

	if err != nil {
		fmt.Println("链接失败", err)

		HttpResponseResult.Code = response.StatusCode
		return HttpResponseResult
	}
	HttpResponseResult.Code = response.StatusCode
	bodys, _ := ioutil.ReadAll(response.Body)
	HttpResponseResult.Bytes = bodys
	return HttpResponseResult
}
func DownLoadFromUrl(strUrl string) model.HttpResponseResult {
	var result model.HttpResponseResult
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr}
	request, err := http.NewRequest("GET", strUrl, nil)
	request.Header.Set("User-Agent", "Mozilla/4.0(compatible;MSIE 5.0;Windows NT;DigExt)")

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("code", resp.StatusCode)
	if err != nil {
		result.Code = resp.StatusCode
		return result
	}
	bodys, _ := ioutil.ReadAll(resp.Body)
	result.Bytes = bodys
	result.Code = resp.StatusCode
	return result
}
