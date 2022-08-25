package handlers

import "github.com/gaertnerl/translate-me.git/webserver/lib/sentence"

type SentenceSimilarityReq struct {
	SentenceA string `url:"sentence_a"`
	SentenceB string `url:"sentence_b"`
}

type RegisterSimilarityEndpointReq struct {
	Protocol string
	Port     string
	Key      string
}

type UserTranslationRequest = sentence.UserTranslation

type NextTranslationReq = sentence.Translation
