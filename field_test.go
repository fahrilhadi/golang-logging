package golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T)  {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "abu").Info("Hello World")

	logger.WithField("username", "fahrilhadi").
		WithField("name", "Fahril Hadi").
		Info("Hello World")
}

func TestFields(t *testing.T)  {
	logger := logrus.New() 
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "fahril",
		"name": "Fahril",
	}).Info("Hello World")
}