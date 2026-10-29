package application

import (
	"context"

	"github.com/ReiiSky/SharingVisionBackend/domain"
	"github.com/ReiiSky/SharingVisionBackend/infrastructure"
)

type getArticleService struct {
	kernel IKernel
}

func (gas getArticleService) GetArticles(ctx context.Context, offset, limit int) ([]domain.Article, error) {
	articles, err := gas.kernel.ArticleRepository().GetLatestArticles(ctx, infrastructure.NewQueryOptions(offset, limit))

	if err != nil {
		return []domain.Article{}, err
	}

	return articles, nil
}
