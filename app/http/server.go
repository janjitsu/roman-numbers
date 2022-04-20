package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"studiosol.com.br/janjitsu/roman-numbers/app/roman"

	"studiosol.com.br/janjitsu/roman-numbers/app/http/handlers"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	search := roman.Search{}

	router := gin.Default()
	router.POST("/graphql", handlers.GraphqlHandler(&search))
	router.GET("/", handlers.PlaygroundHandler("/graphql"))
	router.Run("0.0.0.0:" + port)
}
