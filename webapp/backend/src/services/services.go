package services

import (
	"github.com/gaertnerl/translate-me.git/webserver/lib/caching"
	"github.com/gaertnerl/translate-me.git/webserver/lib/communication"
	lb "github.com/gaertnerl/translate-me.git/webserver/lib/load_balancing"
	"github.com/gaertnerl/translate-me.git/webserver/lib/sentence"
	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity/service"
)

// calculation of sentence similarity is slow, therefor use load balanced similarity service that allows distributing the
// similarity calculation requests to multiple external instances
var similarityServiceLoadBalancer lb.LoadBalancer[service.ServiceIdentifier] = lb.New_RandomLB(service.ServiceIdentifier{Protocol: communication.HTTP, Host: "localhost"})
var SimilarityService service.SimilarityService = service.New_LoadBalancedWebBasedSimilarityService(similarityServiceLoadBalancer)

// calculation of sentence similarity is slow, therefor
var SentenceCache caching.Cache[sentence.CachableSentence] = caching.New_LockedCache[sentence.CachableSentence](10000)
