package initializers

import (
    "os"

    "github.com/sirupsen/logrus"
    log "github.com/sirupsen/logrus"
)

func Logging() {
    // setup logrus
    logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
    if err != nil {
        logLevel = log.InfoLevel
    }

    // Only log based on logLevel
    log.SetLevel(logLevel)

    // Log as JSON instead of the default ASCII formatter.
    log.SetFormatter(&log.JSONFormatter{FieldMap: logrus.FieldMap{
        logrus.FieldKeyTime:  "@timestamp",
        logrus.FieldKeyLevel: "@level",
        logrus.FieldKeyMsg:   "@message",
    }})

    // Output to stdout instead of the default stderr
    // Can be any io.Writer, see below for File example
    log.SetOutput(os.Stdout)
}
