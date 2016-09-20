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
		var crawler tools.Crawler
		crawler.Start(fscanner.Text(), tools.SHANGHAI)
	}
}
func loadRawData() {
	go shenzhen()
	go shanghai()
}

//清洗某一个数据项
func dataWash(stockid string) {
	fmt.Println(config.DataFilePath + stockid)
	var arr []model.ExchangeData
	tools.ParseCsvToArr(config.DataFilePath+stockid, &arr)
	tools.CaculMovingAvg(&arr)
	tools.CsvArrSave(&arr, stockid)
}

//数据清洗
func dataClean() {
	file, _ := os.Open(config.MarketsShenZhenFilePath)
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		go dataWash(fscanner.Text())
	}
	file, _ = os.Open(config.MarketsShangHaiFilePath)
	fscanner = bufio.NewScanner(file)
	for fscanner.Scan() {
		go dataWash(fscanner.Text())
	}
}

func main() {
	fmt.Println("enggo start", string(time.Now().Format("2006-01-02")))
	//	loadRawData()
	dataClean()
	for {
		time.Sleep(100 * time.Second)
	}

}
