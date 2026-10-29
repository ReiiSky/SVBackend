package application

import (
	"context"
	"errors"

	"github.com/ReiiSky/SharingVisionBackend/application/dto"
	"github.com/ReiiSky/SharingVisionBackend/domain"
)

type updateArticleService struct {
	kernel IKernel
}

var (
	ArticleNotFoundError = errors.New("article not found")
)

func (uas updateArticleService) UpdateArticle(ctx context.Context, id int, articleInput dto.CreateArticleInput) error {
	article, err := uas.kernel.ArticleRepository().GetArticleByID(ctx, id)

	if err != nil {
		return err
	}

	if article == nil {
		return ArticleNotFoundError
	}

	return uas.kernel.ArticleRepository().UpdateArticle(ctx, id, domain.Article{
		Title:    articleInput.Title,
		Content:  articleInput.Content,
		Category: articleInput.Category,
		Status:   articleInput.Status,
	})
}
