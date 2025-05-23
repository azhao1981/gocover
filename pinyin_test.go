package gocover

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-ego/gpy"
	"github.com/stretchr/testify/assert"
)

func TestOrderCoverWithPYUnlessEnd(t *testing.T) {
	sentence := "广州市大都会大厦3楼"
	sentence2 := "大都会打厦"
	segTag := ""
	minLen := 1
	rate, sen1, sen2 := OrderCoverWithPyUnlessEnd(sentence, sentence2, segTag, minLen)
	fmt.Printf("rate: [%f] sen1: [%s] sen2: [%s]]\n", rate, sen1, sen2)
	assert.True(t, rate > 0.8, "rate")

	rate, sen1, sen2 = OrderCoverWithUnlessEnd(sentence, sentence2, segTag, minLen)
	fmt.Printf("rate: [%f] sen1: [%s] sen2: [%s]]\n", rate, sen1, sen2)
	assert.True(t, rate >= 0.8, "rate")
}

func TestOrderCoverWithPY(t *testing.T) {
	panicArr := []string{
		"附近",
		"局",
		"门",
		"侧约",
		"交叉",
		"街道办",
		"24小时",
	}
	sentences := []string{
		"山东省潍坊市奎文区高新二路与桃园街交叉路口东北侧潍坊景泰园",
		"山东省潍坊市奎文区中心路潍坊景泰园东北侧约40米潍坊综合保税区信访局",
		"山东省潍坊市奎文区中心路潍坊景泰园东北侧约70米潍坊综合保税区24小时智慧警局",
		"山东省潍坊市奎文区桃园街景泰园潍坊景泰园55号楼",
		"山东省潍坊市奎文区高新二路与桃园街交叉路口东北侧潍坊景泰园北门",
	}
	sentence := "山东省潍坊市保税区香田园镇保税区景泰园"
	segTag := ""
	minLen := 3
	for _, s := range sentences {
		rate, sen1, sen2 := OrderCoverWithPyUnlessEnd(s, sentence, segTag, minLen)
		panic := 0.0
		for _, p := range panicArr {
			if strings.Contains(s, p) && !strings.Contains(sentence, p) {
				panic += 0.15
			}
		}
		fmt.Printf("rate: [%f][%f] [%f] sen1: [%s] sen2: [%s]]\n", rate-panic, rate, panic, sen1, sen2)
		assert.True(t, rate-panic > 0.1, "rate")
	}
}

func TestPinyin(t *testing.T) {
	sentence := "广州市大都会大厦14楼"
	a := gpy.NewArgs()
	fmt.Println(gpy.Pinyin(sentence, a))

}
func TestPinyin2(t *testing.T) {
	sentence := "广州市大都会大厦十四楼"
	a := gpy.NewArgs()
	fmt.Println(gpy.Pinyin(sentence, a))
}
