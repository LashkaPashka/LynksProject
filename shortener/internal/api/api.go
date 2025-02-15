package api

import (
	"Lynks/shortener/internal/db"
	"net/http"

	"github.com/gorilla/mux"
)


type API struct{
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

}