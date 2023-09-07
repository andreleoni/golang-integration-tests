//go:build !integration

package request

import (
	"build-tags-integration-tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	requestInterfaceMock := &mocks.RequestInterface{}
	requestInterfaceMock.On("KeepAlive", "https://www.google.com.br").Return(true)

	assert.Equal(t, requestInterfaceMock.KeepAlive("https://www.google.com.br"), true)
}
