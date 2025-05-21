package token

import "time"

type Pair struct {
	AccessToken  Token
	RefreshToken Refresh
}

type Token struct {
	signedString string
	expiresIn    time.Duration
}

func (t Token) SignedString() string {
	return t.signedString
}

func (t Token) ExpiresIn() time.Duration {
	return t.expiresIn
}
