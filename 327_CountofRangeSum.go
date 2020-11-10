package leetcode_go


//方法一: 暴力破解
func countRangeSum(nums []int, lower int, upper int) int {
	res, n:= 0, len(nums)
	for i:=0; i < n; i++ {
		sum := 0
		for j:=i; j < n; j++ {
			sum += nums[j]
			if sum >= lower && sum <= upper {
				res++
			}
		}
	}
	return res
}


//
//func main()  {
//	nums := []int{-2,5,-1}
//	fmt.Println(countRangeSum(nums,-2,2))
//}