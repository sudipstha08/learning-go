package models

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}
