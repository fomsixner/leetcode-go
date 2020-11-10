package leetcode_go

import "sort"

type A struct {
	Value int
	Ones int
}

type Ars []A

func (a Ars) Len() int {
	return len(a)
}

func (a Ars) Less(i, j int) bool {
	if a[i].Ones != a[j].Ones {
		return a[i].Ones < a[j].Ones
	}
	return a[i].Value < a[j].Value
}

func (a Ars) Swap(i, j int)  {
	a[i], a[j] = a[j], a[i]
}

func calone(v int) int {
	count := 0
	for v > 0 {
		if v%2 == 1 {
			count++
		}
		v /= 2
	}
	return count
}

func sortByBits(arr []int) []int {
	temp := Ars{}
	for _, v:= range arr {
		one := calone(v)
		temp = append(temp, A{v, one})
	}
	sort.Sort(temp)          //这里使用sort.Slice可以写的更短
	res := []int{}
	for _, v:= range temp {
		res = append(res, v.Value)
	}
	return res
}

//func sortByBits(arr []int) []int {
//	sort.Slice(arr, func(i, j int) bool {
//		x, y:= arr[i], arr[j]
//		ox, oy:= calone(x), calone(y)
//		return ox < oy || (ox == oy && x < y)
//	})
//	return arr
//}

//func main()  {
//	arr := []int{1024,512,256,128,64,32,16,8,4,2,1}
//	fmt.Println(sortByBits(arr))
//}