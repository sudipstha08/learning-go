package controller

import (
	"bytes"
	"encoding/csv"

	"github.com/gin-gonic/gin"
)

func HandleCsvDownload(c *gin.Context) {
	var csvStruct [][]string
	csvStruct = [][]string{
		{"name", "address", "phone"},
		{"Mike", "United States", "1236524"},
		{"Micheal", "United kingdom", "8575675484"},
	}
	b := new(bytes.Buffer)
	w := csv.NewWriter(b)
	w.WriteAll(csvStruct)
	c.Writer.Write(b.Bytes())
}
