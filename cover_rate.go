package gocover

import "strings"

type CoverRate struct {
	Arr1      []string
	Arr2      []string
	Arr1Index int // 当前下标
	Arr2Index int // 当前下标

	HitArr1   [][]int // 命中位置下标
	HitArr2   [][]int // 命中位置下标
	Hited     bool    // 是否前面已经有命中
	LastHited bool    // 最后一个是否命中状态

	LastArr1HitIndex int // 最后一个命中开始下标
	LastArr2HitIndex int // 最后一个命中开始下标
}

func (c *CoverRate) Arr2Rate() (rate float64) {
	start, mid, end := c.Arr2LenCount()
	return float64(len(c.Arr2)-start-mid-end) / float64(len(c.Arr2))
}

func (c *CoverRate) Arr1Rate() (rate float64) {
	start, mid, end := c.Arr1LenCount()
	return float64(len(c.Arr1)-start-mid-end) / float64(len(c.Arr1))
}

func (c *CoverRate) Count() error {
	x := 0
	for _, item := range c.Arr2 {
		// 忽略空或空白字符
		if strings.TrimSpace(item) == "" {
			c.AddNotHit()
			continue
		}
		var ok bool
		ok, x = ItemAfter(c.Arr1, item, c.Arr1Index)
		if !ok {
			c.AddNotHit()
			continue
		}
		c.Arr1Index = x
		c.AddHit()
	}
	c.EndCount()
	return nil
}

func (c *CoverRate) EndCount() {
	if c.LastHited {
		c.HitArr1 = append(c.HitArr1, []int{c.LastArr1HitIndex, c.Arr1Index})
		c.HitArr2 = append(c.HitArr2, []int{c.LastArr2HitIndex, c.Arr2Index - 1})
	}
}

func (c *CoverRate) Arr1LenCount() (start int, mid int, end int) {
	return c.LenCount(c.Arr1, c.HitArr1)
}
func (c *CoverRate) Arr2LenCount() (start int, mid int, end int) {
	return c.LenCount(c.Arr2, c.HitArr2)
}

func (c *CoverRate) LenCount(arr []string, hiArr [][]int) (start int, mid int, end int) {
	if len(hiArr) == 0 {
		return 0, 0, 0
	}
	start = hiArr[0][0]
	end = len(arr) - hiArr[len(hiArr)-1][1] - 1
	for idx := range hiArr {
		if idx == 0 {
			continue
		}
		mid += hiArr[idx][0] - hiArr[idx-1][1] - 1
	}
	return
}

func (c *CoverRate) AddNotHit() {
	// fmt.Println("AddNotHit", c.Arr1[c.Arr1Index], c.Arr1Index, c.Arr2[c.Arr2Index], c.Arr2Index)

	if c.LastHited {
		c.HitArr1 = append(c.HitArr1, []int{c.LastArr1HitIndex, c.Arr1Index})
		c.HitArr2 = append(c.HitArr2, []int{c.LastArr2HitIndex, c.Arr2Index - 1})
	}
	c.LastHited = false
	c.Arr2Index++
}

func (c *CoverRate) AddHit() {
	// fmt.Println("AddHit", c.Arr1[c.Arr1Index], c.Arr1Index, c.Arr2[c.Arr2Index], c.Arr2Index)
	if !c.Hited {
		c.Hited = true
	}
	if !c.LastHited {
		c.LastArr1HitIndex = c.Arr1Index
		c.LastArr2HitIndex = c.Arr2Index
	}
	c.LastHited = true
	c.Arr2Index++
}

// 判断 item 是否在 arr 中，并且位置在 index 或之后, i是命中的位置下标
func ItemAfter(arr []string, item string, index int) (ok bool, i int) {
	ok = false
	i = -1

	// 确保 index 在有效范围内
	if index < 0 {
		index = 0
	}
	if index >= len(arr) {
		return false, -1
	}

	// 从 index 开始遍历数组
	for j := index; j < len(arr); j++ {
		if arr[j] == item {
			ok = true
			i = j
			return
		}
	}
	return
}
