package leetcode_go

import (
	"math"
)

//单词接龙
//转化为图, BFS
func canTrans(a ,b string) bool {
	count := 0
	for i:=0; i<len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	if count < 2 {
		return true
	}
	return false
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	in := false
	for _, word:= range wordList {
		if endWord == word {
			in = true
		}
	}
	if !in {
		return 0
	}
	n := len(wordList)
	//建图
	D := make([][]int, n+1)
	for i, _:= range D {
		D[i] = make([]int, n+1)
	}
	for i:=0; i<n; i++ {
		if canTrans(beginWord, wordList[i]) {
			D[0][i+1] = 1
		}
	}
	for i:=0; i<n; i++ {
		for j:=i+1; j < n; j++ {
			if canTrans(wordList[i], wordList[j]) {
				D[i+1][j+1], D[j+1][i+1] = 1, 1
			}
		}
	}
	//BFS
	queue := []int{0}
	dist := make([]int, n+1)
	for i, _:= range dist {
		dist[i] = math.MaxInt64       //表示不可达 or 未被访问
	}
	dist[0] = 0
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v!=0 && wordList[v-1] == endWord {
			return dist[v] + 1
		}
		for u:=1; u<n+1; u++ {
			if dist[u] == math.MaxInt64 && D[v][u]!=0 {
				dist[u] = dist[v] + 1
				queue = append(queue, u)
			}
		}
	}
	return 0
}

//func main()  {
//	wordlist := []string{"a","b","c"}
//	beginword , endword:= "a", "c"
//	fmt.Println(ladderLength(beginword, endword, wordlist))
//}