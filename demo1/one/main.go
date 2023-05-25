package main

import (
	"fmt"
	"strings"
)

func main() {
	url := "https: //www.youtube.com/live/LHCTW4pckDo?feature=share"

	fmt.Println(ExtractYoutubeId(url))
}
func ExtractYoutubeId(videoRealUrl string) string {
	videoId := ""
	if strings.Contains(videoRealUrl, "get_video_info") {
		urlArr := strings.Split(videoRealUrl, "=")
		if len(urlArr) > 0 {
			videoId = urlArr[len(urlArr)-1]
		}
	} else if strings.Contains(videoRealUrl, "youtu.be") {
		urlArr := strings.Split(videoRealUrl, "/")
		if len(urlArr) > 0 {
			videoId = urlArr[len(urlArr)-1]
		}
	} else if strings.Contains(videoRealUrl, "watch?v=") {
		urlArr := strings.Split(videoRealUrl, "v=")
		if len(urlArr) > 0 {
			videoId = urlArr[len(urlArr)-1]
			if strings.Contains(videoId, "&") {
				urlArr = strings.Split(videoId, "&")
				if len(urlArr) > 0 {
					videoId = urlArr[0]
				}
			}
		}
	} else if strings.Contains(videoRealUrl, "live") { //录播链接
		split := strings.Split(videoRealUrl, "/live/")
		contains := strings.Split(split[1], "?")
		videoId = contains[0]
	}
	return videoId
}
