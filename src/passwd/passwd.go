package passwd

import (
	"math/rand"
	"time"
	"strings"
)

var randomizer *rand.Rand

func init() {
	randomizer = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func Generate(length int, uppercase bool, numbers bool, specials bool) string {
	var strBuilder strings.Builder

	strBuilder.WriteString("abcdefghijklmnopqrstuvwxyz")
	if uppercase {
		strBuilder.WriteString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	if numbers {
		strBuilder.WriteString("0123456789")
	}
	if specials {
		strBuilder.WriteString("!?[],.()-_=")
	}

	charset := strBuilder.String()

	bytes := make([]byte, length)
	for i := range bytes {
		bytes[i] = charset[randomizer.Intn(len(charset))]
	}
	return string(bytes)
}
