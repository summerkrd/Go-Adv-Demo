package auth

import (
	"adv-demo/configs"
	"adv-demo/pkg/res"
	"fmt"
	"net/http"
)

type AuthHandlerDeps struct {
	*configs.Config
}
type AuthHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Config: deps.Config,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Login")
		fmt.Println(handler.Config.Auth.Secret)

		data := LoginResponce{
			Token: "123",
		}

		res.Json(w, &data, 200)
	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Register")
	}
}
