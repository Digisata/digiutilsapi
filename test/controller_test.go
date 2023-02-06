package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Digisata/digiutilsapi/app"
	"github.com/Digisata/digiutilsapi/controller"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func setupRouter() http.Handler {
	validate := validator.New()
	controller := controller.NewController(validate)
	router := app.NewRouter(controller)

	return router
}

func TestAddBinarySuccess(t *testing.T) {
	router := setupRouter()
	requestBody := strings.NewReader(`{"a": "11", "b": "1"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/add_binary", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, "100", responseBody["data"])
}

func TestAddBinaryFailed(t *testing.T) {
	router := setupRouter()
	requestBody := strings.NewReader(`{"a": "11", "b": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/add_binary", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestIsPalindromeSuccess(t *testing.T) {
	router := setupRouter()
	requestBody := strings.NewReader(`{"a": "A man, a plan, a canal: Panama"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/is_palindrome", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, true, responseBody["data"])
}

func TestIsPalindromeFailed(t *testing.T) {
	router := setupRouter()
	requestBody := strings.NewReader(`{"a": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/is_palindrome", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestRomanToIntSuccess(t *testing.T) {
	router := setupRouter()
	requestBody := strings.NewReader(`{"a": "III"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/roman_to_int", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, 3, int(responseBody["data"].(float64)))
}

func TestRomanToIntFailed(t *testing.T) {
	router := setupRouter()
	requestBody := strings.NewReader(`{"a": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/roman_to_int", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestContainsDuplicateSuccess(t *testing.T) {
	router := setupRouter()
	requestBody := strings.NewReader(`{"a": [1, 2, 3, 1]}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/contains_duplicate", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, true, responseBody["data"])
}

func TestContainsDuplicateFailed(t *testing.T) {
	router := setupRouter()
	requestBody := strings.NewReader(`{}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/contains_duplicate", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestIsPowerOfTwoSuccess(t *testing.T) {
	router := setupRouter()
	requestBody := strings.NewReader(`{"a": 1}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/is_power_of_two", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "Ok", responseBody["status"])
	assert.Equal(t, true, responseBody["data"])
}

func TestIsPowerOfTwoFailed(t *testing.T) {
	router := setupRouter()
	requestBody := strings.NewReader(`{}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/is_power_of_two", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(recorder.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}
