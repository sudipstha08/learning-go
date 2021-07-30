package main

import (
	"fmt"
	"strings"
)

func main() {
	file := "https://storage.googleapis.com/warehouse-3374e.appspot.com/images/1627031765921266092.png"
	aa := strings.Split(file, "warehouse-3374e.appspot.com/")
	fmt.Println(aa[1])
}
