package tools

import (
	"errors"
	"fmt"
)

type Crawler struct {
	Stockid string
	url     string
	baseurl string
}

func (crawler *Crawler) ScrawlerTask() {
	fmt.Println("task begin")
}

func (crawler *Crawler) Start() (err error) {
	if len(crawler.Stockid) == 0 {
		err = errors.New("stock id should not be none")
		fmt.Println("start error")
		return
	}
	go crawler.ScrawlerTask()
	return
}
