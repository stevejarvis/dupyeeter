package config

import (
	"testing"
)

func TestConfigParse(t *testing.T) {
	_, err := NewConfig()
	if err == nil {
		t.Error("Expected config without required params to error.")
	}
}
