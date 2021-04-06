package leetcode_go

// * / + -
// func clumsy(N int) int {
// 	ans := 0
// 	var seq []int
// 	for i := 0; i < N; i++ {
// 		seq = append(seq, N-i)
// 	}
// 	for i := 0; i*4 < N; i++ {
// 		if i*4+2 < N {
// 			if i == 0 {
// 				ans += (N - i*4) * (N - i*4 - 1) / (N - i*4 - 2)
// 			} else {
// 				ans -= (N - i*4) * (N - i*4 - 1) / (N - i*4 - 2)
// 			}
// 		} else if i*4+1 < N {
// 			if i == 0 {
// 				ans += (N - i*4) * (N - i*4 - 1)
// 			} else {
// 				ans -= (N - i*4) * (N - i*4 - 1)
// 			}
// 		} else {
// 			if i == 0 {
// 				ans += (N - i*4)
// 			} else {
// 				ans -= (N - i*4)
// 			}
// 		}
// 		if i*4+3 < N {
// 			ans += (N - i*4 - 3)
// 		}
// 	}

// 	return ans
// }

// 使用栈模拟
func clumsy(N int) int {
	st := []int{N}
	index := 0
	N--
	for N > 0 {
		switch index % 4 {
		case 0:
			st[len(st)-1] *= N
		case 1:
			st[len(st)-1] /= N
		case 2:
			st = append(st, N)
		default:
			st = append(st, -N)
		}
		N--
		index++
	}
	ans := 0
	for _, v := range st {
		ans += v
	}
	return ans
}
