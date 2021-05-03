package main

import (
	"fmt"
	"time"
)

// Magical Reference date -> no two fields same in the above date
// Mon Jan 2 15:04:05 MST 2006
func main() {
	input := "2017-08-31"
	layout := "2006-01-02"
	t, _ := time.Parse(layout, input)
	fmt.Println(t)
	fmt.Println(t.Format("02-Jan-2006"))
	
}
