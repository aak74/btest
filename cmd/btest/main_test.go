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
	statusCode, body := getResultFromNewServer(req)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "OK", body)
}

func TestHelloHandler(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	statusCode, body := getResultFromNewServer(req)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "Hello", body)
}

func TestStopHandler(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/stop", nil)
	responseRecorder := httptest.NewRecorder()
	srv := NewServer()
	srv.ServeHTTP(responseRecorder, req)

	// Test first request. Expected
	statusCode, body := getResult(responseRecorder)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "Stopping", body)

	req, _ = http.NewRequest(http.MethodGet, "/ping", nil)
	responseRecorder = httptest.NewRecorder()
	srv.ServeHTTP(responseRecorder, req)
	statusCode, body = getResult(responseRecorder)
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.Equal(t, "Stopping", body)

	req, _ = http.NewRequest(http.MethodGet, "/hello", nil)
	responseRecorder = httptest.NewRecorder()
	srv.ServeHTTP(responseRecorder, req)
	statusCode, body = getResult(responseRecorder)
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "Hello", body)
}

func getResultFromNewServer(req *http.Request) (int, string) {
	responseRecorder := httptest.NewRecorder()
	srv := NewServer()
	srv.ServeHTTP(responseRecorder, req)
	return getResult(responseRecorder)
}

func getResult(responseRecorder *httptest.ResponseRecorder) (int, string) {
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
