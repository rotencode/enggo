package tools

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

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
	//	var urls = []string{"http://www.google.com.hk/finance/historical?q=SHA:%s&startdate=1990-01-02&enddate=%s&num=200&start=%d","http://www.google.com.hk/finance/historical?q=SHE:%s&startdate=1990-01-02&enddate=%s&num=200&start=%d"}
	//	url_shanghai
	//	#define SHANGHAI_URL "http://www.google.com.hk/finance/historical?q=SHA:%s&startdate=1990-01-02&enddate=%s&num=200&start=%d"
	//#define SHENZHEN_URL "http://www.google.com.hk/finance/historical?q=SHE:%s&startdate=1990-01-02&enddate=%s&num=200&start=%d"
}

func (crawler *Crawler) getUrl() (url string, err error) {
	if crawler.Market == SHANGHAI {
		url = fmt.Sprintf(crawler.url_shanghai, crawler.Stockid, string(time.Now().Format("2006-01-02")), crawler.position)
	} else {
		url = fmt.Sprintf(crawler.url_shenzhen, crawler.Stockid, string(time.Now().Format("2006-01-02")), crawler.position)
	}
	return
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

}
func (crawler *Crawler) CrawlerRegexMatch(str string) {
	//	crawler.cache[]

}

func (crawler *Crawler) CrawlerTask() {
	str, err := crawler.CrawlerRequest()
	if err == nil {

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

		re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
		re.FindAllStringSubmatch(str, 10)
	}
}

func (crawler *Crawler) Start() (err error) {
	crawler.position = 0
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
