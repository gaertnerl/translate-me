package sentence

import (
	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity"
	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity/service"
)

var testSentenceDE = "Du hast gestern unglaublich effizient gearbeitet."
var testSentenceEN = "You have worked unbelievably efficient yesterfay."

func NextSentence() (Sentence, error) {
	s := Sentence{Text: testSentenceEN, Id: 0}
	return s, nil
}

func EvaluateUserTranslation(service service.SimilarityService, ut UserTranslation) (UserTranslationEvaluation, error) {
	var ute UserTranslationEvaluation
	score, err := similarity.GetSimilarityScore(service, testSentenceDE, ut.Text)
	if err != nil {
		return ute, nil
	}
	ute.Score = score
	return ute, nil
}
