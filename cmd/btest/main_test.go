package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	statusCode, body := getResult(nil, req)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "OK", body)
}

func TestHelloHandler(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	statusCode, body := getResult(nil, req)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "Hello", body)
}

func TestStopHandler(t *testing.T) {
	var statusCode int
	var body string

	srv := NewServer()

	req, _ := http.NewRequest(http.MethodGet, "/stop", nil)
	statusCode, body = getResult(srv, req)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "Stopping", body)

	req, _ = http.NewRequest(http.MethodGet, "/ping", nil)
	statusCode, body = getResult(srv, req)
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.Equal(t, "Stopping", body)

	req, _ = http.NewRequest(http.MethodGet, "/hello", nil)
	statusCode, body = getResult(srv, req)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "Hello", body)

	req, _ = http.NewRequest(http.MethodGet, "/restart", nil)
	statusCode, body = getResult(srv, req)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "Restarting", body)

	req, _ = http.NewRequest(http.MethodGet, "/ping", nil)
	statusCode, body = getResult(srv, req)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "OK", body)
}

func getResult(srv *Server, req *http.Request) (int, string) {
	if srv == nil {
		srv = NewServer()
	}
	responseRecorder := httptest.NewRecorder()
	srv.ServeHTTP(responseRecorder, req)
	result := responseRecorder.Result()

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(result.Body)

	var res, err = io.ReadAll(responseRecorder.Body)
	fmt.Println(string(res))
	if err != nil {
		log.Fatal(err)
	}
	return result.StatusCode, string(res)
}
