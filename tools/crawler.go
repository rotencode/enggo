package tools

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
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

func (crawler *Crawler) ScrawlerTask() {
	c_url, _ := crawler.getUrl()

	response, _ := http.Get(c_url)
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("task begin")

	fmt.Println(body)
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
	go crawler.ScrawlerTask()
	return
}
