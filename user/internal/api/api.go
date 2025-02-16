package api

import (
	"Lynks/user/internal/db"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	mongoDb *db.MongoDb
	router *mux.Router
}

func New(mongoDb *db.MongoDb) *API {
	api := &API{
		mongoDb: mongoDb,
		router: mux.NewRouter(),
	}

	api.EndPoints()
	return api
}

func (api *API) Run(addr string) error {
	return http.ListenAndServe(addr, api.router)
}

func (api *API) EndPoints() {
	api.router.HandleFunc("/register", register()).Methods(http.MethodPost)
	api.router.HandleFunc("/login", nil).Methods(http.MethodPost)
	api.router.HandleFunc("/logout", nil).Methods(http.MethodPost)
}

func register() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, man!"))
	})
}