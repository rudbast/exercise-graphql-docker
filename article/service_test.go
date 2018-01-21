package articleservice_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestServiceGetArticleList(t *testing.T) {
	rows := sqlmock.NewRows(
		[]string{"id", "title", "content", "thumbnail"},
	).AddRow(
		1, "foo", "bar", "baz",
	).AddRow(
		2, "bar", "baz", "foo",
	)
	mockObj.ExpectQuery("SELECT").WillReturnRows(rows)

	resp, err := service.GetArticleList(context.Background(), map[string]interface{}{})
	assert.NoError(t, err)
	assert.Equal(t, 2, len(resp))
}

func TestServiceGetArticle(t *testing.T) {
	articleID := int64(1)

	rows := sqlmock.NewRows(
		[]string{"id", "title", "content", "thumbnail"},
	).AddRow(
		articleID, "foo", "bar", "baz",
	)
	mockObj.ExpectQuery("SELECT").WillReturnRows(rows)

	article, err := service.GetArticle(context.Background(), articleID)
	assert.NoError(t, err)
	assert.Equal(t, fmt.Sprintf("%d", articleID), article.ID)
}

func TestServiceAddArticle(t *testing.T) {
	mockObj.ExpectExec("INSERT").WithArgs(
		"foo", "baz", "bar",
	).WillReturnResult(sqlmock.NewResult(1, 1))

	err := service.AddArticle(context.Background(), "foo", "baz", "bar")
	assert.NoError(t, err)
}

func TestServiceUpdateArticle(t *testing.T) {
	mockObj.ExpectExec("UPDATE").WithArgs(
		"foo", "baz", "bar", 1,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	err := service.UpdateArticle(context.Background(), int64(1), "foo", "baz", "bar")
	assert.NoError(t, err)
}
