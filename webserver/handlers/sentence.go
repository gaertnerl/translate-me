package handlers

import (
	"net/http"

	"github.com/gaertnerl/translate-me.git/webserver/lib/sentence"
	"github.com/gin-gonic/gin"
)

func Post_nextSentence(c *gin.Context) {
	s, err := sentence.NextSentence()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, s)
}

func Post_submitTranslation(c *gin.Context) {
	var ut sentence.UserTranslation
	c.BindJSON(ut)
	ute, err := sentence.EvaluateUserTranslation(ut)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, ute)
}