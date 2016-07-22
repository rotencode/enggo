package main

import (
	"fmt"

	"./tools"
)

func main() {
	fmt.Println("enggo start")
	var crawler tools.Crawler
	crawler.Stockid = ""
	crawler.Start()
}
