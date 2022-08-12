package service

import "github.com/gaertnerl/translate-me.git/webserver/lib/similarity/datatypes"

type SimilarityService interface {
	CalcSimilarity(sentence_a string, sentence_b string, similarity *datatypes.Similarity) (datatypes.Similarity, error)
}
