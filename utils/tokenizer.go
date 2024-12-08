package utils

import (
	log "github.com/sirupsen/logrus"
	"github.com/tiktoken-go/tokenizer"
)

func Tokenize(str string) int {
	enc, err := tokenizer.Get(tokenizer.O200kBase)
	if err != nil {
		log.Error("Error getting tokenizer:", err)
		return 0
	}
	ids, _, err := enc.Encode(str)
	if err != nil {
		log.Error("Error encoding string:", err)
		return 0
	}
	return len(ids)
}
