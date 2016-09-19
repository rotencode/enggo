package tools

import (
	"fmt"

	"../model"
)

func CaculMovingAvg(cache *[]model.ExchangeData, days int) (err error) {
	fmt.Println("计算均值")
	fmt.Println(*cache)
	return
}
