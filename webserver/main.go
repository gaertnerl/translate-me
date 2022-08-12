package main

import (
	"github.com/gaertnerl/translate-me.git/webserver/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/app", "./static")
	r.GET("/api/similarity/:sentence_a/:sentence_b", handlers.Get_SimilarityScore)
	r.Run()
}
