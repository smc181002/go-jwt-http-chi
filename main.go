package main

import (
	"flag"
	"fmt"
	"net/http"

	// "github.com/smc181002/go-jwt-http-chi/middlewares"
	"github.com/smc181002/go-jwt-http-chi/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
  "github.com/go-chi/cors"
)

var port int

/* func basicGet (w http.ResponseWriter, r *http.Request)  {
  w.WriteHeader(http.StatusOK)
  fmt.Fprint(w, "Hello there - GET");
}

func basicPost(w http.ResponseWriter, r *http.Request)  {
  w.WriteHeader(http.StatusOK)
  fmt.Fprint(w, "Hello there - POST");
} */


func main() {
  flag.IntVar(&port, "p", 8080, "Enter the port number for the server")
  flag.Parse()
  app := chi.NewRouter()

  app.Use(middleware.RequestID)
  app.Use(middleware.RealIP)
  app.Use(middleware.Logger)
  app.Use(middleware.Recoverer)

  app.Use(cors.Handler(cors.Options{
    // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
    // AllowedOrigins:   []string{"https://*", "http://*"},
    AllowedOrigins:   []string{"*"},
    // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))

  // app.Use(middlewares.MethodCheck)
  app.Group(routes.Posts)
  app.Group(routes.Auth)

  /* app.Get("/", basicGet)

  app.With(middlewares.Authenticate).Post("/", basicPost) */

  fmt.Printf("server listening on port %v\n", port)
  http.ListenAndServe(fmt.Sprintf(":%v", port), app)

  fmt.Println("Gracefully Shutting down")
}

