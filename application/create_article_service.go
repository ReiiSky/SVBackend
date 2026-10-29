package application

import (
	"context"

	"github.com/ReiiSky/SharingVisionBackend/application/dto"
	"github.com/ReiiSky/SharingVisionBackend/domain"
)

type createArticleService struct {
	kernel IKernel
}

func (cas createArticleService) CreateArticle(ctx context.Context, articleInput dto.CreateArticleInput) error {
	return cas.kernel.ArticleRepository().CreateArticle(ctx, domain.Article{
		Title:    articleInput.Title,
		Content:  articleInput.Content,
		Category: articleInput.Category,
		Status:   articleInput.Status,
	})
}
