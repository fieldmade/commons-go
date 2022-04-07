package idx

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func randomInt(max int) int {
	if max <= 0 {
		max = 1
	}

	binN, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(binN.Int64())
}

// Removing 0, O, I, L since it's hard to know the difference.
var codeChars = []rune{
	'1', '2', '3', '4', '5', '6', '7', '8', '9',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'J',
	'K', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'U',
	'V', 'W', 'X', 'Y', 'Z',
}

func generateCodeSegment(num int) string {
	var values = make([]rune, num)

	for i := 0; i < num; i++ {
		values[i] = codeChars[randomInt(len(codeChars))]
	}

	return string(values)
}

func GenerateCode(numChars ...int) string {
	var segments []string

	for _, num := range numChars {
		segments = append(segments, generateCodeSegment(num))
	}

	return strings.Join(segments, "-")
}
