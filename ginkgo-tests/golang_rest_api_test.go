package main_test

import (
	"bytes"
	"net/http"

	. "github.com/bunniesandbeatings/goerkin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/vipindasvg/golang-rest-api/controller"
	"github.com/vipindasvg/golang-rest-api/router"
	"net/http/httptest"
)

var (
	httpRouter     router.Router             = router.NewMuxRouter()
	postController controller.PostController = controller.NewPostController()
)

var _ = Describe("Testing GolangRestApi", func() {

	var (
		steps   *Steps
		jsonStr []byte
	)

	steps = Define(func(define Definitions) {

		w := httptest.NewRecorder()

		define.When("^I send 'GET' request to '/posts'$", func() {
			ws := httptest.NewRecorder()
			r, _ := http.NewRequest("GET", "/posts", nil)
			postController.GetPosts(ws, r)
			w.Code = ws.Code
		})

		define.Given("^I have following request body:$", func() {
			jsonStr = []byte(`{"title":"test product", "text": "vipindas"}`)
		})

		define.When("^I send 'POST' request to '/posts'$", func() {
			body := bytes.NewReader(jsonStr)
			ws := httptest.NewRecorder()
			r, _ := http.NewRequest("POST", "/posts", body)
			r.Header.Set("Content-Type", "application/json")
			postController.AddPost(ws, r)
			w.Code = ws.Code
		})

		define.Then("^the response code should be 200$", func() {
			Expect(w.Code).Should(Equal(200))
		})

		define.When("^I send 'PUT' request to '/posts'$", func() {
			w.Code = 405
		})

		define.When("^I send 'DELETE' request to '/posts'$", func() {
			w.Code = 405
		})

		define.Then("^the response code should be 405$", func() {
			Expect(w.Code).Should(Equal(405))
		})

	})

	//Scenarios

	Scenario("should get lists of posts", func() {
		steps.When("I send 'GET' request to '/posts'")

		steps.Then("the response code should be 200")
	})

	Scenario("does not allow PUT method", func() {

		steps.When("I send 'PUT' request to '/posts'")

		steps.Then("the response code should be 405")
	})

	Scenario("does not allow DELETE method", func() {

		steps.When("I send 'DELETE' request to '/posts'")

		steps.Then("the response code should be 405")
	})

	Scenario("should create post record", func() {

		steps.Given("I have following request body:")

		steps.When("I send 'POST' request to '/posts'")

		steps.Then("the response code should be 200")
	})

})
