package logutils

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Fields map[string]interface{}

// InitLogger configures the logger with a standardized format and log level.
// With().Caller().Timestamp().Logger() is used to include the caller and timestamp in the log output.
// The log level can be set using the LOG_LEVEL environment variable.
// The log output can be redirected to a file by setting the LOG_PATH environment variable.
// If the LOG_PATH environment variable is not set, the log output will be directed to stdout.
// The default log level is DEBUG.
func InitLogger() zerolog.Logger {

	// Set Zerolog's output format to JSON
	zerolog.TimeFieldFormat = time.RFC3339

	var log zerolog.Logger

	// Log file path: can be modified based on the environment
	logFilePath := os.Getenv("LOG_PATH")
	if logFilePath != "" {
		// Open the log file
		f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to open log file")
		}

		// Set the log file as the output
		log = zerolog.New(f)

	} else {
		// Standard output (stdout) to be captured by the container logging driver
		// Log to stdout
		log = zerolog.New(os.Stdout).With().Caller().Logger()
	}

	// Log level: can be modified based on the environment (debug, info, etc.)
	LogLevel := os.Getenv("LOG_LEVEL")
	if LogLevel != "" {
		typeLevelZeroLog, _ := zerolog.ParseLevel(LogLevel)
		zerolog.SetGlobalLevel(typeLevelZeroLog)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	return log
}

// Info logs an information message.
func Info(message string, fields map[string]interface{}) {
	event := log.Info().Timestamp()
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(message)
}

// Error logs an error message.
func Error(message string, err error, fields map[string]interface{}) {
	event := log.Error().Err(err).Timestamp()
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(message)
}

// Debug logs a debug message.
func Debug(message string, fields map[string]interface{}) {
	event := log.Debug().Timestamp()
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(message)
}

// Warn logs a warning message.
func Warn(message string, fields map[string]interface{}) {
	event := log.Warn().Timestamp()
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(message)
}

// Fatal logs a critical error message and terminates the program.
func Fatal(message string, err error, fields Fields) {
	event := log.Fatal().Err(err).Timestamp()
	for k, v := range fields {
		event = event.Interface(k, v)
	}
	event.Msg(message)
	os.Exit(1)
}
