package sentence

import "github.com/gaertnerl/translate-me.git/webserver/lib/similarity/datatypes"

type Sentence struct {
	Text string
	Id   int32
}

type UserTranslation struct {
	SentenceId int32
	Text       string
}

type UserTranslationEvaluation struct {
	Score datatypes.SimilarityScore
}
