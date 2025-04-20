package belajar_golang_logging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SampleHook struct {
}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Sample Hooks", entry.Level, entry.Message) // -> akan dipanggil ketika ada logger yang levelnya logrus.ErrorLevel, logrus.WarnLevel
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(&SampleHook{})
	logger.SetLevel(logrus.TraceLevel)

	logger.Info("Hello Info")
	logger.Warn("Hello Warn")
	logger.Error("Hello Error")
	logger.Debug("Hello Debug")
}
