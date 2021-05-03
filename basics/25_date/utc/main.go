package main

import (
	"fmt"
	"time"
)

func main() {
	ktm, _ := time.LoadLocation("Asia/Kathmandu")
	utc, _ := time.LoadLocation("UTC")
	fmt.Println(utc)
	fmt.Println(time.Now().In(ktm).Format("2006-01-02 15:04:05"))
	// 2021-02-11 10:16:26
}
