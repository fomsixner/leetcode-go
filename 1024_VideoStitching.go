package leetcode_go

//视频拼接

func videoStitching(clips [][]int, T int) (ans int) {
	maxn := make([]int, T)
	last, pre := 0, 0
	for _, c := range clips {
		l, r := c[0], c[1]
		if l < T && r > maxn[l] {
			maxn[l] = r
		}
	}
	for i, v := range maxn {
		if v > last {
			last = v
		}
		if i == last {
			return -1
		}
		if i == pre {
			ans++
			pre = last
		}
	}
	return

}

//func main()  {
//	var clips [][]int
//	clips = [][]int{{16, 18}, {16, 20}, {3, 13}, {1, 18}, {0, 8}, {5, 6}, {13, 17}, {3, 17}, {5, 6}}
//	fmt.Println(videoStitching(clips, 15))
//}