package log

type Logger interface {
	Infof(format string, args ...any)
	Warn(error)
	Error(error)
}
