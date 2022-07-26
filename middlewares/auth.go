package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var Salt = []byte("ChiAuthSalt")

type UserDataClaims struct {
  Uname string `json:"uname"`
  jwt.StandardClaims
}

func GenJWT(uname string) (string, error) {
  claims := UserDataClaims {
    uname,
    jwt.StandardClaims {
      ExpiresAt: time.Now().Add(time.Hour * 10).Unix(),
    },
  }

  token :=  jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  return token.SignedString(Salt)
}

func breakBearer(bearer string) (string, error) {
  if bearer == "" {
    return "", errors.New("Bearer token not found")
  }

  bearerSlice := strings.Split(bearer, " ")

  if (len(bearerSlice) != 2 && bearerSlice[0] != "Bearer") {
    return "", errors.New("Bearer token not found")
  }

  return bearerSlice[1], nil
}

// lets try to write monadic middlewares

func Authenticate(next http.Handler) http.Handler {
  return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
    bearer := r.Header.Get("Authorization")

    token, err := breakBearer(bearer)

    if err != nil {
      http.Error(w, fmt.Sprintf("%v. %v", http.StatusText(403), err), 403)
      return
    }

    value, err := jwt.ParseWithClaims(
      token,
      &UserDataClaims{},
      func (*jwt.Token) (interface{}, error)  {
        return Salt, nil
    })

    if claims, ok := value.Claims.(*UserDataClaims); !(ok && value.Valid && claims.Uname == "smc") {
      http.Error(w, fmt.Sprintf("%v. %v", http.StatusText(403), err), 403)
      return
    }

    if err != nil {
      http.Error(w, fmt.Sprintf("%v. %v", http.StatusText(403), err), 403)
      return
    }

    next.ServeHTTP(w, r)
  })
}
