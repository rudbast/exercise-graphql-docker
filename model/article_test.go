package model

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetArticleList(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows(
		[]string{"id", "title", "content", "thumbnail"},
	).AddRow(
		1, "foo", "bar", "baz",
	).AddRow(
		2, "bar", "baz", "foo",
	)

	mock.ExpectQuery("SELECT (.+) FROM articles").WillReturnRows(rows)

	_, err = GetArticleList(context.Background(), db)
	assert.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetArticle(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	articleID := int64(1)

	rows := sqlmock.NewRows(
		[]string{"id", "title", "content", "thumbnail"},
	).AddRow(
		articleID, "foo", "bar", "baz",
	)

	mock.ExpectQuery("SELECT (.+) FROM articles WHERE id = ?").WithArgs(articleID).WillReturnRows(rows)

	article, err := GetArticle(context.Background(), db, articleID)
	assert.NoError(t, err)
	assert.NotNil(t, article)

	require.NoError(t, mock.ExpectationsWereMet())

	rows = sqlmock.NewRows(
		[]string{"id", "title", "content", "thumbnail"},
	)

	// Extra test for err no rows (should not be tampered).
	mock.ExpectQuery("SELECT (.+) FROM articles WHERE id = ?").WithArgs(0).WillReturnRows(rows)

	article, err = GetArticle(context.Background(), db, 0)
	assert.Equal(t, sql.ErrNoRows, err)
}

func TestGetArticleListByTitle(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows(
		[]string{"id", "title", "content", "thumbnail"},
	).AddRow(
		1, "foo", "bar", "baz",
	)

	mock.ExpectQuery("SELECT (.+) FROM articles WHERE (.+)").WillReturnRows(rows)

	_, err = GetArticleListByTitle(context.Background(), db, "foo")
	assert.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestAddArticle(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	article := Article{0, "foo", "bar", "baz"}

	mock.ExpectExec("INSERT INTO articles").WithArgs(
		article.Title, article.Content, article.Thumbnail,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = AddArticle(context.Background(), db, article)
	assert.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateArticle(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	article := Article{1, "foo", "bar", "baz"}

	mock.ExpectExec("UPDATE articles").WithArgs(
		article.Title, article.Content, article.Thumbnail, article.ID,
	).WillReturnResult(sqlmock.NewResult(0, 1))

	_, err = UpdateArticle(context.Background(), db, article)
	assert.NoError(t, err)

	require.NoError(t, mock.ExpectationsWereMet())
}
