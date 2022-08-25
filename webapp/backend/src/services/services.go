package services

import (
	"database/sql"

	"github.com/gaertnerl/translate-me.git/webserver/lib/caching"
	"github.com/gaertnerl/translate-me.git/webserver/lib/database"
	lb "github.com/gaertnerl/translate-me.git/webserver/lib/load_balancing"
	"github.com/gaertnerl/translate-me.git/webserver/lib/sentence"
	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity"
)

// calculation of sentence similarity is slow, therefor use load balanced similarity service that allows distributing the
// similarity calculation requests to multiple external instances
var similarityServiceLoadBalancer lb.LoadBalancer[similarity.ServiceIdentifier]
var SimilarityService similarity.SimilarityService

// calculation of sentence similarity is slow, therefor cache similarity calculation
var SentenceCache caching.Cache[sentence.CachableSentence]

var DB *sql.DB

func SetupServices() {
	similarityServiceLoadBalancer = lb.New_RandomLB[similarity.ServiceIdentifier]()
	SimilarityService = similarity.New_LoadBalancedWebBasedSimilarityService(similarityServiceLoadBalancer)
	SentenceCache = caching.New_LockedCache[sentence.CachableSentence](10000)
	DB = database.DBHandle()
}
