package handlers

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

//Playground handler enables graphql insception on Web Browser
func PlaygroundHandler(route string) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := playground.Handler("GraphQL", route)

		h.ServeHTTP(c.Writer, c.Request)
	}
}
