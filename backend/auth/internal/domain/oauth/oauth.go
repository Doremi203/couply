package oauth

type Token string

type Code string

type CodeVerifier string

type DeviceID string

type State string

type Request struct {
	Provider     ProviderType
	Code         Code
	State        State
	CodeVerifier CodeVerifier
	DeviceID     DeviceID
}
