//go:build integration

package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest_Integration(t *testing.T) {
	assert.Equal(t, true, NewRequest().KeepAlive("https://www.google.com.br/"))
}
