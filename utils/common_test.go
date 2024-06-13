package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandString(t *testing.T) {
	str := GenerateRandString(8)
	assert.NotEmpty(t, str)
}
