package leetcode_go

func maxProfit(prices []int) int {
	n := len(prices)
	if n==1 {return 0}
	profit, buy := 0, prices[0]
	for i:=1; i<n; i++ {
		if prices[i] > buy {
			profit += prices[i] - buy
		}
		buy = prices[i]
	}
	return profit
}
