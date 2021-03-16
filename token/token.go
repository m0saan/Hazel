package token

type tokenType string

type Token struct {
	Type tokenType
	Literal string
}
