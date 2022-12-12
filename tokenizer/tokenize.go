package tokenizer

type TokenKind int

const (
	NumberToken TokenKind = iota
	IdentifierToken
	PlusToken
	MinusToken
	StarToken
	SlashToken
	PercentToken
	ExponentiationToken
	OpenParenthesisToken
	CloseParenthesisToken
	CommaToken
)

type Token struct {
	Kind     TokenKind
	Value    string
	Position int
}

func Tokenize(code string) ([]Token, error) {
	return nil, nil
}
