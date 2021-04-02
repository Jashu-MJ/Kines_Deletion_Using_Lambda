package handler_test

import (
	"awslambdamonitor/handler"
	"testing"
)

func TestHandler(t *testing.T) {
	err := handler.Handler()
	if err != nil {
		t.Error("error occured at TestHandler", err)
	}
}
