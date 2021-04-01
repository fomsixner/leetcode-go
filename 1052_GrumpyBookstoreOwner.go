package leetcode_go

func maxSatisfied(customers []int, grumpy []int, X int) int {
	n, ans := len(customers), 0
	add := 0
	for i := 0; i < n; i++ {
		for i < n && grumpy[i] == 0 {
			ans += customers[i]
			i++
		}
		temp := 0
		for left, right := i, i+X; left < right && left < n; left++ {
			if grumpy[left] == 1 {
				temp += customers[left]
			}
		}
		if temp > add {
			add = temp
		}
	}
	return ans + add
}
