package services

import (
	"github.com/gaertnerl/translate-me.git/webserver/lib/communication"
	lb "github.com/gaertnerl/translate-me.git/webserver/lib/load_balancing"
	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity/service"
)

var similarityServiceLoadBalancer lb.LoadBalancer[service.ServiceIdentifier] = lb.New_RandomLB(service.ServiceIdentifier{Protocol: communication.HTTP, Host: "localhost"})
var SimilarityService service.SimilarityService = service.New_LoadBalancedWebBasedSimilarityService(similarityServiceLoadBalancer)
