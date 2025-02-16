package api

import (
	"Lynks/user/internal/db"
	"Lynks/user/internal/payload"
	"Lynks/user/internal/repository"
	"encoding/json"
	"log"
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
		var user payload.UserRequest

		repo := repository.NewUserRepository()

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Fatal(err)
		}
		
		if err := repo.InsertDocs(repo.MongoClient, user.Email, user.Password, user.Name); err != nil {
			log.Fatal(err)
		}

	})
}