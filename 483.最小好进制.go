package leetcode_go

import (
	"math"
	"math/bits"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=483 lang=golang
 *
 * [483] 最小好进制
 */

// @lc code=start
func smallestGoodBase(n string) string {
	number, _ := strconv.Atoi(n)
	mMax := bits.Len(uint(number)) - 1
	for m := mMax; m > 1; m-- {
		k := int(math.Pow(float64(number), 1/float64(m)))
		mul, sum := 1, 1
		for i := 0; i < m; i++ {
			mul *= k
			sum += mul
		}
		if sum == number {
			return strconv.Itoa(k)
		}
	}
	return strconv.Itoa(number - 1)
}

// @lc code=end
// func main() {
// 	fmt.Println(smallestGoodBase("13"))
// }
