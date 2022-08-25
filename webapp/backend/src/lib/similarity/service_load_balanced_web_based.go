package similarity

import (
	"github.com/gaertnerl/translate-me.git/webserver/lib/communication"
	lb "github.com/gaertnerl/translate-me.git/webserver/lib/load_balancing"
)

type ServiceIdentifier struct {
	Protocol string
	Port     string
	Host     string
}

type LoadBalancedWebBasedSimilarityService struct {
	Loadb lb.LoadBalancer[ServiceIdentifier]
}

type SimilarityJsonResp struct {
	Similarity float64 `json:"similarity"`
}

func New_LoadBalancedWebBasedSimilarityService(loadb lb.LoadBalancer[ServiceIdentifier]) *LoadBalancedWebBasedSimilarityService {
	return &LoadBalancedWebBasedSimilarityService{Loadb: loadb}
}

func (ws *LoadBalancedWebBasedSimilarityService) AddService(si ServiceIdentifier) {
	ws.Loadb.Add(si)
}

func (ws *LoadBalancedWebBasedSimilarityService) create_url(sentence_a string, sentence_b string) string {
	service := ws.Loadb.Next()
	return service.Protocol + "://" + service.Host + ":" + service.Port + "/similarity" + "/" + sentence_a + "/" + sentence_b
}

func (ws *LoadBalancedWebBasedSimilarityService) CalcSimilarity(sentence_a string, sentence_b string) (Similarity, error) {
	var sim Similarity
	url := ws.create_url(sentence_a, sentence_b)
	resp := SimilarityJsonResp{}
	err := communication.GetJson(url, &resp)
	if err == nil {
		sim = resp.Similarity
	}
	return sim, err
}
