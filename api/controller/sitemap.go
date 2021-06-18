package controller

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type URL struct {
	Loc      string     `xml:"loc"`
	LastMod  *time.Time `xml:"lastmod,omitempty"`
	Priority float32    `xml:"priority,omitempty"`
}

// Sitemap represents a complete sitemap which can be marshaled to XML.
// New instances must be created with New() in order to set the xmlns
// attribute correctly. Minify can be set to make the output less human
// readable.
type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []*URL   `xml:"url"`

	Minify bool `xml:"-"`
}

// New returns a new Sitemap.
func NewSiteMap() *Sitemap {
	return &Sitemap{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  make([]*URL, 0),
	}
}

// Add adds an URL to a Sitemap.
func (s *Sitemap) Add(u *URL) {
	s.URLs = append(s.URLs, u)
}

const (
	Header     = `<?xml version="1.0" encoding="UTF-8"?>`
	StyleSheet = `<?xml-stylesheet type="text/css" href="http://localhost:5005/sitemap.css"?>`
)

func GetSiteMap(c *gin.Context) {
	sm := NewSiteMap()
	t := time.Unix(0, 0).UTC()
	sm.Add(&URL{
		Loc:     "http://example.com/",
		LastMod: &t,
	})

	marshaledData, _ := xml.MarshalIndent(sm, " ", "  ")

	data := StyleSheet + string(marshaledData)

	c.Data(http.StatusOK, "text/xml; charset=UTF-8", []byte(data))
}
