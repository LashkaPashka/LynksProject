package api

import (
	"User/internal/db"
	"User/internal/payload"
	"User/internal/repository"
	"User/pkg/logger"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
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
	api.router.HandleFunc("/register", Register()).Methods(http.MethodPost)
	api.router.HandleFunc("/login", Login()).Methods(http.MethodPost)
	api.router.HandleFunc("/logout", Logout()).Methods(http.MethodDelete)
}

func Register() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user payload.RegisterUserRequest
		repo := repository.NewUserRepository()

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			logger.Log.Error(
				"Неудачное декодирование",
			)
			w.WriteHeader(http.StatusConflict)
			return
		}
		
		if err := repo.InsertDocs(repo.MongoClient, user.Email, user.Password, user.Name); err != nil {
			logger.Log.Error(
				"Невозможно добавить пользователя в БД",
				slog.String("email", user.Email),
				slog.String("password", user.Password),
				slog.String("name", user.Name),
			)
			w.WriteHeader(http.StatusConflict)
			return
		}

		w.Write([]byte("Пользователь добавлен"))
	})
}

func Login() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user payload.LoginUserRequest
		repo := repository.NewUserRepository()

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			logger.Log.Error(
				"Неудачаное декодирование",
			)
			return
		}
		
		existUser, _, _ := repo.GetByEmail(repo.MongoClient, user.Email)
		if err := bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(user.Password)); err != nil {
			http.Error(w, "Пользователь не существует", http.StatusConflict)
			logger.Log.Error(
				"Не удалось декодировать пароль",
				slog.String(user.Email, "Пароль не соответствует"),
			)
			return
		}
	
		stoken, _ := repo.Token.CreateJWT(existUser.Email)
		w.Header().Set("Authorization", fmt.Sprintf("Bearer %s", stoken))
		w.WriteHeader(http.StatusOK)
	})
}

func Logout() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Del("Authorization")
		w.WriteHeader(http.StatusOK)
	})
}