package tools

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"../config"

	"../model"
)

const (
	SHENZHEN = iota
	SHANGHAI
)

type Crawler struct {
	Stockid string
	url     string
	baseurl string
	//本地最新的成交记录日期。更新数据从当前日期开始。
	start_date   string
	position     int
	Market       int
	cache        map[string]model.ExchangeData
	url_shanghai string //= "http://www.google.com.hk/finance/historical?q=SHA:%s&startdate=%s&enddate=%s&num=200&start=%d"
	url_shenzhen string //=  "http://www.google.com.hk/finance/historical?q=SHE:%s&startdate=%s&enddate=%s&num=200&start=%d"

}

func (crawler *Crawler) CrawlerRegexMatch(str string) (hasNext bool) {

	hasNext = false
	begin := strings.Index(str, "class=\"lm\">")
	if begin == -1 {
		return false
	}
	end := strings.Index(str[begin:], "</table>")
	if end == -1 {
		return false
	}
	pure := str[begin : begin+end]

	items := strings.Split(pure, "<tr>")
	for _, item := range items {
		r, _ := regexp.Compile("class=\"lm\">([\\d-]+)[\\W]+<td class=\"rgt\">([\\d.]+)[\\W]+<td class=\"rgt\">([\\d.]+)[\\W]+<td class=\"rgt\">([\\d.]+)[\\W]+<td class=\"rgt\">([\\d.]+)[\\W]+<td class=\"rgt rm\">([\\d*,]+)")
		arr := r.FindStringSubmatch(item)
		fmt.Println("==========>")
		fmt.Println(arr)
		fmt.Println("《==========")
		if len(arr) == 7 {

			fmt.Println(arr[1], arr[2], ";", arr[3], ";", arr[4], ";", arr[5], ";", arr[6])

			var exchange model.ExchangeData
			exchange.ExchangeDate = arr[1]
			exchange.PriceFirst, _ = strconv.ParseFloat(arr[2], 32)
			exchange.PriceHigh, _ = strconv.ParseFloat(arr[3], 32)
			exchange.PriceLow, _ = strconv.ParseFloat(arr[4], 32)
			exchange.PriceLast, _ = strconv.ParseFloat(arr[5], 32)
			exchange.ExchangeAmount, _ = strconv.ParseInt(strings.Replace(arr[6], ",", "", -1), 10, 32)
			//
			fmt.Println("==exchange===>", exchange)
			crawler.cache[exchange.ExchangeDate] = exchange
		} else {
			fmt.Println("....", item, "---")
			fmt.Println("错误数据")
		}

	}
	fmt.Print("内容的长度%d\n\n", len(items))
	if len(items) == 200 {
		hasNext = true
	}
	return
}

func (crawler *Crawler) getUrl() (url string, err error) {
	if crawler.Market == SHANGHAI {
		url = fmt.Sprintf(crawler.url_shanghai, crawler.Stockid, crawler.start_date, string(time.Now().Format("2006-01-02")), crawler.position)
	} else {
		url = fmt.Sprintf(crawler.url_shenzhen, crawler.Stockid, crawler.start_date, string(time.Now().Format("2006-01-02")), crawler.position)
	}
	return
}

func (crawler *Crawler) crawlerSave() {
	fmt.Println("crawlerSave", crawler.cache)
	ofile := string(config.DataFilePath + string(crawler.Stockid))
	fout, err := os.Create(ofile)
	defer fout.Close()
	if err == nil {

		var keys []string
		for k := range crawler.cache {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			str := fmt.Sprintf("%s,%.3f,%.3f,%.3f,%.3f,%d\n", crawler.cache[k].ExchangeDate, crawler.cache[k].PriceFirst, crawler.cache[k].PriceHigh, crawler.cache[k].PriceLow, crawler.cache[k].PriceLast, crawler.cache[k].ExchangeAmount)
			fmt.Println(str)
			fout.WriteString(str)
		}
		fmt.Sprintln("%v")
	} else {
		fmt.Print("错误", err)
	}
}

func (crawler *Crawler) CrawlerRequest() (str string, err error) {
	c_url, _ := crawler.getUrl()
	//	fmt.Println("\n%s\n", c_url)
	timeout := time.Duration(120 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	//	client.Get(url)
	response, h_err := client.Get(c_url)
	fmt.Println("==============%+v", h_err)
	//	defer response.Body.Close()
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	str = string(body)

	crawler.position += 200
	return
}
func (crawler *Crawler) CrawlerTask() {
	for {
		str, err := crawler.CrawlerRequest()
		if err == nil {
			hasNext := crawler.CrawlerRegexMatch(str)
			if !hasNext {
				fmt.Println("保存到数据文件中")
				crawler.crawlerSave()
				return
			}
		}
	}

}

//func ParseCSVToMap(filename string, items *map[string]model.ExchangeData) {

func (crawler *Crawler) initCache() (err error) {
	ParseCSVToMap(config.DataFilePath+crawler.Stockid, &crawler.cache)
	if len(crawler.cache) == 0 {
		crawler.start_date = "1990-01-01"
	} else {
		var keys []string
		for k := range crawler.cache {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		crawler.start_date = keys[len(keys)-1]
		fmt.Println(keys[len(keys)-1])
	}
	return
}

func (crawler *Crawler) Start(stockid string, market int) (err error) {
	if len(stockid) == 0 {
		return
	}
	crawler.Market = market
	crawler.Stockid = stockid
	fmt.Printf("%+v, %+v", crawler.Market, crawler.Stockid)
	crawler.position = 0
	crawler.cache = make(map[string]model.ExchangeData)
	crawler.url_shanghai = "http://www.google.com.hk/finance/historical?q=SHA:%s&startdate=%s&enddate=%s&num=200&start=%d"
	crawler.url_shenzhen = "http://www.google.com.hk/finance/historical?q=SHE:%s&startdate=%s&enddate=%s&num=200&start=%d"
	crawler.initCache()

	if string(time.Now().Format("2006-01-02")) == crawler.start_date {
		fmt.Println("已经缓存了最新数据")
		return
	}
	if len(crawler.Stockid) == 0 {
		err = errors.New("stock id should not be none")
		fmt.Println("start error")
		time.Sleep(5 * time.Second)
		return
	}
	crawler.CrawlerTask()
	time.Sleep(5 * time.Second)

	return
}
