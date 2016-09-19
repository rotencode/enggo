package algorithm

import (
	"fmt"

	"../model"
)

//http://ju.outofmemory.cn/entry/103439
type AlgorithmBaseInterface interface {
	IsMeetBuy(date string, data *map[string]model.ExchangeData) (meetable bool)
	IsMeetSell(date string, data *map[string]model.ExchangeData) (meetable bool)
}

func test() {
	fmt.Println("...")
}
