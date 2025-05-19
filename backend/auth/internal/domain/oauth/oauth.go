package oauth

type Token string

type Code string

type State string

type Request struct {
	Provider    ProviderType
	Code        Code
	State       State
	AccessToken Token
}
