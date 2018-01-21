package articleservice_test

import (
	"log"
	"os"
	"testing"

	"github.com/rudbast/exercise-graphql-docker/article"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var (
	mockObj sqlmock.Sqlmock
	service articleservice.ArticleService
)

func TestMain(m *testing.M) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal("Initiate database mock error:", err)
		return
	}

	mockObj = mock
	service = articleservice.New(db)

	os.Exit(m.Run())
}
