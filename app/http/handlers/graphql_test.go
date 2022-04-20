package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"studiosol.com.br/janjitsu/roman-numbers/app/roman"
)

func TestGraphqlHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()

	search := roman.Search{}

	r.POST("/graphql", GraphqlHandler(&search))

	mutation := `{"query":"mutation {search(text:\"ABCXZZZZXC\") {number value}}"}`

	req, _ := http.NewRequest("POST", "/graphql", strings.NewReader(mutation))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Should return status OK")
	}

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if string(body) != `{"data":{"search":{"number":"CX","value":110}}}` {
		t.Errorf("Should return roman number search result")
	}
}
