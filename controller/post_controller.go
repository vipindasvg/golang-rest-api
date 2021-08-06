package controller

import (
	"math/rand"
	"encoding/json"
	"net/http"
	
	"github.com/vipindasvg/golang-rest-api/entity"
	"github.com/vipindasvg/golang-rest-api/service"
)

var (
	postService service.PostService = service.NewPostService()
)

type controller struct{}

type PostController interface {
	AddPost(resp http.ResponseWriter, req *http.Request)
	GetPosts(resp http.ResponseWriter, req *http.Request) 
}

func NewPostController() PostController {
	return &controller{}
}

func (*controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "Application/json")
	posts, err := postService.FindAll() 
	_, err = json.Marshal(posts)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error":"Error Marshalling Post Array"}`))
		return 
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func (*controller) AddPost(resp http.ResponseWriter, req *http.Request)  {
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error":"Error Marshalling Post Array"}`))
		return
	}
	post.ID = rand.Int63()
	err =postService.Validate(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error":"Validation error"}`))
		return  
	}
	result, err :=	postService.Create(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error":"creation error"}`))
		return 
	}
	//postService.Create(&post)
	resp.WriteHeader(http.StatusOK)
	results, _ := json.Marshal(result)
	resp.Write(results)
}