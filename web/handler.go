package web

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rudbast/exercise-graphql-docker/config"
	"github.com/rudbast/exercise-graphql-docker/jsonapi"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request) (interface{}, error)

func (fn HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	timeout, err := time.ParseDuration(config.Get().Server.Timeout)
	if err != nil {
		timeout = time.Second * 5
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	r = r.WithContext(ctx)

	res, err := fn(w, r)
	if err != nil {
		// TODO: Properly map http status on error.
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := jsonapi.Response{
		Data: res,
	}
	byteResp, _ := json.Marshal(resp)

	w.Header().Set("Content-Type", "application/json")
	w.Write(byteResp)
}
