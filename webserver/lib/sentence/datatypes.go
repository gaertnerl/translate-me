package sentence

import (
	"time"

	"github.com/gaertnerl/translate-me.git/webserver/lib/similarity/datatypes"
)

type Sentence struct {
	Text string
	Id   int32
}

type CachableSentence struct {
	Sentence Sentence
	timeout  time.Time
}

func (cs CachableSentence) Timeout() time.Time {
	return cs.timeout
}

type UserTranslation struct {
	SentenceId int32
	Text       string
}

type UserTranslationEvaluation struct {
	Score datatypes.SimilarityScore
}
