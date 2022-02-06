package models

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// TableName will be used wherever we need to access the table name.
func (b *Article) TableName() string {
	return "todo"
}