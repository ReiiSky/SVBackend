package dto

type CreateArticleInput struct {
	Title    string `validate:"required,min=20,max=60" json:"title"`
	Content  string `validate:"required,min=200,max=800" json:"content"`
	Category string `validate:"required,min=3,max=30" json:"category"`
	Status   string `validate:"required,oneof=publish draft trash" json:"status"`
}
