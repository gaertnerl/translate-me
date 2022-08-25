package similarity

const (
	FAILED  SimilarityScore = 0
	GOOD    SimilarityScore = 100
	PERFECT SimilarityScore = 500
)

func SimilarityToScore(similarity Similarity) SimilarityScore {
	if similarity > 0.9 {
		return PERFECT
	} else if similarity > 0.8 {
		return GOOD
	} else {
		return FAILED
	}
}
