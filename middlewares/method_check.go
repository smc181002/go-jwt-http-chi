package middlewares

import (
	"github.com/go-chi/chi/v5"
  "net/http"
)

// Chi by default doesnt show message for 405 error text
// This middleware checks the gloobal router context and
// displays an error message of "Method not Allowed"
func MethodCheck(next http.Handler) http.Handler {
  return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
      // get the context from the given router
      rctx := chi.RouteContext(r.Context())

      // Temporary context
      ctx := chi.NewRouteContext()
      // Matching the router context with the current request.
      // If there is no method matching for the current 
      // request route, throw method not allowed (405) err.
      if !rctx.Routes.Match(ctx, r.Method, r.URL.Path) {
        http.Error(w, http.StatusText(405), 405)
        return
      }
      next.ServeHTTP(w, r.WithContext(r.Context()))
  })
}
// No need to worry about the cors pre-flight checks as 
// Chi router adds the OPTIONS method to the router context
