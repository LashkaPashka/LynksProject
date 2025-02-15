package api

import (
	"Lynks/shortener/internal/db"
	"Lynks/shortener/internal/model"
	"Lynks/shortener/internal/repository"
	"Lynks/shortener/pkg/logger"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	router *mux.Router
	db *db.Db
}


func New(db *db.Db) *API{
	api := &API{
		router: mux.NewRouter(),
		db: db,
	}

	api.Endpoints()
	return api
}


func (api *API) Run(addr string) error {
	return http.ListenAndServe(addr, api.router)
}

func (api *API) Endpoints(){
	api.router.HandleFunc("/hello", Create()).Methods(http.MethodGet)
}

func Create() http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var link model.Links
		repo := repository.NewLinkRepository()
	
		if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
			logger.Log.Error(
				"Failed decoding of the response",
				slog.String("Msg", err.Error()),
			)
		}
		var context = context.WithValue(context.Background(), model.HostAPI, r.Host)

		newLink := model.NewLink(context, link.Url)

		repo.CreateLinks(r.Context(), []model.Links{*newLink})
		///////////////////////////////////////////////////////
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&link)
		w.WriteHeader(http.StatusOK)
	})
}