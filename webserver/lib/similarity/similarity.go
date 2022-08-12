package similarity

import (
	"fmt"

	"github.com/gaertnerl/translate-me.git/webserver/lib/communication"
	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity/datatypes"
	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity/score"
	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity/service"
)

func GetSimilarityScore(sentence_a string, sentence_b string) (datatypes.SimilarityScore, error) {
	var scr datatypes.SimilarityScore
	service := service.WebBasedSimilarityService{Protocol: communication.HTTP, Host: "localhost"}
	sim, err := service.CalcSimilarity(sentence_a, sentence_b)
	fmt.Println(sim)
	if err == nil {
		scr = score.SimilarityToScore(sim)
		fmt.Println(scr)
	}
	return scr, err
}
