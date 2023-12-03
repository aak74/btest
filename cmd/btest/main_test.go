package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	response := httptest.NewRecorder()

	srv := NewServer()
	srv.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Result().StatusCode)
}

func TestHelloHandler(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	response := httptest.NewRecorder()

	srv := NewServer()
	srv.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Result().StatusCode)
}
