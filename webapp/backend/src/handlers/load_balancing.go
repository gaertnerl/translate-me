package handlers

import (
	"fmt"
	"net/http"

	"github.com/gaertnerl/translate-me.git/webserver/lib/env"
	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity"
	"github.com/gaertnerl/translate-me.git/webserver/services"
	"github.com/gin-gonic/gin"
)

func Post_RegisterSimilarityEndpoint(c *gin.Context) {

	var req RegisterSimilarityEndpointReq
	c.BindJSON(&req)
	key := env.Env.REGISTER_SIM_ENDPOINT_KEY

	if req.Key != key {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	simS := services.SimilarityService
	_, ok := simS.(*similarity.LoadBalancedWebBasedSimilarityService)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	si := similarity.ServiceIdentifier{Protocol: req.Protocol, Port: req.Port, Host: c.ClientIP()}
	fmt.Println(si)
	simS.(*similarity.LoadBalancedWebBasedSimilarityService).Loadb.Add(si)
	c.Status(http.StatusOK)
}
