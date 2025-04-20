package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "Jhon Doe").Info("Hello logging")

	logger.WithField("username", "rossi").WithField("name", "Valentino rossi").Info("Hello logging")
}

func TestFieldS(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "doe",
		"name":     "Jhon Doe",
	}).Info("Hello logging")
}
