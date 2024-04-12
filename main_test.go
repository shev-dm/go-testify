package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=8&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.Len(t, list, totalCount)
}

func TestMainHandlerWhenCorrectRequest(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()

	require.Equal(t, responseRecorder.Code, http.StatusOK)
	assert.NotEmpty(t, body)
}

func TestMainHandlerWhenNotCorrectCity(t *testing.T) {
	answer := "wrong city value"
	req := httptest.NewRequest("GET", "/cafe?count=4&city=london", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()

	require.Equal(t, responseRecorder.Code, http.StatusBadRequest)
	assert.Equal(t, answer, body)
}
