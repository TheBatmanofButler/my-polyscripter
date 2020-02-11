package token

import (
	"math/rand"
	"strings"
	"time"
)

var CHARS = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"abcdefghijklmnopqrstuvwxyz")

func getRandomString() string {
	rand.Seed(time.Now().UnixNano())
	length := rand.Intn(3) + 5
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(CHARS[rand.Intn(len(CHARS))])
	}
	str := b.String()

	return str

}

func getRandomKeywords() map[string]TokenType {

	var keywords = map[string]TokenType{
		getRandomString(): FUNCTION,
		getRandomString(): LET,
		getRandomString(): TRUE,
		getRandomString(): FALSE,
		getRandomString(): IF,
		getRandomString(): ELSE,
		getRandomString(): RETURN,
	}

	return keywords
}
