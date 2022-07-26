package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/smc181002/go-jwt-http-chi/controllers"
)

func Auth(r chi.Router)  {
  r.Post("/get-token", controllers.GetToken)
}
