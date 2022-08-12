package main

import (
	"github.com/gaertnerl/translate-me.git/webserver/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/similarity/:sentence_a/:sentence_b", handlers.Get_SimilarityScore)
	r.Run()
}
