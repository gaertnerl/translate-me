package handlers

import (
	"net/http"

	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity"
	"github.com/gaertnerl/translate-me.git/webserver/services"
	"github.com/gin-gonic/gin"
)

func Get_SimilarityScore(c *gin.Context) {
	var req SentenceSimilarityReq
	c.Bind(&req)
	score, err := similarity.GetSimilarityScore(services.SimilarityService, req.SentenceA, req.SentenceB)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"score": score,
	})
}
