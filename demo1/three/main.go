package main

import (
	"fmt"
	"github.com/ozgio/strutil"
	"github.com/shmokmt/go-zhuyin"
	"strings"
)

func main() {
	/*Some (MV)
	Some
	Some (Karaoke)
	不能講的秘密 (純伴奏)
	為你放袂落 (MV)
	真愛只有你
	恨你無情 (MV)
	恨你無情 (純伴奏)
	情人淚 (MV)
	情人淚 (純伴奏)
	刻字
	刻字 (純伴奏)
	許容欣,陳芝緗,陳韋慈=ㄒㄖㄒ,ㄔㄓㄒ,ㄔㄨㄘ
	彼得潘(純伴奏)=ㄅㄉㄆ(ㄔㄅㄗ)
	離別的酒(純伴奏)=ㄌㄅㄉㄐ(ㄔㄅㄗ)
	亞特蘭提斯=ㄧㄊㄌㄊㄙ
	是什麼讓我遇見這樣的你(純伴奏)=ㄕㄕㄇㄖㄨㄩㄐㄓㄧㄉㄋ(ㄔㄅㄗ)
	*/
	yinNew := GetZhuYinNew("許容欣,陳芝緗,陳韋慈")
	//s2 := zhuyin.Convert("情人淚 (純伴奏)")
	fmt.Println(yinNew)
}
func GetZhuYinNew(word string) string {

	var newWord string
	//处理歌名空格
	newWord = strings.ReplaceAll(word, " ", "")
	//fmt.Println("处理空格后的歌名:", newWord)
	//处理歌名特殊字符和字母数字
	arr2 := []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "j", "k", "l",
		"z", "x", "c", "v", "b", "n", "m", "\\+", "\\-", "\\_", "1", "\\.", "\\'", "2", "3", "4", "5", "6", "7", "8", "9", "0", "'"}

	for _, s := range arr2 {
		newWord = strings.ReplaceAll(newWord, s, "")
		newWord = strings.ReplaceAll(newWord, strings.ToUpper(s), "")
	}
	newWord = strings.ReplaceAll(newWord, "()", "")
	var s2 []string

	for _, s := range Str(newWord) {
		if s == "(" || s == ")" || s == "," {
			s2 = append(s2, s)
			continue
		}
		s2 = append(s2, zhuyin.Convert(s))

	}

	var content string
	if len(Str(newWord)) != len(s2) {
		return ""
	}
	for i := 0; i < len(Str(newWord)); i++ {
		if i == len(Str(newWord)) { //跳过最后一次
			break
		}
		arrStr, _ := strutil.Substring(s2[i], 0, 1)
		content += arrStr
	}

	return content

}

// Str string="原子少年"==》[]String{"原","子","少","年"}
func Str(str string) []string {
	sp := make([]string, 0)
	for _, v := range str {
		if string(v) == " " { //去除空格
			continue
		}
		if string(v) == "\"" {
			continue
		}
		sp = append(sp, string(v))
	}
	return sp
}
