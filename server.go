package main

import (
	"fmt"
	"log"
	"net/http"

	//"github.com/gorilla/mux"
	// router "./http"
	// controller "./controller"
	"github.com/vipindasvg/golang-rest-api/controller"
	"github.com/vipindasvg/golang-rest-api/router"
)

var (
	postController controller.PostController = controller.NewPostController()
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {

	const port string = ":8000"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and Runnnin....")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	log.Println("Server Listening on Port ", port)
	// log.Fatalln(http.ListenAndServe(port, router))
	httpRouter.SERVE(port)
}
