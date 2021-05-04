package main

import (
	"fmt"
)

//PaddingNum 함수는 숫자가 한 자리일 때 앞에 0을 붙여주는 template 함수이다.
func PaddingNum(num int) string {
	n := fmt.Sprintf("%02d", num)
	return n
}
