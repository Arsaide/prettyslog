package prettyslog

import (
	"context"
	"fmt"
	"log/slog"
)

const (
	Reset  = "\033[0m"
	Red    = "\\e[0;31m"
	Green  = "\\e[0;32m"
	Yellow = "\\e[0;33m\t"
	Blue   = "\\e[0;34m\t"
)

type ColorHandler struct {
	slog.Handler
}

func NewColorHandler() *ColorHandler {
	return &ColorHandler{}
}

func (h *ColorHandler) Handle(ctx context.Context, r slog.Record) error {
	var color string
	switch r.Level {
	case slog.LevelInfo:
		color = Green
	case slog.LevelWarn:
		color = Yellow
	case slog.LevelError:
		color = Red
	case slog.LevelDebug:
		color = Blue
	default:
		color = Reset
	}

	fmt.Printf("%s%s: %s%s\n", color, r.Level.String(), r.Message, Reset)

	return nil
}

func (h *ColorHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *ColorHandler) WithGroup(name string) slog.Handler {
	return h
}
