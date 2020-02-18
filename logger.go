package logger

import (
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// Logger is a global logger object
var Logger log.FieldLogger // Global logger object

func init() {
	if os.Getenv("LOG_FORMAT") == "json" { // Fluentd field name conventions

		log.SetFormatter(&log.JSONFormatter{
			TimestampFormat: time.RFC3339Nano,
			FieldMap: log.FieldMap{
				log.FieldKeyTime:  "@timestamp",
				log.FieldKeyLevel: "level",
				log.FieldKeyMsg:   "message",
				log.FieldKeyFunc:  "caller",
			},
		})
		log.SetReportCaller(true)

	} else if os.Getenv("LOG_FORMAT") == "json_plain" { // Logrus field names
		log.SetFormatter(&log.JSONFormatter{})
	}
	if os.Getenv("LOG_STDOUT") == "true" {
		log.SetOutput(os.Stdout)
	}

	// Logrus has seven log levels:
	// Trace, Debug, Info, Warning, Error, Fatal and Panic.
	level, _ := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	log.SetLevel(level)

	Logger = log.WithFields(log.Fields{
		"@version": "1",
		"logger":   "kpi-uploader",
	})
}