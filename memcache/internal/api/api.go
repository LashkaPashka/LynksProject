package api

import (
	"memCache/internal/db"
	"memCache/internal/model"
	"memCache/internal/payload"
	"memCache/internal/repository"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	redisDb *db.Db
	router *mux.Router
}

func New(db *db.Db) (*API, error) {
	api := &API{
		redisDb: db,
		router: mux.NewRouter(),
	}

	api.EndPoints()

	return api, nil
}

func (api *API) Run(addr string) error {
	return http.ListenAndServe(addr, api.router)
}

func (api *API) EndPoints() {
	api.router.HandleFunc("/Save", Save()).Methods(http.MethodPost)
	api.router.HandleFunc("/Get/{hash}", Get()).Methods(http.MethodGet)
}

func Save() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Init memoryRepository
		repo, err := repository.New()
		if err != nil {
			http.Error(w, "не получилось инициализировать данные", http.StatusBadRequest)
			return
		}

		// Response user
		var data payload.LinkCachePayload
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "не получилось сериализовать данные", http.StatusBadRequest)
			return
		}
		
		// Save data inMemory 
		if err := repo.SaveInMemory(model.LinkInMemory{
			Url: data.Url,
			Short_url: data.ShortUrl,
		}); err != nil {
			http.Error(w, "не получилось добавить данные", http.StatusBadRequest)
			return
		}

		w.Write([]byte("Данные добавились в кэш"))
		w.WriteHeader(http.StatusOK)

	})
}

func Get() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hash := mux.Vars(r)["hash"]
		// Init memoryRepository
		repo, err := repository.New()
		if err != nil {
			http.Error(w, "не получилось инициализировать данные", http.StatusBadRequest)
			return
		}

		mem, err := repo.GetInMemory(hash)
		if err != nil {
			w.WriteHeader(http.StatusConflict)
			http.Error(w, "не получилось получить данные", http.StatusBadRequest)
			return
		}

		var memResp = payload.ResponseLinkPayload{
			Url: mem.Url,
			ShortUrl: mem.Short_url,
		}

		fmt.Println(mem)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&memResp)
		w.WriteHeader(http.StatusOK)
	})
}