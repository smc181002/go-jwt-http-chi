package main

import (
	"flag"
	"fmt"
	"net/http"

	// "github.com/smc181002/go-jwt-http/"
	"github.com/go-chi/chi/v5"
  "github.com/go-chi/cors"

)

func AllowMethod(next http.Handler) http.Handler {
  return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
      // get the context from the given router
      rctx := chi.RouteContext(r.Context())

      // Temporary context
      ctx := chi.NewRouteContext()
      fmt.Println(rctx.Routes.Match(ctx, r.Method, r.URL.Path))
      next.ServeHTTP(w, r.WithContext(r.Context()))
      fmt.Println("something")
      /* fmt.Println("Options")
      next.ServeHTTP(w, r.WithContext(r.Context())) */
  })
}

var port int
func main() {
  flag.IntVar(&port, "p", 8080, "Enter the port number for the server")
  flag.Parse()
  app := chi.NewRouter()

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

  app.Use(AllowMethod)

  app.Get("/", func (w http.ResponseWriter, r *http.Request)  {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Hello there - GET");
  })

  app.Post("/", func (w http.ResponseWriter, r *http.Request)  {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Hello there - POST");
  })

  /* r := mux.NewRouter()

  r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, "Hello there");
  }).Methods("GET")
 */

  fmt.Printf("server listening on port %v\n", port)
  http.ListenAndServe(fmt.Sprintf(":%v", port), app)

  fmt.Println("Gracefully Shutting down")
}

