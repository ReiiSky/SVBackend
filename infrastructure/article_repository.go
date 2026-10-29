package infrastructure

import (
	"context"

	"github.com/ReiiSky/SharingVisionBackend/domain"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type articleRepository struct {
	connection DatabaseConnector
}

func NewArticleRepository(connection DatabaseConnector) articleRepository {
	return articleRepository{connection}
}

func (ar articleRepository) CreateArticle(ctx context.Context, article domain.Article) error {
	return ar.connection.RunCommand(ctx, func(tx *sqlx.Tx) error {
		dml, args := sqlbuilder.NewInsertBuilder().
			InsertInto("posts").
			Cols("title", "content", "category", "status").
			Values(article.Title, article.Content, article.Category, article.Status).
			BuildWithFlavor(sqlbuilder.MySQL)

		_, err := tx.ExecContext(ctx, dml, args...)

		if err != nil {
			return err
		}

		return nil
	})
}
