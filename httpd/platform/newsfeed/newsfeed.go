package newsfeed

type Getter interface {
	// GetAll will return the slice of items
	GetAll() []Item
}

type Adder interface {
	// Adder will take one otem
	Add(item Item)
}
type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Repo struct {
	Items []Item
}

// constructor
func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

// returns the slice of Item
func (r *Repo) GetAll() []Item {
	return r.Items
}
