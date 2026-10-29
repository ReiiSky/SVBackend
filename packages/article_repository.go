package packages

import (
	"context"

	"github.com/ReiiSky/SharingVisionBackend/domain"
	"github.com/ReiiSky/SharingVisionBackend/infrastructure"
)

type ArticleRepository interface {
	GetArticleByID(context.Context, int) (*domain.Article, error)
	GetLatestArticles(context.Context, infrastructure.QueryOptions) ([]domain.Article, error)
	CreateArticle(context.Context, domain.Article) error
}
