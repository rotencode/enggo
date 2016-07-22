package main

import (
	"fmt"
	"time"

	"./tools"
)

func main() {
	fmt.Println("enggo start", string(time.Now().Format("2006-01-02")))
	var crawler tools.Crawler
	crawler.Stockid = ""
	crawler.Start()
}
