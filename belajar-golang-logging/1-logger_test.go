package belajar_golang_logging

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func Test(t *testing.T) {
	logger := logrus.New()

	logger.Println("Hello Logger")
	fmt.Println("Hello Logger")
}
