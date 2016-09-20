package tools

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"sort"

	"../config"
	"../model"
)

func CsvMapSave(cache *map[string]model.ExchangeData, stockid string) {
	ofile := string(config.DataFilePath + string(stockid))
	fout, err := os.Create(ofile)
	defer fout.Close()
	if err == nil {

		var keys []string
		for k := range *cache {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			str := fmt.Sprintf("%s,%.3f,%.3f,%.3f,%.3f,%d\n", (*cache)[k].ExchangeDate, (*cache)[k].PriceFirst,
				(*cache)[k].PriceHigh, (*cache)[k].PriceLow,
				(*cache)[k].PriceLast, (*cache)[k].ExchangeAmount)
			fmt.Println(str)
			fout.WriteString(str)
		}
		fmt.Sprintln("%v")
	} else {
		fmt.Print("错误", err)
	}
}

func CsvArrSave(arr *[]model.ExchangeData, stockid string) {
	ofile := string(config.DataFilePath + string(stockid))
	fout, err := os.Create(ofile)
	defer fout.Close()
	if err == nil {

		//		var keys []string
		//		for k := range *arr {
		//			keys = append(keys, k)
		//		}
		//		sort.Strings(keys)

		for _, val := range *arr {
			if val.MA10 > 0 && val.MA30 > 0 && val.MA60 > 0 && val.MA365 > 0 {
				str := fmt.Sprintf("%s,%.3f,%.3f,%.3f,%.3f,%d,%.3f,%.3f,%.3f,%.3f\n", val.ExchangeDate, val.PriceFirst,
					val.PriceHigh, val.PriceLow,
					val.PriceLast, val.ExchangeAmount,
					val.MA10, val.MA30, val.MA60, val.MA365)
				fmt.Println(str)
				fout.WriteString(str)
			} else {
				str := fmt.Sprintf("%s,%.3f,%.3f,%.3f,%.3f,%d\n", val.ExchangeDate, val.PriceFirst,
					val.PriceHigh, val.PriceLow,
					val.PriceLast, val.ExchangeAmount)
				fmt.Println(str)
				fout.WriteString(str)
			}
		}
		fmt.Sprintln("%v")
	} else {
		fmt.Print("错误", err)
	}
}

func ParseCSVToMap(filename string, items *map[string]model.ExchangeData) {
	csvfile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		//		os.Exit(1)
		return
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sanity check, display to standard output
	//	for _, each := range rawCSVdata {
	//		//		fmt.Printf("email : %s and timestamp : %s\n", each[0], each[1])
	//		fmt.Println(each)
	//	}

	// now, safe to move raw CSV data to struct

	var item model.ExchangeData

	for _, each := range rawCSVdata {
		//		*日期 / 開市價 / 最高價 / 最低價 / 收市價 / 成交量
		item.ExchangeDate = each[0]
		item.PriceFirst, _ = strconv.ParseFloat(each[1], 32)
		item.PriceHigh, _ = strconv.ParseFloat(each[2], 32)
		item.PriceLow, _ = strconv.ParseFloat(each[3], 32)
		item.PriceLast, _ = strconv.ParseFloat(each[4], 32)
		item.ExchangeAmount, _ = strconv.ParseInt(each[5], 10, 32)
		//		allRecords = append(allRecords, oneRecord)
		(*items)[item.ExchangeDate] = item
	}

	// second sanity check, dump out allRecords and see if
	// individual record can be accessible
	//	fmt.Println(allRecords)
	//	fmt.Println(allRecords[2].Email)
	//	fmt.Println(allRecords[2].Date)
}
func ParseCsvToArr(filename string, arr *[]model.ExchangeData) {
	csvfile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		//		os.Exit(1)
		return
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//	// sanity check, display to standard output
	//	for _, each := range rawCSVdata {
	//		//		fmt.Printf("email : %s and timestamp : %s\n", each[0], each[1])
	//		fmt.Println(each)
	//	}

	// now, safe to move raw CSV data to struct

	var item model.ExchangeData

	for _, each := range rawCSVdata {
		//		*日期 / 開市價 / 最高價 / 最低價 / 收市價 / 成交量
		item.ExchangeDate = each[0]
		item.PriceFirst, _ = strconv.ParseFloat(each[1], 32)
		item.PriceHigh, _ = strconv.ParseFloat(each[2], 32)
		item.PriceLow, _ = strconv.ParseFloat(each[3], 32)
		item.PriceLast, _ = strconv.ParseFloat(each[4], 32)
		item.ExchangeAmount, _ = strconv.ParseInt(each[5], 10, 32)
		//		allRecords = append(allRecords, oneRecord)
		(*arr) = append((*arr), item)
	}

}

func SerilizeMapToCSV(filename string, items *map[string]model.ExchangeData) {
	f, err := os.Create("haha2.csv")

	if err != nil {

		panic(err)

	}
	defer f.Close()

	w := csv.NewWriter(f)

	for _, k := range *items {
		w.Write([]string{k.ExchangeDate, strconv.FormatFloat(k.PriceFirst, 'g', 1, 64), strconv.FormatFloat(k.PriceHigh, 'g', 1, 64), strconv.FormatFloat(k.PriceLow, 'g', 1, 64), strconv.FormatFloat(k.PriceLast, 'g', 1, 64), strconv.Itoa(int(k.ExchangeAmount))})
		//		w.Write([]string{"1", "张三", "23"})
		//		w.Write([]string{"2", "李四", "24"})
		//		w.Write([]string{"3", "王五", "25"})
		//		w.Write([]string{"4", "赵六", "26"})
	}

	w.Flush()
}
