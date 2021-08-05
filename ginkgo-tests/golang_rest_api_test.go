package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"net/http"
	"net/http/httptest"
	// router "../http"
	// controller "../controller"
	// entity "../entity"
	"github.com/vipindasvg/golang-rest-api/router"
	"github.com/vipindasvg/golang-rest-api/controller"
	"github.com/vipindasvg/golang-rest-api/entity"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
	postController controller.PostController = controller.NewPostController()
)

var _ = Describe("Testing GolangRestApi", func() {
	var server *ghttp.Server

	BeforeEach(func() {
		server = ghttp.NewServer()
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("fetching posts", func() {
		var statusCode int
		var posts []entity.Post

		BeforeEach(func() {
			statusCode = http.StatusOK
			posts = []entity.Post{}
			server.AppendHandlers(ghttp.CombineHandlers(
				ghttp.VerifyRequest("GET", "/posts"),
				ghttp.RespondWithJSONEncodedPtr(&statusCode, &posts),
			))
		})

		When("requesting all posts", func() {
			When("the response is successful", func() {
				BeforeEach(func() {
					posts = []entity.Post{}
				})
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