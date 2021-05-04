package main

import "strconv"

func PaddingNum(num int) string {
	n := strconv.Itoa(num)
	for len(n) < 2 {
		n = "0" + n
	}
	return n
}
