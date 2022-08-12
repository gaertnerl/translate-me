package service

import (
	"github.com/gaertnerl/translate-me.git/webserver/lib/communication"
	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity/datatypes"
)

type WebBasedSimilarityService struct {
	Protocol string
	Host     string
}

type SimilarityJsonResp struct {
	Similarity float64 `json:"similarity"`
}

func (ws WebBasedSimilarityService) create_url(sentence_a string, sentence_b string) string {
	return ws.Protocol + "://" + ws.Host + "/similarity" + "/" + sentence_a + "/" + sentence_b
}

func (ws WebBasedSimilarityService) CalcSimilarity(sentence_a string, sentence_b string) (datatypes.Similarity, error) {
	var sim datatypes.Similarity
	url := ws.create_url(sentence_a, sentence_b)
	resp := SimilarityJsonResp{}
	err := communication.GetJson(url, &resp)
	if err == nil {
		sim = resp.Similarity
	}
	return sim, err
}
