package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	//"github.com/onsi/gomega/ghttp"
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

var _ = Describe("Testing GolangRestApi", func() {

	Describe("Given a request for fetching posts", func() {
		When("requesting all posts", func() {
			When("the response is successful", func() {
				It("should return the posts", func() {
					r, _ := http.NewRequest("GET", "/posts", nil)
   					w := httptest.NewRecorder()
					postController.GetPosts(w, r)
					Expect(w.Code).Should(Equal(200))
				})
			})
		})

	})
})