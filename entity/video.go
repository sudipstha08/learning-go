package entity

type Person struct {
	FirstName string `json: "firstname"`
	LastName string `json: "lastname"`
	Age int8 `json: "age"`
	Email string `json: "email"`
}
type Video struct {
	Title       string `json: "title"`
	Description string `json: "description"`
	URL         string `json: "url"`
	Author string `json: "author"`
}


