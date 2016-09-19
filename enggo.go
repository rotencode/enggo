package main

import (
	"fmt"

	"time"

	"bufio"
	"os"

	"./config"
	"./model"
	"./tools"
)

func shenzhen() {
	file, _ := os.Open(config.MarketsShenZhenFilePath)
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		fmt.Println(fscanner.Text())
		var crawler tools.Crawler
		crawler.Start(fscanner.Text(), tools.SHENZHEN)
		//		time.Sleep(30 * time.Second)
	}
}
func shanghai() {

	file, _ := os.Open(config.MarketsShangHaiFilePath)
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		fmt.Println(fscanner.Text())
		fmt.Println(fscanner.Text())
		//		var crawler tools.Crawler
		var arr []model.ExchangeData
		tools.ParseCsvToArr(config.DataFilePath+fscanner.Text(), &arr)
		tools.CaculMovingAvg(&arr, 10)
	}

	//CaculMovingAvg
	//	 http://www.google.com.hk/finance/historical?q=SHA:600000&startdate=1990-01-02&enddate=2016-09-17&num=200&start=0
}
func loadRawData() {
	//	go shenzhen()
	go shanghai()
}
func main() {
	fmt.Println("enggo start", string(time.Now().Format("2006-01-02")))
	loadRawData()
	for {
		time.Sleep(100 * time.Second)
	}

}
