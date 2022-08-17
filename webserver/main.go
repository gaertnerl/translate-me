package main

import (
	"github.com/gaertnerl/translate-me.git/webserver/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/app", "./static")
	r.POST("/api/sentence/next", handlers.Post_nextSentence)
	r.POST("/api/sentence/submit", handlers.Post_submitTranslation)
	r.GET("/api/similarity/:sentence_a/:sentence_b", handlers.Get_SimilarityScore)
	r.POST("/api/load-balancing/register", handlers.Post_RegisterSimilarityEndpoint)
	r.Run()
}
