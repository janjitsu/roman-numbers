package handlers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"studiosol.com.br/janjitsu/roman-numbers/app/http/graph"
	"studiosol.com.br/janjitsu/roman-numbers/app/http/graph/generated"
	"studiosol.com.br/janjitsu/roman-numbers/app/roman"
)

func GraphqlHandler(search roman.Finder) gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{search}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
