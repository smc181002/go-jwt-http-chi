package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/smc181002/go-jwt-http-chi/middlewares"
)

type User struct {
  Uname string `json:"uname"`
  Password string `json:"password"`
}

type TokenSchema struct {
  Token string `json:"token"`
  Error error `json:"error"`
}

func GetToken (w http.ResponseWriter, r *http.Request) {
  user := User{}
  tokenJson := TokenSchema{}

  decoder := json.NewDecoder(r.Body)

  err := decoder.Decode(&user)

  if err != nil {
    http.Error(w, http.StatusText(403), 403)
    return
  }

  token, err := middlewares.GenJWT(user.Uname)
  if err != nil {
    tokenJson.Error = err
    fmt.Println("Some error")
  }

  tokenJson.Token = token

  render.JSON(w, r, tokenJson)
  // user.uname
}
