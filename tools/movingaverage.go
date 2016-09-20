package tools

import (
	"fmt"

	"../model"
)

func CaculMovingAvg(cache *[]model.ExchangeData) (err error) {
	for index, _ := range *cache {
		//		fmt.Println(index, val)
		if (*cache)[index].MA10 > 0 && (*cache)[index].MA30 > 0 && (*cache)[index].MA60 > 0 && (*cache)[index].MA365 > 0 {
			fmt.Println("已经计算过数据")
			continue
		}
		days := 10
		if index < days {
			var sum float64 = 0
			for _, sub_val := range (*cache)[0:index] {
				sum += sub_val.PriceLast
			}
			(*cache)[index].MA10 = sum / float64(index)
		} else {
			var sum float64 = 0
			for _, sub_val := range (*cache)[index-days : index] {
				sum += sub_val.PriceLast
			}
			(*cache)[index].MA10 = sum / float64(days)
		}
		days = 30
		if index < days {
			var sum float64 = 0
			for _, sub_val := range (*cache)[0:index] {
				sum += sub_val.PriceLast
			}
			(*cache)[index].MA30 = sum / float64(index)
		} else {
			var sum float64 = 0
			for _, sub_val := range (*cache)[index-days : index] {
				sum += sub_val.PriceLast
			}
			(*cache)[index].MA30 = sum / float64(days)
		}
		days = 60
		if index < days {
			var sum float64 = 0
			for _, sub_val := range (*cache)[0:index] {
				sum += sub_val.PriceLast
			}
			(*cache)[index].MA60 = sum / float64(index)
		} else {
			var sum float64 = 0
			for _, sub_val := range (*cache)[index-days : index] {
				sum += sub_val.PriceLast
			}
			(*cache)[index].MA60 = sum / float64(days)
		}
		days = 365
		if index < days {
			var sum float64 = 0
			for _, sub_val := range (*cache)[0:index] {
				sum += sub_val.PriceLast
			}
			(*cache)[index].MA365 = sum / float64(index)
		} else {
			var sum float64 = 0
			for _, sub_val := range (*cache)[index-days : index] {
				sum += sub_val.PriceLast
			}
			(*cache)[index].MA365 = sum / float64(days)
		}
	}
	//	fmt.Println("计算均值", cache)
	//	fmt.Println((*cache)[0])
	//	fmt.Println("最后一个")
	//	fmt.Println("最后一个", (*cache)[len(*cache)-1])
	return
}
