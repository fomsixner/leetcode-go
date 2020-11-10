package leetcode_go

//方法一: 贪心
//可以通过映射来简化代码
func intToRoman(num int) string {
	var res string
	if num >= 1000 {
		for i:=0; i < num/1000; i++ {
			res += "M"
		}
		num %= 1000
	}
	if num >= 900 {
		res += "CM"
		num %= 900
	}
	if num >= 500 {
		res += "D"
		num %= 500
	}
	if num >= 400 {
		res += "CD"
		num %= 400
	}
	if num >= 100 {
		for i:=0; i < num/100; i++ {
			res += "C"
		}
		num %= 100
	}
	if num >= 90 {
		res += "XC"
		num %= 90
	}
	if num >= 50 {
		res += "L"
		num %= 50
	}
	if num >= 40 {
		res += "XL"
		num %= 40
	}
	if num >= 10 {
		for i:=0; i < num/10; i++ {
			res += "X"
		}
		num %= 10
	}
	if num >= 9 {
		res += "IX"
		num %= 9
	}
	if num >= 5 {
		res += "V"
		num %= 5
	}
	if num >= 4 {
		res += "IV"
		num %= 4
	}
	for i:=0; i < num; i++ {
		res += "I"
	}
	return res
}


//方法二: 硬编码
func intToRoman2(num int) string {
	thousands := []string{"", "M", "MM", "MMM"}
	hundreds := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	return thousands[num / 1000] + hundreds[num % 1000 / 100] + tens[num % 100 / 10] + ones[num % 10]
}
