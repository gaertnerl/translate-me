package service

import (
	"github.com/gaertnerl/translate-me.git/webserver/lib/communication"
	lb "github.com/gaertnerl/translate-me.git/webserver/lib/load_balancing"
	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity/datatypes"
)

type ServiceIdentifier struct {
	Protocol string
	Host     string
}

type LoadBalancedWebBasedSimilarityService struct {
	loadb lb.LoadBalancer[ServiceIdentifier]
}

type SimilarityJsonResp struct {
	Similarity float64 `json:"similarity"`
}

func New_LoadBalancedWebBasedSimilarityService(loadb lb.LoadBalancer[ServiceIdentifier]) *LoadBalancedWebBasedSimilarityService {
	return &LoadBalancedWebBasedSimilarityService{loadb: loadb}
}

func (ws LoadBalancedWebBasedSimilarityService) create_url(sentence_a string, sentence_b string) string {
	service := ws.loadb.Next()
	return service.Protocol + "://" + service.Host + "/similarity" + "/" + sentence_a + "/" + sentence_b
}

func (ws LoadBalancedWebBasedSimilarityService) CalcSimilarity(sentence_a string, sentence_b string) (datatypes.Similarity, error) {
	var sim datatypes.Similarity
	url := ws.create_url(sentence_a, sentence_b)
	resp := SimilarityJsonResp{}
	err := communication.GetJson(url, &resp)
	if err == nil {
		sim = resp.Similarity
	}
	return sim, err
}
