package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	LoadConfig()
}

func TestConfig(t *testing.T) {
	assert.NotEmpty(t, App.Port)
}
