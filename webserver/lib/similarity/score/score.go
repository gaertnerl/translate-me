package score

import (
	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity/datatypes"
)

const (
	FAILED  datatypes.SimilarityScore = 0
	GOOD    datatypes.SimilarityScore = 100
	PERFECT datatypes.SimilarityScore = 500
)

func SimilarityToScore(similarity datatypes.Similarity) datatypes.SimilarityScore {
	if similarity > 0.9 {
		return PERFECT
	} else if similarity > 0.7 {
		return GOOD
	} else {
		return FAILED
	}
}
