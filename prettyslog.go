package prettyslog

import (
	"context"
	"fmt"
	"io"
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
	writer io.Writer
	level  slog.Level
	slog.Handler
}

func NewColorHandler(w io.Writer, opts *slog.HandlerOptions) *ColorHandler {
	if opts == nil {
		opts = &slog.HandlerOptions{Level: slog.LevelInfo}
	}

	return &ColorHandler{writer: w, level: opts.Level.Level()}
}

func (h *ColorHandler) Handle(ctx context.Context, r slog.Record) error {
	if r.Level < h.level {
		return nil
	}

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

	fmt.Fprintf(h.writer, "%s%s: %s%s\n", color, r.Level.String(), r.Message, Reset)

	return nil
}

func (h *ColorHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *ColorHandler) WithGroup(name string) slog.Handler {
	return h
}
