package similarity

func GetSimilarityScore(service SimilarityService, sentence_a string, sentence_b string) (SimilarityScore, error) {
	var scr SimilarityScore
	sim, err := service.CalcSimilarity(sentence_a, sentence_b)
	if err == nil {
		scr = SimilarityToScore(sim)
	}
	return scr, err
}
