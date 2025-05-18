package oauth

type Token string

type Request struct {
	Provider    Provider
	AccessToken Token
}
