package gocover

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrderCoverWithUnlessEnd(t *testing.T) {
	sentence := "广州市大都会大厦3楼"
	sentence2 := "大都会大厦"
	segTag := ""
	minLen := 1
	rate, sen1, sen2 := OrderCoverWithUnlessEnd(sentence, sentence2, segTag, minLen)
	fmt.Printf("rate: [%f] sen1: [%s] sen2: [%s]]\n", rate, sen1, sen2)
	assert.Equal(t, rate, 1.0, "rate")
}

func TestOrderCover(t *testing.T) {
	sentence := "广州市大都会大厦"
	sentence2 := "大都会大厦"
	segTag := ""
	minLen := 1
	rate := OrderCover(sentence, sentence2, segTag, minLen)
	fmt.Println(rate)
	assert.Equal(t, rate, 1.0, "rate")
}

func TestOrderCover2(t *testing.T) {
	sentence := "广州市大都会大厦"
	sentence2 := "在大都会大厦"
	segTag := ""
	minLen := 1
	rate := OrderCover(sentence, sentence2, segTag, minLen)
	fmt.Println(rate)
	assert.Equal(t, rate > 0.8, true, "rate")
}

func TestCoverRate(t *testing.T) {
	time.Sleep(10 * time.Microsecond)
	arr1 := strings.Split("广州市大都会大厦3楼", "")
	arr2 := strings.Split("我到大都会的大厦了", "")
	coverRate := &CoverRate{
		Arr1: arr1,
		Arr2: arr2,
	}
	coverRate.Count()

	fmt.Println(coverRate)
	fmt.Println(coverRate.HitArr1)
	fmt.Println(coverRate.Arr1[len(coverRate.Arr1)-1])
	fmt.Println(coverRate.HitArr2)
	fmt.Println(coverRate.Arr2[len(coverRate.Arr2)-1])

	assert.Equal(t, [][]int{{3, 5}, {6, 7}}, coverRate.HitArr1)
	assert.Equal(t, [][]int{{2, 4}, {6, 7}}, coverRate.HitArr2)

	start, mid, end := coverRate.Arr1LenCount()
	assert.Equal(t, start, 3, "arr1 start")
	assert.Equal(t, mid, 0, "arr1 mid")
	assert.Equal(t, end, 2, "arr1 end")
	start, mid, end = coverRate.Arr2LenCount()
	assert.Equal(t, start, 2, "arr2 start")
	assert.Equal(t, mid, 1, "arr2 mid")
	assert.Equal(t, end, 1, "arr2 end")
}
