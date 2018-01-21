package model

import (
	"context"
	"database/sql"
	"fmt"
)

type (
	Article struct {
		ID        int64  `json:"id"`
		Title     string `json:"title"`
		Content   string `json:"content"`
		Thumbnail string `json:"thumbnail"`
	}
)

func GetArticleList(ctx context.Context, db *sql.DB) ([]Article, error) {
	rows, err := db.QueryContext(ctx, stmtQueryAllArticles)
	if err != nil {
		return nil, fmt.Errorf("database query article list error: { %+v }", err)
	}
	defer rows.Close()

	articleList := []Article{}

	for rows.Next() {
		var article Article

		err = rows.Scan(&article.ID, &article.Title, &article.Content, &article.Thumbnail)
		if err != nil {
			return nil, fmt.Errorf("database scan rows error: { %+v }", err)
		}

		articleList = append(articleList, article)
	}

	return articleList, nil
}

func GetArticle(ctx context.Context, db *sql.DB, id int64) (*Article, error) {
	var article Article

	err := db.QueryRowContext(ctx, stmtQueryArticleByID, id).Scan(&article.ID,
		&article.Title, &article.Content, &article.Thumbnail)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, fmt.Errorf("database query article by id error: { %+v }", err)
	}

	return &article, nil
}

func GetArticleListByTitle(ctx context.Context, db *sql.DB, title string) ([]Article, error) {
	rows, err := db.QueryContext(ctx, stmtQueryArticlesByTitle, title)
	if err != nil {
		return nil, fmt.Errorf("database query article list by title error: { %+v }", err)
	}
	defer rows.Close()

	articleList := []Article{}

	for rows.Next() {
		var article Article

		err = rows.Scan(&article.ID, &article.Title, &article.Content, &article.Thumbnail)
		if err != nil {
			return nil, fmt.Errorf("database scan rows error: { %+v }", err)
		}

		articleList = append(articleList, article)
	}

	return articleList, nil
}

func AddArticle(ctx context.Context, db *sql.DB, article Article) (int64, error) {
	res, err := db.ExecContext(ctx, stmtInsertArticle,
		article.Title, article.Content, article.Thumbnail)
	if err != nil {
		return 0, fmt.Errorf("database insert article error: { %+v }", err)
	}

	return res.LastInsertId()
}

func UpdateArticle(ctx context.Context, db *sql.DB, article Article) (int64, error) {
	res, err := db.ExecContext(ctx, stmtUpdateArticle,
		article.Title, article.Content, article.Thumbnail, article.ID)
	if err != nil {
		return 0, fmt.Errorf("database update article error: { %+v }", err)
	}

	return res.RowsAffected()
}
