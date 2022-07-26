package controllers

import (
	"fmt"
	"math/rand"
	"net/http"

  "github.com/go-chi/render"
)

type Post struct {
  Title string `json:"title"`
  Description string `json:"description"`
  Likes int32 `json:"likes"`
  Comments int32 `json:"comments"`
  ShareLink string `json:"shareLink"`
}

func GetPosts (w http.ResponseWriter, r *http.Request)  {
  w.Header().Add("Accept", "application/json")

  dataList := []Post{ }
  var data Post
  const n int = 4
  for i:=0; i < n ; i++ {
    data = Post{
      Title: fmt.Sprintf("JWT auth with golang - %v", i+1),
      Description: "We are working on a new project template with golang with jwt authentication included",
      Likes: rand.Int31n(1024),
      Comments: rand.Int31n(1024),
      ShareLink: "https://github.com/smc181002/smc181002",
    }
    dataList = append(dataList, data)
  }
  
  render.JSON(w, r, dataList)
}

func GetSecretPosts (w http.ResponseWriter, r *http.Request)  {
  w.Header().Add("Accept", "application/json")

  dataList := []Post{ }
  var data Post
  const n int = 4
  for i:=0; i < n ; i++ {
    data = Post{
      Title: fmt.Sprintf("Secret Mission from Agent Anya - %v", i+1),
      Description: "This Mission is sent by agent Anya from the Norden HeadQuarters.",
      Likes: rand.Int31n(1024),
      Comments: rand.Int31n(1024),
      ShareLink: "https://spyhq.com/forger/anya",
    }
    dataList = append(dataList, data)
  }
  
  render.JSON(w, r, dataList)
}
