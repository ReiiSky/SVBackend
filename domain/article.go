package domain

type Article struct {
	ID       int    `json:"id" db:"id"`
	Title    string `json:"title" db:"title"`
	Content  string `json:"content" db:"content"`
	Category string `json:"category" db:"category"`
	Status   string `json:"status" db:"status"`
}
