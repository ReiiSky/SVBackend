package application

import (
	"context"

	"github.com/ReiiSky/SharingVisionBackend/domain"
)

type getDetailArticleService struct {
	kernel IKernel
}

func (gdas getDetailArticleService) GetArticleByID(ctx context.Context, id int) (*domain.Article, error) {
	article, err := gdas.kernel.ArticleRepository().GetArticleByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return article, nil
}
