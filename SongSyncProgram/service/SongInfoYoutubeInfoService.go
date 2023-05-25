package service

import (
	"SongSyncProgram/dao"
	"SongSyncProgram/model"
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"github.com/tidwall/gjson"
	"regexp"
	"strconv"
	"strings"
)

func GetAdjustmentValue(typeInt int) int {
	return dao.GetAdjustmentValue(typeInt)
}
func AddSongInfoYoutubeInfo(SourceId int, YouTubeId string, videoStartTimeOffset string) {
	//fmt.Println("添加SongInfoYoutubeInfo表")
	dao.AddSongInfoYoutubeInfo(SourceId, YouTubeId, videoStartTimeOffset)
}

//清空YoutubeOffest"
func ClearYoutubeOffest() {
	dao.ClearYoutubeOffest()
}

func getChannel(SexTagStr string) int {
	sexTag, _ := strconv.Atoi(SexTagStr)
	lrcChannel := 1
	if (sexTag & 32768) == 32768 {
		lrcChannel = 3
	} else if (sexTag & 16384) == 16384 {
		lrcChannel = 7
	} else {
		lrcChannel = 5
	}
	return lrcChannel
}

//判断歌曲是否存在
//如果存在就将该条数据返回
func GetIsExt(SongInfoMergeList []model.SongInfoMerger, sourceId int) (bool, model.SongInfoMerger) {
	var flag bool
	var Song model.SongInfoMerger
	for _, SongInfoMerge := range SongInfoMergeList {
		if SongInfoMerge.SourceId == sourceId {

			flag = true
			//	fmt.Println("flag", flag, sourceId, SongInfoMerge.SourceId)
			return flag, SongInfoMerge
		}

	}
	return flag, Song
}

func getIsDuet(duetModeEnabledStr string) int {
	var isDuet int
	duetModeEnabled, _ := strconv.Atoi(duetModeEnabledStr)
	if duetModeEnabled == 1 {
		isDuet = 1
	}
	return isDuet
}

func getYouTubeSongType(songName string) int {

	youtubeSongType := 1 // youtube伴奏类型1练唱，2MV,3纯伴奏
	if strings.Contains(songName, "(MV)") || strings.Contains(songName, "（MV）") {
		youtubeSongType = 2
	} else if strings.Contains(songName, "(純伴奏)") || strings.Contains(songName, "（純伴奏）") || strings.Contains(songName, "(Karaoke)") || strings.Contains(songName, "（Karaoke）") {
		youtubeSongType = 3
	} else if strings.Contains(songName, "(練唱版)") {
		youtubeSongType = 4
	}
	return youtubeSongType
}

func getKeyWord(songIndex, singerIndex, songPhoneticIndex, singerPhoneticIndex string) string {
	var keyword string
	if songIndex != "" {
		keyword += songIndex
	}
	if singerIndex != "" {
		keyword += " " + singerIndex
	}

	if songPhoneticIndex != "" {
		keyword += " " + songPhoneticIndex
	}
	if singerPhoneticIndex != "" {
		keyword += " " + singerPhoneticIndex
	}
	return keyword
}

func getSex(sexTagStr string) int {
	sexTag, _ := strconv.Atoi(sexTagStr)
	var sex int
	if (sexTag&1) != 1 && (sexTag&2) != 2 {
		if (sexTag & 4) == 4 {
			sex = 3
		} else {
			sex = 0
		}
	}
	if (sexTag & 1) == 1 {
		sex = 1
	}
	if (sexTag & 2) == 2 {
		sex = 2
	}
	return sex
}

func getLanguage(languageStr string) int {
	language, _ := strconv.Atoi(languageStr)
	if language == 1 || language == 2 {
		language = 1
	} else if language == 3 {
		language = 5
	} else if language == 4 {
		language = 4
	} else if language == 5 {
		language = 2
	} else if language == 6 {
		language = 7
	} else if language == 7 {
		language = 18
	} else if language == 8 {
		language = 20
	} else if language == 9 {
		language = 3
	} else if language == 10 {
		language = 19
	} else if language == 11 {
		language = 21
	}
	return language
}

func getLrcSizeAndLrc2Size(hasLrc int, duetModeEnabledStr string) (int, int, int) {
	duetModeEnabled, _ := strconv.Atoi(duetModeEnabledStr)

	var lrcSize int
	var lrc2Size int
	if hasLrc == 1 || duetModeEnabled == 1 {
		hasLrc = 0
		lrcSize = 1
	} else {
		hasLrc = 4
	}
	if duetModeEnabled == 1 {
		lrc2Size = 1
	}
	return lrcSize, lrc2Size, hasLrc
}

func SongNamePinYinProcessing(songName string, songChinesePhoneticArray []string) string {
	//歌名处理判断是否为纯伴奏等歌曲
	index := JudgmentStringIndex(songName)
	b, Commaindex := JudgmentCommaindex(songName)
	//fmt.Println("111", songChinesePhoneticArray)
	//fmt.Println("index", index)
	var songChinesePhoneticStr string
	if len(songChinesePhoneticArray) == 0 { //英文歌曲
		return songName
	}
	if index == 0 { //不是其他版本的歌曲不需要对歌曲首字母做处理
		for i, pingYin := range songChinesePhoneticArray {
			if b && i == Commaindex-1 {
				songChinesePhoneticStr += pingYin + ","
				continue
			}
			songChinesePhoneticStr += pingYin

		}
	} else {
		if len(songChinesePhoneticArray)-index <= 0 { //mv版本
			for _, pingYin := range songChinesePhoneticArray {

				songChinesePhoneticStr += pingYin
			}
			songChinesePhoneticStr += "(MV)"
		}
		if len(songChinesePhoneticArray)-1-index > 0 { //纯伴奏版本
			for i, pingYin := range songChinesePhoneticArray {

				if i == index {
					songChinesePhoneticStr += "(" + pingYin
					continue
				}
				if len(songChinesePhoneticArray)-1-i == 0 {
					songChinesePhoneticStr += pingYin + ")"
					continue
				}
				songChinesePhoneticStr += pingYin
			}
		}
	}

	return songChinesePhoneticStr
}
func SongNamePinYinProcessing2(songName string) string {
	var SongNamePinYin string
	var suffixPinYin1 string
	var suffixPinYin2 string
	typePinYin := pinyin.NewArgs()
	typePinYin.Style = pinyin.FirstLetter //首字母

	////处理【你聽見我了沒有_Del】情况
	songName = strings.ReplaceAll(songName, "_Del", "地")
	arr := pinyin.LazyPinyin(songName, typePinYin)
	//fmt.Println("arr", arr)
	if len(arr) == 0 { //纯英语或者纯日语直接返回歌名
		return songName
	}

	//去除歌名的特殊符号
	//1.去除所有特殊字符 除了空格符号
	punctuateStrList := []string{"!", "！", "，", "'", "‘", "？", "?", ":", "：",
		"\\", "、", "|", "《", "<", "》", ">", ".", "。", "/", "、", "·", ";", "；", "’",
		"`", "@", "#", "￥", "$", "%", "……", "^", "&", "*", "-", "——", "+", "=", "+"}
	for i := 0; i < len(punctuateStrList); i++ {
		songName = strings.ReplaceAll(songName, punctuateStrList[i], "")
	}
	//处理歌名判断是否是纯伴奏登歌曲
	str1, suffixPinYin1 := ProcessingPinYin(songName)
	if strings.Contains(str1, "(") || strings.Contains(str1, "（") {
		str1, suffixPinYin2 = ProcessingPinYin(str1)
	}
	//处理没有拼音的字符
	typePinYin.Fallback = func(r rune, a pinyin.Args) []string {
		pattern := "[A-Za-z]+$"
		//正则匹配英文【含英语字母的也之间返回歌名】
		reg, _ := regexp.Compile(pattern)
		flag1 := reg.MatchString(string(r))
		if flag1 { //如果匹配直接返回
			return []string{strings.ToUpper(string(r))}
		}
		// u0800-u4e00 日语Unicode编码范围
		//日语rune的返回 >12448
		//英语 65-122
		if r >= 10 { //日语
			return []string{"."}
		}
		return []string{string(r)}
	}
	//判断是否含有"," //纯中文处理逗号[判断是否存在存在返回下标]
	//b, commaIndex := JudgmentCommaindex(songName)
	/*fmt.Println("songName", songName)
	fmt.Println("b", b)
	fmt.Println("commaIndex", commaIndex)*/
	SongNamePinYinArr := pinyin.LazyPinyin(str1, typePinYin)
	//	fmt.Println("SongNamePinYinArr", SongNamePinYinArr)
	for i, s := range SongNamePinYinArr {
		//拼接日语
		if s == "." {
			strArr := Str2(songName)
			s = strArr[i]
		}
		/*if b && i == commaIndex-1 {
			SongNamePinYin += s + ","
			continue
		}*/

		SongNamePinYin += s
	}
	//fmt.Println(SongNamePinYin)
	return SongNamePinYin + suffixPinYin2 + suffixPinYin1
}

//判断”(“的下标位置
func JudgmentStringIndex(str string) int {
	//去除空格
	strArray := Str(str)
	var index int
	for i, s := range strArray {
		//	fmt.Println(i, s)
		if s == "(" || s == "（" {

			index = i
		}
	}
	return index
}

func JudgmentCommaindex(singer string) (bool, int) {
	for i, s := range Str2(singer) {
		if s == "," {
			return true, i
		}
	}
	return false, 0
}
func JudgmentStringIndex2(str string) (int, []string) {
	//保留空格
	strArray := Str2(str)
	var index int
	for i, s := range strArray {
		//	fmt.Println(i, s)
		if s == "(" || s == "（" {
			index = i
		}
	}
	return index, strArray
}
func JudgmentStringIndex3(str string) (int, []string) {
	//去除空格
	strArray := Str(str)
	var index int
	for i, s := range strArray {
		//	fmt.Println(i, s)
		if s == "(" || s == "（" {
			index = i
		}
	}
	return index, strArray
}

//处理字段工具
func GetString(column, jsonStr string) string {
	//字段为String的类型还需要转码
	//获取字段Unicode编码
	//获取songlist数组里面的数据
	UnicodeStr := gjson.Get(jsonStr, column).String()
	//处理编码问题
	unicode, err := zhToUnicode([]byte(delBrackets(UnicodeStr)))
	if err != nil {
		fmt.Println("编码转换异常")
		return ""
	}
	return unicode
}
func GetInt(column, jsonStr string) string {
	//int类型的字段是需要去除。中阔号
	UnicodeStr := gjson.Get(jsonStr, column).String()
	return delBrackets(UnicodeStr)
}

//去除获取的字符串中的中括号
func delBrackets(str string) string {
	UnicodeStr1 := strings.ReplaceAll(str, "[", "")
	return strings.ReplaceAll(UnicodeStr1, "]", "")
}

//将Unicode编码转为汉字
func zhToUnicode(raw []byte) (string, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return "", err
	}
	return str, nil
}

func GetWordPart(word string) int {
	var songNameStr string
	songNameStr = ProcessingBroadHorn(word)
	//fmt.Println("songNameStr333=", songNameStr)
	if strings.Contains(songNameStr, "(") {
		songNameStr = ProcessingBroadHorn(songNameStr)
	}
	//fmt.Println("songNameStr444=", songNameStr)
	//1.去除所有特殊字符 除了空格符号
	punctuateStrList := []string{"!", "！", ",", "，", "'", "‘", "？", "?", ":", "：",
		"\\", "、", "|", "《", "<", "》", ">", ".", "。", "/", "、", "·", ";", "；", "’",
		"`", "@", "#", "￥", "$", "%", "……", "^", "&", "*", "-", "——", "+", "=", "+"}
	for i := 0; i < len(punctuateStrList); i++ {
		songNameStr = strings.ReplaceAll(songNameStr, punctuateStrList[i], "")
	}
	//处理与英语混合的歌曲
	//fmt.Println("处理后的word", songNameStr)
	//fmt.Println("长度", len(Str(songNameStr)))
	//2.定义两个临时变量
	var songTemp string
	var partNum int
	pattern := "[A-Za-z]+$"
	pattern2 := "\\s+"
	for i, str := range Str2(songNameStr) {

		//	fmt.Println("str", str)
		//fmt.Println("str:", string(str))
		r, _ := regexp.Compile(pattern)
		flag1 := r.MatchString(str)
		r2, _ := regexp.Compile(pattern2)
		flag2 := r2.MatchString(str)
		//	fmt.Printf("flag1=%t   flag2=%t \n", flag1, flag2)
		if flag1 { //正则匹配英文字母
			songTemp += str
			//		fmt.Println("songTemp", songTemp)
			if i == len(Str2(songNameStr))-1 {
				partNum += 1
			}
		} else if flag2 {
			//		fmt.Println(22222)
			//正则表达式匹配空格符号
			if str != "" && len(songTemp) > 0 {
				partNum += 1
				songTemp = ""
			}
		} else {
			//	fmt.Println(3)
			//fmt.Println(songTemp)
			//汉字,数字,日文 字部 +1
			if len(songTemp) > 0 {
				partNum += 1
			}
			partNum += 1
			songTemp = ""
		}
	}
	return partNum
}

func ProcessingBroadHorn(word string) string {
	var songNameStr string //处理阔号前面的字符串
	var KuoHao string
	if strings.Contains(word, "(") || strings.Contains(word, "（") {
		index, nameList := JudgmentStringIndex2(word)
		/*fmt.Println("-----", nameList)
		fmt.Println("index==", index)
		fmt.Println("Len", len(nameList))*/
		for i := 0; i <= index; i++ {
			if i >= 1 { //消除” （）“对结果的影响
				songNameStr += nameList[i-1]
				//fmt.Println(songNameStr)
			}
		}
		for i := index + 1; i < len(nameList); i++ {
			KuoHao += nameList[i]
		}
	} else {
		songNameStr = word
	}
	KuoHao = strings.ReplaceAll(KuoHao, ")", "")
	/*	fmt.Println(1, songNameStr)
		fmt.Println(2, KuoHao)*/
	//对阔号里面的字符串做判断如果是 （mv） （纯伴奏） 等歌曲就无需拼接字符
	if strings.Contains(KuoHao, "Karaoke") || strings.Contains(KuoHao, "MV") ||
		strings.Contains(KuoHao, "純伴奏") || strings.Contains(KuoHao, "单") ||
		strings.Contains(KuoHao, "A Happier Song") || strings.Contains(KuoHao, "Remix") ||
		strings.Contains(KuoHao, "On the Bayou") || strings.Contains(KuoHao, "Taylor’s Version") ||
		strings.Contains(KuoHao, "現場版") || strings.Contains(KuoHao, "A Man After Midnight") ||
		strings.Contains(KuoHao, "韓") || strings.Contains(KuoHao, "台") || strings.Contains(KuoHao, "日") ||
		strings.Contains(KuoHao, "客") || strings.Contains(KuoHao, "短版") || strings.Contains(KuoHao, "粵語") ||
		strings.Contains(KuoHao, "Japanese Version") || strings.Contains(KuoHao, "DJ名龍版") || strings.Contains(KuoHao, "副歌") ||
		strings.Contains(KuoHao, "Stupid Mistake") || strings.Contains(KuoHao, "From DreamWorks Animation's Trolls") ||
		strings.Contains(KuoHao, "高雄觀光主題曲") || strings.Contains(KuoHao, "低音聲部練習版") || strings.Contains(KuoHao, "吹管聲部練習版") ||
		strings.Contains(KuoHao, "彈撥聲部練習版") || strings.Contains(KuoHao, "拉弦聲部練習版") || strings.Contains(KuoHao, "抒情版") {
		songNameStr += ""
	} else {
		songNameStr += " " + KuoHao //【注意拼接空格】
	}
	return songNameStr
}
func ProcessingPinYin(word string) (string, string) {
	var songNameStr string //处理阔号前面的字符串
	var KuoHao string
	if strings.Contains(word, "(") || strings.Contains(word, "（") {
		index, nameList := JudgmentStringIndex3(word)
		for i := 0; i <= index; i++ {
			if i >= 1 { //消除” （）“对结果的影响
				songNameStr += nameList[i-1]
				//fmt.Println(songNameStr)
			}
		}
		for i := index + 1; i < len(nameList); i++ {
			KuoHao += nameList[i]
		}
	} else {
		songNameStr = word
	}
	var PinYin string
	KuoHao = strings.ReplaceAll(KuoHao, ")", "")
	if strings.Contains(KuoHao, "MV") {
		KuoHao = "(MV)"
	} else if strings.Contains(KuoHao, "純伴奏") {
		KuoHao = "(cbz)"
	} else if strings.Contains(KuoHao, "Karaoke") {
		KuoHao = "(Karaoke)"
	} else if strings.Contains(KuoHao, "DJ名龍版") {
		typePinYin := pinyin.NewArgs()
		typePinYin.Style = pinyin.FirstLetter //首字母
		SongNamePinYinArr := pinyin.LazyPinyin(KuoHao, typePinYin)
		for i, s := range SongNamePinYinArr {
			if i == 0 {
				PinYin += "(" + "DJ" + s
				continue
			}
			if i == len(SongNamePinYinArr)-1 {
				PinYin += s + (")")
				continue
			}
			PinYin += s
		}
		return songNameStr, PinYin
	} else { //不是以上情况就
		typePinYin := pinyin.NewArgs()
		typePinYin.Style = pinyin.FirstLetter //首字母
		SongNamePinYinArr := pinyin.LazyPinyin(KuoHao, typePinYin)
		for i, s := range SongNamePinYinArr {
			if len(SongNamePinYinArr) == 1 {
				PinYin += "(" + s + ")"
				continue
			}
			if i == 0 {
				PinYin += "(" + s
				continue
			}
			if i == len(SongNamePinYinArr)-1 {
				PinYin += s + (")")
				continue
			}
			PinYin += s
		}
		return songNameStr, PinYin
	}
	return songNameStr, KuoHao
}
