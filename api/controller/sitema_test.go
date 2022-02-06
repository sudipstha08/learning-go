package controller

// import (
// 	"bytes"
// 	"encoding/xml"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/snabb/sitemap"
// )

// type Servers struct {
// 	XMLName xml.Name `xml:"urlset"`
// 	Version string   `xml:"version,attr"`
// 	XmlNs   string   `xml:"xmlns,attr"`
// 	Svs     []URL    `xml:"url"`
// }

// type URL struct {
// 	Loc     string `xml:"loc"`
// 	LastMod string `xml:"lastmod"`
// }

// const (
// 	// A generic XML header suitable for use with the output of Marshal.
// 	// This is not automatically added to any output of this package,
// 	// it is provided as a convenience.
// 	Header     = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
// 	StyleSheet = `<?xml-stylesheet type="text/css" href="xml-styles.css"?>` + "\n"
// )

// func GetSiteMap(c *gin.Context) {
// 	c.Header("Content-Type", "application/xml")

// 	v := &Servers{Version: "1", XmlNs: "http://www.sitemaps.org/schemas/sitemap/0.9"}
// 	v.Svs = append(v.Svs, URL{"http://google.com", "2021Jan13"})
// 	v.Svs = append(v.Svs, URL{"http://youtube.com", "2021J21"})
// 	// output, err := xml.MarshalIndent(v, "  ", "    ")
// 	bytexml, err := xml.MarshalIndent(v, " ", "  ")
// 	bytexml = []byte(StyleSheet + string(bytexml))
// 	bytexml = []byte(xml.Header + string(bytexml))
// 	// bytexml = []byte(xml. + string(bytexml))
// 	if err != nil {
// 		fmt.Printf("error: %v\n", err)
// 	}
// 	// os.Stdout.Write([]byte(xml.Header))
// 	fmt.Println("ddddd")
// 	os.Stdout.Write(bytexml)
// 	sm := sitemap.New()
// 	t := time.Unix(0, 0).UTC()
// 	sm.Add(&sitemap.URL{
// 		Loc:        "http://example.com/",
// 		LastMod:    &t,
// 		ChangeFreq: sitemap.Daily,
// 	})
// 	c.XML(http.StatusOK, bytes.NewBuffer(bytexml))
// }

// // https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/07.1.html
// https://github.com/beevik/etree
// https://www.geeksforgeeks.org/displaying-xml-using-css/