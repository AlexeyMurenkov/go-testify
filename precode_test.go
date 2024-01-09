package main

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func request(city string, count int) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(http.MethodGet, "/cafe", http.NoBody)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Add("city", city)
	query.Add("count", strconv.Itoa(count))
	req.URL.RawQuery = query.Encode()

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	return responseRecorder, nil
}
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	city := "moscow"

	responseRecorder, err := request(city, totalCount+1)

	require.NoError(t, err)
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	require.Len(t, strings.Join(cafeList[city], ","), responseRecorder.Body.Len())
}

func TestMainHandlerWhenValidQuery(t *testing.T) {
	count := 1
	city := "moscow"

	responseRecorder, err := request(city, count)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, responseRecorder.Code)
	require.NotEmpty(t, responseRecorder.Body)
}

func TestMainHandlerWhenInvalidCityParam(t *testing.T) {
	count := 1
	city := "non-existent city"

	responseRecorder, err := request(city, count)
	require.NoError(t, err)
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}
