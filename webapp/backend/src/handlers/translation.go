package handlers

import (
	"net/http"

	"github.com/gaertnerl/translate-me.git/webserver/lib/sentence"
	"github.com/gaertnerl/translate-me.git/webserver/services"
	"github.com/gin-gonic/gin"
)

func Post_NextTranslation(c *gin.Context) {

	var req NextTranslationReq
	var resp NextTranslationResp

	c.Bind(&req)
	resp, exists := sentence.GetNextTranslation(services.DB, req.Translation_id, req.Translation_type)
	if !exists {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func Post_SubmitTranslation(c *gin.Context) {

	var req UserTranslationRequest
	var resp UserTranslationEvaluationResponse

	c.BindJSON(&req)
	resp, err := sentence.EvaluateUserTranslation(services.DB, services.SimilarityService, req)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, resp)
}
