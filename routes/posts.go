package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/smc181002/go-jwt-http-chi/controllers"
	"github.com/smc181002/go-jwt-http-chi/middlewares"
)

func Posts(r chi.Router) {
  r.Get("/posts", controllers.GetPosts)
  r.With(middlewares.Authenticate).Get("/secret-posts", controllers.GetSecretPosts)
}

