package tools

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"strconv"

	"../model"
)

const (
	SHENZHEN = iota
	SHANGHAI
)

type Crawler struct {
	Stockid      string
	url          string
	baseurl      string
	position     int
	Market       int
	cache        map[string]model.ExchangeData
	url_shanghai string //= "http://www.google.com.hk/finance/historical?q=SHA:%s&startdate=1990-01-02&enddate=%s&num=200&start=%d"
	url_shenzhen string //=  "http://www.google.com.hk/finance/historical?q=SHE:%s&startdate=1990-01-02&enddate=%s&num=200&start=%d"

}

func (crawler *Crawler) CrawlerRegexMatch(str string) (hasNext bool) {
	hasNext = false
	begin := strings.Index(str, "<td  class=\"lm\">")
	end := strings.Index(str[begin:], "</table>")
	pure := str[begin : begin+end]

	items := strings.Split(pure, "<tr>")
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
			exchange.ExchageDate = arr[1]
			exchange.PriceFirst, _ = strconv.ParseFloat(arr[2], 32)
			exchange.PriceHigh, _ = strconv.ParseFloat(arr[3], 32)
			exchange.PriceLow, _ = strconv.ParseFloat(arr[4], 32)
			exchange.PriceLast, _ = strconv.ParseFloat(arr[5], 32)
			exchange.ExchangeAmount, _ = strconv.ParseInt(strings.Replace(arr[6], ",", "", -1), 10, 32)
			//
			fmt.Println("==exchange===>", exchange)
			crawler.cache[exchange.ExchageDate] = exchange
		}
	}
	if len(items) == 200 {
		hasNext = true
	}
	return
}

func (crawler *Crawler) getUrl() (url string, err error) {
	if crawler.Market == SHANGHAI {
		url = fmt.Sprintf(crawler.url_shanghai, crawler.Stockid, string(time.Now().Format("2006-01-02")), crawler.position)
	} else {
		url = fmt.Sprintf(crawler.url_shenzhen, crawler.Stockid, string(time.Now().Format("2006-01-02")), crawler.position)
	}
	return
}
func (crawler *Crawler) crawlerSave() {
	ofile := string("./data/" + string(crawler.Market))
	fout, err := os.Create(ofile)
	defer fout.Close()
	if err == nil {
		for k, v := range crawler.cache {
			fmt.Printf("k=%v, v=%v\n", k, v)
			fout.WriteString(fmt.Sprintln("%v", v))

		}
		fmt.Sprintln("%v")
	}
}
func (crawler *Crawler) CrawlerRequest() (str string, err error) {
	c_url, _ := crawler.getUrl()

	response, _ := http.Get(c_url)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	str = string(body)
	fmt.Println("task begin")

	fmt.Println(body)
	return
}
func (crawler *Crawler) CrawlerInitCache() {
	dat, _ := ioutil.ReadFile("./buffer.txt")

	str := string(dat)
	fmt.Println(str)
}
func (crawler *Crawler) CrawlerTask() {
	for {
		str, err := crawler.CrawlerRequest()
		if err == nil {
			hasNext := crawler.CrawlerRegexMatch(str)
			if !hasNext {
				crawler.crawlerSave()
			}
		}
	}

}

func (crawler *Crawler) Start() (err error) {
	crawler.position = 0
	crawler.cache = make(map[string]model.ExchangeData)
	crawler.url_shanghai = "http://www.google.com.hk/finance/historical?q=SHA:%s&startdate=1990-01-02&enddate=%s&num=200&start=%d"
	crawler.url_shenzhen = "http://www.google.com.hk/finance/historical?q=SHE:%s&startdate=1990-01-02&enddate=%s&num=200&start=%d"
	if len(crawler.Stockid) == 0 {
		err = errors.New("stock id should not be none")
		fmt.Println("start error")
		return
	}
	go crawler.CrawlerTask()
	return
}
