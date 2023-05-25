package model

import "time"

/*private int sourceId;
private boolean duet;
private String updateTime;
private String lrcChannel;
private int youtubeSongType;*/

type DoremiSong struct {
	SourceId        int
	YoutubeSongType int
	Duet            bool
	UpdateTime      time.Time
	LrcChannel      string
}
