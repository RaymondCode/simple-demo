package util

import (
	"github.com/google/uuid"
	"github.com/ser163/WordBot/generate"
)

func CreateUuid() int64 {
	return int64(uuid.New().ID())
}

func CreateRandomString(count int) []string {
	var randStrs []string
	for i := 1; i <= count; i++ {
		wordList, _ := generate.GenRandomMix(10)
		randStrs = append(randStrs, wordList.Word)
	}
	return randStrs
}
