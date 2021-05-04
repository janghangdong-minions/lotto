package main

import "strconv"

//PaddingNum 함수는 숫자가 한 자리일 때 앞에 0을 붙여주는 template 함수이다.
func PaddingNum(num int) string {
	n := strconv.Itoa(num)
	for len(n) < 2 {
		n = "0" + n
	}
	return n
}
