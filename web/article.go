package web

import (
	"encoding/json"
	"html"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rudbast/exercise-graphql-docker/article"
	"github.com/rudbast/exercise-graphql-docker/database"
)

type (
	ArticleForm struct {
		Title     string `json:"title"`
		Content   string `json:"content"`
		Thumbnail string `json:"thumbnail"`
	}

	ArticleHandler struct {
		articleService articleservice.ArticleService
	}
)

func (f *ArticleForm) Sanitize() {
	f.Title = html.EscapeString(f.Title)
	f.Content = html.EscapeString(f.Content)
	f.Thumbnail = html.EscapeString(f.Thumbnail)
}

func NewArticleHandler(r *mux.Router) {
	articleHandler := ArticleHandler{
		articleService: articleservice.New(database.Get()),
	}

	r.Handle("/api/v1/articles", HandlerFunc(articleHandler.GetArticleListHandler)).Methods("GET")
	r.Handle("/api/v1/articles", HandlerFunc(articleHandler.AddArticleHandler)).Methods("POST")
	r.Handle("/api/v1/articles/{id:[0-9]+}", HandlerFunc(articleHandler.GetArticleHandler)).Methods("GET")
	r.Handle("/api/v1/articles/{id:[0-9]+}", HandlerFunc(articleHandler.UpdateArticleHandler)).Methods("POST")
}

func (h ArticleHandler) GetArticleListHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	params := map[string]interface{}{}
	if title := strings.TrimSpace(r.FormValue("title")); title != "" {
		params["title"] = html.EscapeString(title)
	}

	return h.articleService.GetArticleList(ctx, params)
}

func (h ArticleHandler) GetArticleHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return nil, err
	}

	return h.articleService.GetArticle(ctx, id)
}

func (h ArticleHandler) AddArticleHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()

	var data ArticleForm

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, err
	}

	data.Sanitize()
	return nil, h.articleService.AddArticle(ctx, data.Title, data.Content, data.Thumbnail)
}

func (h ArticleHandler) UpdateArticleHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		return nil, err
	}

	var data ArticleForm

	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return nil, err
	}

	data.Sanitize()
	return nil, h.articleService.UpdateArticle(ctx, id, data.Title, data.Content, data.Thumbnail)
}
