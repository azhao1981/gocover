package gocover

import (
	"strings"

	"github.com/go-ego/gpy"
)

// 句子顺序覆盖率
// sentence1 被查询
// sentence2 被覆盖
// segTag 分隔符
// minLen 最小长度，太小的不算
func OrderCover(sentence1 string, sentence2 string, segTag string, minLen int) (rate float64) {
	if sentence1 == sentence2 {
		return 1.0
	}
	if strings.TrimSpace(sentence1) == "" || strings.TrimSpace(sentence2) == "" {
		return 0
	}
	arr1 := strings.Split(sentence1, segTag)
	arr2 := strings.Split(sentence2, segTag)
	if len(arr1) < minLen || len(arr2) < minLen {
		return 0
	}
	coverRate := &CoverRate{
		Arr1: arr1,
		Arr2: arr2,
	}
	coverRate.Count()
	return coverRate.Arr2Rate()
}

// 顺序覆盖率，返回被覆盖的句子和句尾无效的句子
func OrderCoverWithUnlessEnd(
	sentence1 string, sentence2 string, segTag string, minLen int,
) (rate float64, sen1 string, sen2 string) {
	arr1 := strings.Split(sentence1, segTag)
	arr2 := strings.Split(sentence2, segTag)
	if len(arr1) < minLen || len(arr2) < minLen {
		return 0, "", sentence1
	}
	coverRate := &CoverRate{
		Arr1: arr1,
		Arr2: arr2,
	}
	coverRate.Count()
	rate = coverRate.Arr2Rate()
	if len(coverRate.HitArr1) == 0 {
		return rate, "", sentence1
	}
	endIndex := coverRate.HitArr1[len(coverRate.HitArr1)-1][1]
	sen1 = strings.Join(coverRate.Arr1[:endIndex+1], segTag)
	if endIndex < len(coverRate.Arr1)-1 {
		sen2 = strings.Join(coverRate.Arr1[endIndex+1:], segTag)
	}
	return rate, sen1, sen2
}

// 顺序覆盖率，返回被覆盖的句子和句尾无效的句子，先转成拼音再比覆盖
func OrderCoverWithPyUnlessEnd(
	sentence1 string, sentence2 string, segTag string, minLen int,
) (rate float64, sen1 string, sen2 string) {
	arr1 := strings.Split(sentence1, segTag)
	arr2 := strings.Split(sentence2, segTag)
	if len(arr1) < minLen || len(arr2) < minLen {
		return 0, "", sentence1
	}
	arr1Py := make([]string, len(arr1))
	arr2Py := make([]string, len(arr2))
	for i, v := range arr1 {
		arr1Py[i] = HasToPY(v)
	}
	for i, v := range arr2 {
		arr2Py[i] = HasToPY(v)
	}
	coverRate := &CoverRate{
		Arr1: arr1Py,
		Arr2: arr2Py,
	}
	coverRate.Count()
	rate = coverRate.Arr2Rate()
	if len(coverRate.HitArr1) == 0 {
		return rate, "", sentence1
	}
	endIndex := coverRate.HitArr1[len(coverRate.HitArr1)-1][1]
	sen1 = strings.Join(arr1[:endIndex+1], segTag)
	if endIndex < len(arr1)-1 {
		sen2 = strings.Join(arr1[endIndex+1:], segTag)
	}
	return rate, sen1, sen2
}

var GpyArgs = gpy.NewArgs()

func HasToPY(sentence string) string {
	py := gpy.Pinyin(sentence, GpyArgs)
	if len(py) == 0 {
		return ""
	}
	pyStr := make([]string, len(py))
	for i, v := range py {
		pyStr[i] = strings.Join(v, " ")
	}
	return strings.Join(pyStr, " ")
}
