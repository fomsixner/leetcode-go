package leetcode_go

func canCompleteCircuit(gas []int, cost []int) int {
	res := -1
	for i := 0; i < len(gas); i++ {
		cur, j := 0, 0
		for ; j < len(gas); j++ {
			if cur+gas[(i+j)%len(gas)]-cost[(i+j)%len(gas)] < 0 {
				break
			}
			cur = cur + gas[(i+j)%len(gas)] - cost[(i+j)%len(gas)]
		}
		if j == len(gas) {
			res = i
			break
		}
	}
	return res
}

//方法二: 一次遍历
//我们首先检查第 00 个加油站，并试图判断能否环绕一周；如果不能，就从第一个无法到达的加油站开始继续检查。

//func main()  {
//	fmt.Println(canCompleteCircuit([]int{3,3,4}, []int{3,4,4}))
//}
