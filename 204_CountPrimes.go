package leetcode_go

//筛表法
//埃氏筛
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

//线性筛
//func countPrimes(n int) int {
//	primes := []int{}
//	isPrime := make([]bool, n)
//	for i := range isPrime {
//		isPrime[i] = true
//	}
//	for i := 2; i < n; i++ {
//		if isPrime[i] {
//			primes = append(primes, i)
//		}
//		for _, p := range primes {
//			if i*p >= n {
//				break
//			}
//			isPrime[i*p] = false
//			if i%p == 0 {
//				break
//			}
//		}
//	}
//	return len(primes)
//}
