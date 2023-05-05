package utils

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRequest(t *testing.T) {
	response, err := Request(http.MethodGet, "https://www.google.com/ncr")
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response)
}

func TestRequestWithSocks5(t *testing.T) {
	response, err := RequestWithSocks5(http.MethodGet, "https://www.google.com/ncr", "127.0.0.1:7890")
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.NotEmpty(t, response)
}
