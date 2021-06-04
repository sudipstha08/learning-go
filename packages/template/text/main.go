package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"text/template"
)

type Inventory struct {
	Material string
	Count    uint
}

func main() {
	sweaters := Inventory{"wool", 17}
	data, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		fmt.Println("errr", err)
	}
	tmpl, err := template.New("test-template").Parse(string(data))
	if err != nil {
		fmt.Println("Error while parsing", err)
	}
	buf := new(bytes.Buffer)

	err = tmpl.Execute(buf, sweaters)
	body := buf.String()
	fmt.Println("body", body)
	if err != nil {
		fmt.Println("Error while executing", err)
	}
}
