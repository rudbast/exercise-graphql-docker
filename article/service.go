package articleservice

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rudbast/exercise-graphql-docker/jsonapi"
	"github.com/rudbast/exercise-graphql-docker/model"
)

type (
	ArticleService struct {
		database *sql.DB
	}
)

func New(db *sql.DB) ArticleService {
	return ArticleService{
		database: db,
	}
}

func (s ArticleService) GetArticleList(ctx context.Context, params map[string]interface{}) ([]jsonapi.ResponseObject, error) {
	var articles []model.Article
	var err error

	// NOTE: Will need to create a query builder if parameter count gets larger.
	if title, ok := params["title"].(string); ok {
		articles, err = model.GetArticleListByTitle(ctx, s.database, title)
	} else {
		articles, err = model.GetArticleList(ctx, s.database)
	}

	if err != nil {
		return nil, err
	}

	var resp []jsonapi.ResponseObject

	for _, article := range articles {
		obj := jsonapi.ResponseObject{
			Type:      "article",
			ID:        fmt.Sprintf("%d", article.ID),
			Attribute: article,
		}

		resp = append(resp, obj)
	}

	return resp, nil
}

func (s ArticleService) GetArticle(ctx context.Context, id int64) (*jsonapi.ResponseObject, error) {
	article, err := model.GetArticle(ctx, s.database, id)
	if err != nil {
		return nil, err
	}

	return &jsonapi.ResponseObject{
		Type:      "article",
		ID:        fmt.Sprintf("%d", article.ID),
		Attribute: article,
	}, nil
}

func (s ArticleService) AddArticle(ctx context.Context, title, content, thumbnail string) error {
	articleModel := model.Article{
		Title:     title,
		Content:   content,
		Thumbnail: thumbnail,
	}

	_, err := model.AddArticle(ctx, s.database, articleModel)
	return err
}

func (s ArticleService) UpdateArticle(ctx context.Context, id int64, title, content, thumbnail string) error {
	articleModel := model.Article{
		ID:        id,
		Title:     title,
		Content:   content,
		Thumbnail: thumbnail,
	}

	_, err := model.UpdateArticle(ctx, s.database, articleModel)
	return err
}
