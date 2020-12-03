package leetcode_go

//筛表法
func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	count := 0
	temp := make([]bool, n+1)
	for i := 2; i < n; i++ {
		if temp[i] == false {
			count++
			for j := i * 2; j < n; j += i {
				temp[j] = true
			}
		}
	}
	return count
}

//func isPrime(n int) bool {
//	if n == 2 {
//		return true
//	}
//	m := int(math.Sqrt(float64(n)))
//	for i:=2; i <= m; i++ {
//		if n % i == 0 {
//			return false
//		}
//	}
//	return true
//}
