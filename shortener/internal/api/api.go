package api

import (
	"ShorteNer/configs"
	"ShorteNer/internal/client"
	"ShorteNer/internal/db"
	"ShorteNer/internal/model"
	"ShorteNer/internal/payload"
	"ShorteNer/internal/repository"
	"ShorteNer/pkg/logger"
	"ShorteNer/pkg/middleware"
	"ShorteNer/pkg/res"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type API struct {
	router *mux.Router
	db *db.Db
	conf *configs.Config
}


func New(db *db.Db, conf *configs.Config) *API{
	api := &API{
		router: mux.NewRouter(),
		db: db,
		conf: conf,
	}

	api.Endpoints()
	return api
}


func (api *API) Run(addr string) error {
	return http.ListenAndServe(addr, api.router)
}

func (api *API) Endpoints(){
	api.router.Handle("/Create", middleware.IsAuthed(Create(), api.conf)).Methods(http.MethodPost)
	api.router.HandleFunc("/{hash}", GoTo()).Methods(http.MethodGet)
	api.router.Handle("/{hash}", middleware.IsAuthed(Delete(), api.conf)).Methods(http.MethodDelete)
}

func Create() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repo := repository.NewLinkRepository()
		var link payload.LinkRequest

		if err := json.NewDecoder(r.Body).Decode(&link); err != nil {
			logger.Log.Error(
				"Failed decoding of the response",
				slog.String("Msg", err.Error()),
			)
		}
		
		newLink := model.NewLink(link.Destination)
		repo.CreateLinks(r.Context(), newLink)
		
		res.Encode(w, &payload.LinkResponse{
			ShortUrl: "http://" + r.Host + "/" + newLink.Hash,
			Destination: newLink.Url,
		}, http.StatusOK)
	})
}

func GoTo() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hash := mux.Vars(r)["hash"]
		repo := repository.NewLinkRepository()

		ctx := context.WithValue(r.Context(), repository.Hash, hash)
		
		// Get url from Cache
		var url string
		mp, err := client.GetCache(hash)
		if err != nil {
			url, err = repo.GetLinks(ctx)
			if err != nil {
				logger.Log.Error(
					"Failed decoding of the response",
					slog.String("Msg", err.Error()),
				)
			}
		} else {
			url = mp["url"]
		}

		repo.Kafka.Producer(map[string]string{
			"url": url,
			"hash": hash,
		})

		http.Redirect(w, r, url, http.StatusOK)
	})
}

func Delete() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repo := repository.NewLinkRepository()
		hash := mux.Vars(r)["hash"]
		
		ctx := context.WithValue(r.Context(), repository.Hash, hash)
		repo.DeleteLinks(ctx)

		w.Write([]byte("Link removed"))
		w.WriteHeader(http.StatusOK)
	})
}