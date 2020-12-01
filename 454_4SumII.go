package leetcode_go

import (
	"sort"
)

func fourSumCount(A []int, B []int, C []int, D []int) int {
	sort.Ints(A)
	sort.Ints(B)
	sort.Ints(C)
	sort.Ints(D)
	res := 0
	count1, count2 := 0, 0
	for a := 0; a < len(A); a++ {
		if a != 0 && A[a] == A[a-1] {
			res += count1
			continue
		}
		count1 = 0
		for b := 0; b < len(B); b++ {
			if b != 0 && B[b] == B[b-1] {
				res += count2
				count1 += count2
				continue
			}
			count2 = 0
			for c, d := 0, len(D)-1; c < len(C) && d >= 0; {
				if A[a]+B[b]+C[c]+D[d] == 0 {
					x, y := 1, 1
					for c+1 < len(C) && C[c] == C[c+1] {
						c++
						x++
					}
					for d-1 >= 0 && D[d] == D[d-1] {
						d--
						y++
					}
					res += x * y
					count1 += x * y
					count2 += x * y
					c++
				} else if A[a]+B[b]+C[c]+D[d] > 0 {
					d--
				} else {
					c++
				}
			}
		}
	}
	return res
}

//方法二: 分组+哈希表
func fourSumCount2(a, b, c, d []int) (ans int) {
	countAB := map[int]int{}
	for _, v := range a {
		for _, w := range b {
			countAB[v+w]++
		}
	}
	for _, v := range c {
		for _, w := range d {
			ans += countAB[-v-w]
		}
	}
	return
}

//func main()  {
//	A := []int{1, 1, -1, -1}
//	B := []int{-1, -1, 1, 1}
//	C := []int{1, -1, 0, -1}
//	D := []int{1, 1, -1, 1}
//	fmt.Println(fourSumCount(A, B, C, D))
//}
