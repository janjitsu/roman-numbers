package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGraphqlPlayground(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()

	r.GET("/", PlaygroundHandler("/graphql"))

	req, _ := http.NewRequest("GET", "/", nil)

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Should return status OK")
	}
}
