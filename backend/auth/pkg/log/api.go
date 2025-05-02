package log

type Logger interface {
	Info(msg string, args ...any)
	Warn(error)
	Error(error)
}
