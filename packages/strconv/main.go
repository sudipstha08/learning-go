// https://golang.org/pkg/strconv/#pkg-examples

package main

import (
	"fmt"
	"strconv"
)

func main() {
	//string to int
	i, err := strconv.Atoi("-42")
	if err == nil {
		fmt.Println(i)
	}
	//int to string
	s := strconv.Itoa(-42)
	fmt.Println(s)

	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	s3 := strconv.FormatInt(-42, 16)
	s4 := strconv.FormatUint(42, 16)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	q := strconv.Quote("Hello, 世界")
	z := strconv.QuoteToASCII("Hello, 世界")
	fmt.Println(q)
	fmt.Println(z)

}
