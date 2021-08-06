package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"github.com/vipindasvg/golang-rest-api/router"
	"github.com/vipindasvg/golang-rest-api/controller"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
	postController controller.PostController = controller.NewPostController()
)

func TestGetRequest(t *testing.T) {
    t.Parallel()

    r, _ := http.NewRequest("GET", "/posts", nil)
    w := httptest.NewRecorder()


    postController.GetPosts(w, r)

    assert.Equal(t, http.StatusOK, w.Code)
}