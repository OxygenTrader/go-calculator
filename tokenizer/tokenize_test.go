package tokenizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenize(test *testing.T) {
	type data struct {
		name           string
		code           string
		expectedTokens []Token
		expectedErr    string
	}

	tests := []data{}
	for _, testData := range tests {
		actualTokens, actualErr := Tokenize(testData.code)

		assert.Equal(test, testData.expectedTokens, actualTokens, testData.name)
		if testData.expectedErr == "" {
			assert.NoError(test, actualErr, testData.name)
		} else {
			assert.EqualError(test, actualErr, testData.expectedErr, testData.name)
		}
	}
}
