package logutils

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// InitLogger configures the logger with a standardized format and log level.
func InitLogger() {

	// Log level: can be modified based on the environment (debug, info, etc.)
	LogLevel := os.Getenv("LOG_LEVEL")
	if LogLevel != "" {
		typeLevelZeroLog, _ := zerolog.ParseLevel(LogLevel)
		zerolog.SetGlobalLevel(typeLevelZeroLog)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	// Set Zerolog's output format to JSON
	zerolog.TimeFieldFormat = time.RFC3339

	// Standard output (stdout) to be captured
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
}

// Info logs an information message.
func Info(message string, fields map[string]interface{}) {
	event := log.Info().Timestamp()
	for k, v := range fields {
		event = event.Str(k, v.(string))
	}
	event.Msg(message)
}

// Error logs an error message.
func Error(message string, err error, fields map[string]interface{}) {
	event := log.Error().Err(err).Timestamp()
	for k, v := range fields {
		event = event.Str(k, v.(string))
	}
	event.Msg(message)
}

// Debug logs a debug message.
func Debug(message string, fields map[string]interface{}) {
	event := log.Debug().Timestamp()
	for k, v := range fields {
		event = event.Str(k, v.(string))
	}
	event.Msg(message)
}

// Warn logs a warning message.
func Warn(message string, fields map[string]interface{}) {
	event := log.Warn().Timestamp()
	for k, v := range fields {
		event = event.Str(k, v.(string))
	}
	event.Msg(message)
}

// Fatal logs a critical error message and terminates the program.
func Fatal(message string, err error, fields map[string]interface{}) {
	event := log.Fatal().Err(err).Timestamp()
	for k, v := range fields {
		event = event.Str(k, v.(string))
	}
	event.Msg(message)
	os.Exit(1)
}
