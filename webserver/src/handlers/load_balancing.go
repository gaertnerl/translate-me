package handlers

import (
	"net/http"

	"github.com/gaertnerl/translate-me.git/webserver/lib/env"
	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity/service"
	"github.com/gaertnerl/translate-me.git/webserver/services"
	"github.com/gin-gonic/gin"
)

type RegisterSimilarityEndpointReq struct {
	Service service.ServiceIdentifier
	Key     string
}

func Post_RegisterSimilarityEndpoint(c *gin.Context) {

	var req RegisterSimilarityEndpointReq
	c.BindJSON(&req)
	key := env.Env.REGISTER_SIM_ENDPOINT_KEY

	if req.Key != key {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	simS := services.SimilarityService
	_, ok := simS.(*service.LoadBalancedWebBasedSimilarityService)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	simS.(*service.LoadBalancedWebBasedSimilarityService).Loadb.Add(req.Service)
	c.Status(http.StatusOK)
}
