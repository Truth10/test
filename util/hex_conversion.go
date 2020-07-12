package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	RadixMap = make(map[int]string, 80)
)

func init() {
	for i := 0; i <= 9; i++ {
		RadixMap[i] = fmt.Sprintf("%d", i)
	}
	tempLowerWordIndex := 10
	for j := 'a'; j <= 'z'; j++ {
		RadixMap[tempLowerWordIndex] = fmt.Sprintf("%c", j)
		tempLowerWordIndex++
	}
	tempUpperWordIndex := 36
	for k := 'A'; k <= 'Z'; k++ {
		RadixMap[tempUpperWordIndex] = fmt.Sprintf("%c", k)
		tempUpperWordIndex++
	}
	RadixMap[62] = "!"
	RadixMap[63] = "@"
	RadixMap[64] = "#"
	RadixMap[65] = "$"
	RadixMap[66] = "%"
	RadixMap[67] = "^"
	RadixMap[68] = "&"
	RadixMap[69] = "?"
	RadixMap[70] = "("
	RadixMap[71] = ")"
	RadixMap[72] = "-"
	RadixMap[73] = "+"
	RadixMap[74] = "="
	RadixMap[75] = "~"
	RadixMap[76] = "{"
	RadixMap[77] = "}"
	RadixMap[78] = "["
	RadixMap[79] = "]"
}

func DecimalToAny(num, n int) string {
	new_num_str := ""
	var remainder int
	var remainder_string string
	for num != 0 {
		remainder = num % n
		if 80 > remainder && remainder > 9 {
			remainder_string = RadixMap[remainder]
		} else {
			remainder_string = strconv.Itoa(remainder)
		}
		new_num_str = remainder_string + new_num_str
		num = num / n
	}
	return new_num_str
}

func FindKey(in string) int {
	result := -1
	for k, v := range RadixMap {
		if in == v {
			result = k
		}
	}
	return result
}

// 任意进制转10进制
func AnyToDecimal(num string, n int) int {
	var new_num float64
	new_num = 0.0
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := float64(FindKey(value))
		if tmp != -1 {
			new_num = new_num + tmp*math.Pow(float64(n), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int(new_num)
}
