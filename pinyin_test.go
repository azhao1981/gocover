package gocover

import (
	"fmt"
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
	assert.Equal(t, rate, 1.0, "rate")

	rate, sen1, sen2 = OrderCoverWithUnlessEnd(sentence, sentence2, segTag, minLen)
	fmt.Printf("rate: [%f] sen1: [%s] sen2: [%s]]\n", rate, sen1, sen2)
	assert.Equal(t, rate, 1.0, "rate")
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
