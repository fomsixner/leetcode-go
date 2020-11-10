package leetcode_go

//单词拆分
//给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

//动态规划

func wordBreak(s string, wordDict []string) bool {
	var match func(s string, wordDict []string) bool
	match = func(s string, wordDict []string) bool {
		for _, word := range wordDict {
			if s == word {
				return true
			}
		}
		return false
	}
	n := len(s)
	dp := make([]bool, n+1)
	//初始化, dp[0]表示空串且合法
	dp[0] = true
	for i:=1; i <= n; i+=1 {
		for j:=0; j < i; j+=1 {
			if dp[j] && match(s[j:i], wordDict) {
				dp[i] = true
				break
			}
		}
	}
	if dp[n] {
		return true
	}
	return false
}