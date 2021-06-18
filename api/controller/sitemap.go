package controller

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Servers struct {
	XMLName xml.Name `xml:"servers"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server"`
}

type server struct {
	ServerName string `xml:"serverName"`
	ServerIP   string `xml:"serverIP"`
}

func GetSiteMap(c *gin.Context) {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{"Shanghai_VPN", "127.0.0.1"})
	v.Svs = append(v.Svs, server{"Beijing_VPN", "127.0.0.2"})
	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	// os.Stdout.Write([]byte(xml.Header))
fmt.Println("ddddd")
	os.Stdout.Write(output)
	c.JSON(http.StatusOK, output)
}

// https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/07.1.html