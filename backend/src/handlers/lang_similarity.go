package handlers

import (
	"fmt"
	"net/http"

	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity"
	"github.com/gaertnerl/translate-me.git/webserver/services"
	"github.com/gin-gonic/gin"
)

func Get_SimilarityScore(c *gin.Context) {
	sentence_a := c.Param("sentence_a")
	sentence_b := c.Param("sentence_b")
	fmt.Println(sentence_a)
	fmt.Println(sentence_b)
	score, err := similarity.GetSimilarityScore(services.SimilarityService, sentence_a, sentence_b)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"score": score,
	})
}
