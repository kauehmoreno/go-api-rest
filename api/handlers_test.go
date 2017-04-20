package api

import (
	"gopkg.in/check.v1"
	"testing"
)

func TestHandlers(t *testing.T) {
	check.TestingT(t)
}

type MyHandlerSuite struct {
	handler http.Handler
}
