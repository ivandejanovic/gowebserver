package service

import (
	"testing"
)

func TestXYZ(t *testing.T) {
	value, _ := Method("test")
	if value != "Your key is: test" {
		t.Error("Expected returned string to be 'Your key is: test'")
	}
}
