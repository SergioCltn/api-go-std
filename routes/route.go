package routes

import (
	"net/http"

	"github.com/sergiocltn/api-go-std/controllers"
	"github.com/sergiocltn/api-go-std/routes/middleware"
)

type Router struct {
	controller controllers.Controller
}

func NewRouter(c controllers.Controller) *Router {
	return &Router{controller: c}
}

func (ro *Router) SetupRoutes() http.Handler {
	mux := http.NewServeMux()

	usersHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			ro.controller.CreateUser(w, r)
		case http.MethodGet:
			ro.controller.GetUser(w, r)
		}
	})

	mux.Handle("/user", middleware.NewLogMiddleware(usersHandler))
	return mux
}
