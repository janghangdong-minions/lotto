package main

import (
	"math/rand"
	"time"
)

// GenRandomNums 랜덤 값 출력
func GenRandomNums(min, max, ranges int) LottoNum {
	// 로또 1~45까지 숫자 중 7개의 숫자를 뽑는다.
	var nums []int
	for i := 0; i < ranges; i++ {
		rand.Seed(time.Now().UnixNano())
		nums = append(nums, rand.Intn(max-min+1)+min)
	}
	l := LottoNum{nums}
	return l
}
