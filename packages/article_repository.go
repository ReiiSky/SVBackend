package packages

import (
	"context"

	"github.com/ReiiSky/SharingVisionBackend/domain"
)

type ArticleRepository interface {
	// GetLatestArticles(*infrastructure.QueryOptions) ([]domain.Article, error)
	CreateArticle(context.Context, domain.Article) error
}
