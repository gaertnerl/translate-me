package sentence

import (
	"time"

	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity"
)

type TranslationId = uint32

type Sentence = string

type TranslationType = uint8

var (
	EN_DE TranslationType = 0 // sentence 1 english, sentence 2 german
	EN_ES                 = 1 // sentence 1 english, sentence 2 spanish
	ES_DE                 = 2 // sentence 1 spanish, sentence 2 german
)

type Translation struct {
	Translation_id   TranslationId
	Translation_type TranslationType
	Sentence_1       Sentence
	Sentence_2       Sentence
}

type UserTranslation struct {
	Correct        Sentence
	UserTranslated Sentence
	Id             TranslationId
}

type CachableSentence struct {
	Sentence Sentence
	timeout  time.Time
}

func (cs CachableSentence) Timeout() time.Time {
	return cs.timeout
}

type UserTranslationEvaluation struct {
	Score similarity.SimilarityScore
}
