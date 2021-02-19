package models

import "time"

type ArticleModel struct {
	Title   string    `json:"title"`
	Content string    `json:"content"`
}
