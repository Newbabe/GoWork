package util

import (
	"github.com/semicircle/gozhszht"
	"github.com/siongui/gojianfan"
	"strings"
)

const SIMPLIFED_CHINESE = "zh-Hans"
const TRADITIONAL_CHINESE = "zh-Hant"

func Transform(content string, appLanguage string) string {
	if content == "" {
		return content
	}
	if SIMPLIFED_CHINESE == appLanguage {
		return gozhszht.ToSimple(content)
	}
	if strings.Contains(content, "没") {
		content = strings.Replace(content, "没", "沒", -1)
	}
	if strings.Contains(content, "伦") {
		strings.Replace(content, "伦", "倫", -1)
	}
	return gozhszht.ToTradition(content)
}

func TransformChinese(text string, appLanguage string) string {
	str := gojianfan.S2T(text)
	if appLanguage == SIMPLIFED_CHINESE {
		str = gojianfan.T2S(text)
	}
	str = strings.Replace(str, "回復", "回覆", -1)
	if str == "周傑倫" {
		str = "周杰倫"
	}
	if strings.Contains(str, "癡") {
		str = strings.Replace(str, "癡", "痴", -1)
	}
	if strings.Contains(str, "範") {
		str = strings.Replace(str, "範", "范", -1)
	}
	if strings.Contains(str, "誌") {
		str = strings.Replace(str, "誌", "志", -1)
	}

	return str
}
