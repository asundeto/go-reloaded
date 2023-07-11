package reloaded

import (
	"fmt"
	"strconv"
)

func hexAndBinToDec(s string, n int) string {
	if len(s) == 0 {
		return ""
	}
	numByte := 2
	if n == 1 {
		numByte = 16
	}
	num, err := strconv.ParseInt(s, numByte, 64)
	if err != nil {
		fmt.Println("Error! Incorrect format!")
		return s
	}
	return Int64ToString(num)
}

func toLowerCase(s rune) rune {
	res := ' '
	if s >= 'A' && s <= 'Z' {
		res = s + 32
	} else {
		res = s
	}
	return res
}

func toUpperCase(s rune) rune {
	res := ' '
	if s >= 'a' && s <= 'z' {
		res = s - 32
	} else {
		res = s
	}
	return res
}

func myReverse(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}

func Int64ToString(x int64) string {
	res := ""
	if x == 0 {
		return "0"
	}
	for x > 0 {
		res = string(rune(x%10+'0')) + res
		x = x / 10
	}
	return res
}

func Atoi(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res *= 10
		res += int(s[i] - 48)
	}
	return res
}

