package tools

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"../model"
)

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
	for _, each := range rawCSVdata {
		//		fmt.Printf("email : %s and timestamp : %s\n", each[0], each[1])
		fmt.Println(each)
	}

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
