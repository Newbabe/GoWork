package service

import (
	"SongSyncProgram/util"
	"bufio"
	"fmt"
	"github.com/ozgio/strutil"
	"github.com/shmokmt/go-zhuyin"
	"net/url"
	"os"
	"strings"
)

//测试文件路径
//const cacheFile = "/home/ec2-user/GOTest/MalaysiaTEST/cache_map.txt"

// 120服务器文件路径
const cacheFile = "/home/ec2-user/SongSyncProgram/cache_map.txt"

var cacheMap = make(map[string]string) //定义一个全局map

/*1.读取cacheFile文件
2.读取每一行数据
3.处理数据存入cacheFile中
*/
// TODO 相当于java的static代码块
func init() {
	open, err := os.Open(cacheFile)
	defer open.Close()
	if err != nil {
		fmt.Println("打开文件出错")
		return
	}
	txt, err := util.ReadTxt(open)
	if err != nil {
		return
	}

	for i := 0; i < len(txt)-1; i++ {
		if i == 0 {

		}
		//TODO 此处暂时不用考虑txt数组中的元素只用一个的情况
		str := strings.Split(txt[i], "=")
		cacheMap[str[0]] = str[1]
		if i == len(txt)-2 {
			str = strings.Split(txt[i+1], "=")
			cacheMap[str[0]] = str[1]
		}

	}

}
func UpdateCacheMap() {
	//读取txt文件
	open, err := os.Create(cacheFile)
	defer open.Close()
	if err != nil {
		fmt.Println("创建文件出错")
		return
	}
	//遍历cacheMap集合
	for k, v := range cacheMap {
		//	将cacheMap的数据写入txt文件并且格式为k=v
		writer := bufio.NewWriter(open)
		writer.Write([]byte(k + "=" + v))
		//将缓存中的数据写入文件中
		writer.Flush()
		if err != nil {
			fmt.Println("写入文件失败", err)
			return
		}

	}
}

func GetZhuYin(word string) string {
	cache := cacheMap[word]
	if cache != "" { //如果cache中有改数据直接返回
		return cache
	}
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
	//去掉空（）
	//	fmt.Println("处理特殊符号后的歌名:", newWord)
	newWord = strings.ReplaceAll(newWord, "()", "")
	//fmt.Println("去掉阔号后的歌名", newWord)
	//处理连接
	path := "https://pinyin.thl.tw/converter?s=" + url.QueryEscape(newWord)
	//fmt.Println("链接:", path)
	body := util.HttpGet(path)
	str := "ui very basic selectable table"
	indexs := strings.Index(body, str)
	//fmt.Println("小标", indexs)
	result, _ := strutil.Substring(body, indexs-141, indexs-70)
	result1 := strings.TrimSpace(result)
	//fmt.Println("result1", result1)
	ss := strings.Trim(result1, "\">\n          </td></tr>\n </table>\n        <h3")
	//fmt.Println("ss==", ss)
	s1 := strings.ReplaceAll(ss, "</td></tr>", "")
	s2 := strings.ReplaceAll(s1, "<tr><td>", "")
	//s3 := strings.ReplaceAll(s3, "</table>", "")
	arr := strings.Fields(s2)
	/*fmt.Println("旧歌名长度--:", len(Str(word)))
	fmt.Println("处理后的歌名长度--:", len(Str(newWord)))
	fmt.Println("获取后的注音", arr)*/
	/*for i, s := range arr {
		fmt.Println("获取后的注音遍历：", i, s)
	}*/
	var content string
	if len(Str(newWord)) != len(arr) {
		return ""
	}
	for i := 0; i < len(Str(newWord)); i++ {
		if i == len(Str(newWord)) { //跳过最后一次
			break
		}
		arrStr, _ := strutil.Substring(arr[i], 0, 1)
		content += arrStr
	}
	if content != "" { //如果注音不为空就存入caCheMap中
		cacheMap[word] = content
	}
	return content

}
func GetZhuYinNew(word string) string {
	cache := cacheMap[word]
	if cache != "" { //如果cache中有改数据直接返回
		return cache
	}
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
	if content != "" {
		cacheMap[word] = content
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
func Str2(str string) []string {
	sp := make([]string, 0)
	for _, v := range str {
		if string(v) == "\"" {
			continue
		}
		sp = append(sp, string(v))
	}
	return sp
}
