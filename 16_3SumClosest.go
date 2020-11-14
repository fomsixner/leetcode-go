package leetcode_go

import (
	"math"
	"sort"
)

//func abs(x int) int {
//	if x < 0 {
//		return -x
//	}
//	return x
//}

//使用math.Abs比自己写的Abs快很多

//双指针法
//固定一个数,再双指针
func threeSumClosest(nums []int, target int) int {
	n, dist := len(nums), math.MaxInt64
	res := nums[0] + nums[1] + nums[2]
	sort.Ints(nums)
	for i := 0; i <= n-3; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			for j, k := i+1, n-1; j < k; {
				count := nums[i] + nums[j] + nums[k]
				temp := int(math.Abs(float64(nums[i] + nums[j] + nums[k] - target)))
				if temp < dist {
					dist, res = temp, count
				}
				if count == target {
					return res
				} else if count > target {
					for oldk := nums[k]; k > j && nums[k] == oldk; k-- {
					}
				} else {
					for oldj := nums[j]; j < k && nums[j] == oldj; j++ {
					}
				}
			}
		}
	}
	return res
}

//func main()  {
//	nums := []int{1,1,-1,-1,3}
//	fmt.Println(threeSumClosest(nums, -1))
//}
