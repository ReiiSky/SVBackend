package application

import (
	"context"
)

type deleteArticleService struct {
	kernel IKernel
}

func (uas deleteArticleService) DeleteArticle(ctx context.Context, id int) error {
	article, err := uas.kernel.ArticleRepository().GetArticleByID(ctx, id)

	if err != nil {
		return err
	}

	if article == nil {
		return ArticleNotFoundError
	}

	return uas.kernel.ArticleRepository().DeleteArticle(ctx, id)
}
