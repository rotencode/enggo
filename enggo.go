package main

import (
	"fmt"
	"io"
	"time"

	"bufio"
	"encoding/csv"
	"io/ioutil"
	"os"

	"regexp"
	"strings"

	"strconv"

	"./model"
	"./tools"
)

func test() {
	var str string

	var cache map[string]model.ExchangeData
	cache = make(map[string]model.ExchangeData)
	dat, _ := ioutil.ReadFile("./buffer.txt")

	str = string(dat)
	//	fmt.Println(str)
	//	var str string = "<tr><td  class=\"lm\">2015-05-20<td  class=\"rgt\">17.04<td  class=\"rgt\">17.49<td  class=\"rgt\">16.97<td  class=\"rgt\">17.15<td  class=\"rgt rm\">265,675,717<tr><td  class=\"lm\">2015-05-20<td  class=\"rgt\">17.04<td  class=\"rgt\">17.49<td  class=\"rgt\">16.97<td  class=\"rgt\">17.15<td  class=\"rgt rm\">265,675,717<tr></table>"
	//<tr>
	//<td  class="lm">2015-05-20
	//<td  class="rgt">17.04
	//<td  class="rgt">17.49
	//<td  class="rgt">16.97
	//<td  class="rgt">17.15
	//<td  class="rgt rm">265,675,717
	//<tr>
	//<td  class="lm">2015-05-19
	//<td  class="rgt">16.37
	//<td  class="rgt">17.09
	//<td  class="rgt">16.35
	//<td  class="rgt">17.04
	//<td  class="rgt rm">242,251,271
	//(-?\d*)

	begin := strings.Index(str, "<td  class=\"lm\">")
	end := strings.Index(str[begin:], "</table>")
	//	fmt.Println(begin, end)
	pure := str[begin : begin+end]
	//	fmt.Println("=======>")
	//	fmt.Println(pure)

	items := strings.Split(pure, "<tr>")
	//	fmt.Println("---->", len(items), "\n")
	for _, item := range items {
		//		fmt.Println("....", item, "---")
		//r, _ := regexp.Compile(".+<td  class=\"lm\">([d-]+)<td  class=\"rgt\">([d\\.]+)<td  class=\"rgt\">([d\\.]+)<td  class=\"rgt\">([d\\.]+)<td  class=\"rgt\">([d\\.]+)")
		//r, _ := regexp.Compile(".+<td  class=\"lm\">([\\d-]+)")
		r, _ := regexp.Compile("<td  class=\"lm\">([\\d-]+)[\\W]+<td  class=\"rgt\">([\\d.]+)[\\W]+<td  class=\"rgt\">([\\d.]+)[\\W]+<td  class=\"rgt\">([\\d.]+)[\\W]+<td  class=\"rgt\">([\\d.]+)[\\W]+<td  class=\"rgt rm\">([\\d,]+)")
		arr := r.FindStringSubmatch(item)
		fmt.Println(arr[1], arr[2], ";", arr[3], ";", arr[4], ";", arr[5], ";", arr[6])

		if len(arr) == 7 {
			// 日期/開市價/最高價/最低價/收市價/成交量
			var exchange model.ExchangeData
			exchange.ExchangeDate = arr[1]
			exchange.PriceFirst, _ = strconv.ParseFloat(arr[2], 32)
			exchange.PriceHigh, _ = strconv.ParseFloat(arr[3], 32)
			exchange.PriceLow, _ = strconv.ParseFloat(arr[4], 32)
			exchange.PriceLast, _ = strconv.ParseFloat(arr[5], 32)
			exchange.ExchangeAmount, _ = strconv.ParseInt(strings.Replace(arr[6], ",", "", -1), 10, 32)
			//
			fmt.Println("==exchange===>", exchange)
			cache[exchange.ExchangeDate] = exchange

		}
		//		for i, sub := range arr {
		//			fmt.Println(i, sub)
		//		}
		//		r.findstr
		//		fmt.Println(arr)
		//		fmt.Println(r.FindAllString(item, -1))

	}

}

func ttt() {
	// Load a TXT file.
	f, _ := os.Open("./file")

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}
		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.
		fmt.Println(record)
		fmt.Println(len(record))
		for value := range record {
			fmt.Printf(" %v, %v,\n", value, record[value])
		}
	}
}

//Contents: file.txt

//cat,dog,bird
//10,20,30,40
//fish,dog,snake
//}
func shenzhen() {
	file, _ := os.Open("./markets/shenzhen")
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		fmt.Println(fscanner.Text())
		var crawler tools.Crawler
		//		crawler.Stockid = ""
		crawler.Start(fscanner.Text(), tools.SHENZHEN)
	}

}
func shanghai() {

	file, _ := os.Open("./markets/shanghai")
	fscanner := bufio.NewScanner(file)
	for fscanner.Scan() {
		fmt.Println(fscanner.Text())
		var crawler tools.Crawler
		//		crawler.Stockid = ""
		crawler.Start(fscanner.Text(), tools.SHANGHAI)
	}
	//	 http://www.google.com.hk/finance/historical?q=SHA:600000&startdate=1990-01-02&enddate=2016-09-17&num=200&start=0

}
func main() {
	//	ttt()
	//	test()
	fmt.Println("enggo start", string(time.Now().Format("2006-01-02")))

	go shenzhen()
	//	shanghai()
	for {
		time.Sleep(100 * time.Second)
	}

}
