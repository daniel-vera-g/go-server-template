package utils

import (
	"reflect"
	"testing"
)

var msg = map[string]interface{}{
	"message": "Test",
	"status":  true,
}

func TestMessage(t *testing.T) {
	res := Message(true, "Test")
	eq := reflect.DeepEqual(msg, res)
	if !eq {
		t.Error("The return value should be a JSON message.")
		t.Error(msg)
	}
}
