package logger

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
)

// Log is a global logger object
var Log log.FieldLogger // Global logger object

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
	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	fmt.Printf("Level a: Loglevel is: %v, LOG_LEVEL specifies: %v, ERR is: %v\n", log.GetLevel(), level, err)
	if err != nil {
		log.SetLevel(log.ErrorLevel)
		fmt.Printf("Level b: Due to error Loglevel is set to default Err loglevel: %v\n", log.ErrorLevel)
	} else {
		log.SetLevel(level)
		fmt.Printf("Level b: We set Loglevel to LOG_LEVEL: %v\n", level)
	}
	fmt.Printf("LEVEL is %v\n", log.GetLevel())

	Log = log.WithFields(log.Fields{
		"@version": "1",
		"logger":   "kpi-uploader",
	})
}
