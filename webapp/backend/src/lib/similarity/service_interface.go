package similarity

type SimilarityService interface {
	CalcSimilarity(sentence_a string, sentence_b string) (Similarity, error)
}
