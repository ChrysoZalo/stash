package server

import (
	"log/slog"
	"time"

	"github.com/go-chi/httplog/v2"
)

// Logger
var logger = httplog.NewLogger("httplog-example", httplog.Options{
	// JSON:             true,
	LogLevel:         slog.LevelDebug,
	Concise:          true,
	RequestHeaders:   true,
	MessageFieldName: "message",
	TimeFieldFormat:  time.RFC850,
	// SourceFieldName: "source",
})
