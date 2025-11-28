package moduleMgr

import "log"

type ILogger interface {
	Info(args ...any)
	Error(args ...any)
	Warn(args ...any)
}

type FmtLogger struct{}

func (s *FmtLogger) Info(args ...any) {
	args = append([]any{"[INFO]"}, args...)
	log.Println(args...)
}

func (s *FmtLogger) Error(args ...any) {
	args = append([]any{"[ERROR]"}, args...)
	log.Println(args...)
}

func (s *FmtLogger) Warn(args ...any) {
	args = append([]any{"[WARN]"}, args...)
	log.Println(args...)
}
