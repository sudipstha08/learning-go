package newsfeed

type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Repo struct {
	Items []Item
}

// constructor
func New() *Repo {
	return &Repo{}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

// returns the slice of Item
func (r *Repo) GetAll() []Item {
	return r.Items
}
