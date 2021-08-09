package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vipindasvg/golang-rest-api/controller"
	"github.com/vipindasvg/golang-rest-api/router"
	"net/http"
	"net/http/httptest"
)

var (
	httpRouter     router.Router             = router.NewMuxRouter()
	postController controller.PostController = controller.NewPostController()
)

func TestGetRequest(t *testing.T) {
	t.Parallel()

	r, _ := http.NewRequest("GET", "/posts", nil)
	w := httptest.NewRecorder()

	postController.GetPosts(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}
