package main_test

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"github.com/vipindasvg/golang-rest-api/router"
	"github.com/vipindasvg/golang-rest-api/controller"
	//"github.com/vipindasvg/golang-rest-api/entity"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
	postController controller.PostController = controller.NewPostController()
)


func TestGolangRestApi(t *testing.T) {
	Convey("Subject: Testing Api Requests", t, func() {
		Convey("Given a requests for get posts", func() {
			r, _ := http.NewRequest("GET", "/posts", nil)
   			w := httptest.NewRecorder()
			postController.GetPosts(w, r)
			Convey("When requests run successfully", func() {
				So(w.Code, ShouldEqual, 200)
				Convey("Then User will get the posts", nil)
			})
		})
	})
}