package similarity

import (
	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity/datatypes"
	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity/score"
	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity/service"
)

func GetSimilarityScore(service service.SimilarityService, sentence_a string, sentence_b string) (datatypes.SimilarityScore, error) {
	var scr datatypes.SimilarityScore
	sim, err := service.CalcSimilarity(sentence_a, sentence_b)
	if err == nil {
		scr = score.SimilarityToScore(sim)
	}
	return scr, err
}
