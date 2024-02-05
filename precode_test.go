package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	reqOk                 = "/cafe?count=3&city=moscow"
	reqWrongCity          = "/cafe?count=3&city=wrongcity"
	reqCountMoreThanTotal = "/cafe?count=100&city=moscow"

	bodyWrongCity = "wrong city value"

	cafeListDelim = ","
)

// TestMainHandlerWhenOk проверяет, что запрос сформирован корректно,
// сервис возвращает код ответа 200 и тело ответа не пустое.
func TestMainHandlerWhenOk(t *testing.T) {
	request := httptest.NewRequest("GET", reqOk, nil)
	respRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(respRecorder, request)

	require.Equal(t, http.StatusOK, respRecorder.Code)

	body := respRecorder.Body.String()
	assert.NotEmpty(t, body)
}

// TestMainHandlerWhenWrongCityValue проверяет, что город, который передаётся
// в параметре city, не поддерживается.
// Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
func TestMainHandlerWhenWrongCityValue(t *testing.T) {
	request := httptest.NewRequest("GET", reqWrongCity, nil)
	respRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(respRecorder, request)

	require.Equal(t, http.StatusBadRequest, respRecorder.Code)

	body := respRecorder.Body.String()
	assert.Equal(t, bodyWrongCity, body)
}

// TestMainHandlerWhenCountMoreThanTotal проверяет, что если значение параметра count
// превышает кол-во кафе в городе, то в ответе возвращается список всех доступных кафе.
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	request := httptest.NewRequest("GET", reqCountMoreThanTotal, nil)
	respRecorder := httptest.NewRecorder()

	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(respRecorder, request)

	// здесь нужно добавить необходимые проверки
	require.Equal(t, http.StatusOK, respRecorder.Code)

	body := respRecorder.Body.String()
	list := strings.Split(body, cafeListDelim)
	assert.Equal(t, totalCount, len(list))
}
