package sentence

import "github.com/gaertnerl/translate-me.git/webserver/lib/similarity"

var testSentenceDE = "Du hast gestern unglaublich effizient gearbeitet."
var testSentenceEN = "You have worked unbelievably efficient yesterfay."

func NextSentence() (Sentence, error) {
	s := Sentence{Text: testSentenceEN, Id: 0}
	return s, nil
}

func EvaluateUserTranslation(ut UserTranslation) (UserTranslationEvaluation, error) {
	var ute UserTranslationEvaluation
	score, err := similarity.GetSimilarityScore(testSentenceDE, ut.Text)
	if err != nil {
		return ute, nil
	}
	ute.Score = score
	return ute, nil
}
