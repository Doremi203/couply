package telegram

type CallbackResult struct {
	ActionText   string
	ResponseText string
	Error        error
}

type CallbackHandler func(userID string) CallbackResult
