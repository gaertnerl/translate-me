package main

import (
	"github.com/gaertnerl/translate-me.git/webserver/handlers"
	"github.com/gaertnerl/translate-me.git/webserver/lib/env"
	"github.com/gaertnerl/translate-me.git/webserver/services"
	"github.com/gin-gonic/gin"
)

func main() {
	env.SetupEnv()
	services.SetupServices()
	r := gin.Default()
	r.Static("/app", "./static")
	r.POST("/api/sentence/translation", handlers.Post_NextTranslation)
	r.POST("/api/sentence/submit", handlers.Post_SubmitTranslation)
	r.GET("/api/similarity/:sentence_a/:sentence_b", handlers.Get_SimilarityScore)
	r.POST("/api/load-balancing/register", handlers.Post_RegisterSimilarityEndpoint)
	r.Run()
}
