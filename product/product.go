package product

type Category struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desciption"`
}

type Product struct {
	ID       int64  `json:"id"`
	Category int64  `json:"category"`
	Name     string `json:"name"`
}
