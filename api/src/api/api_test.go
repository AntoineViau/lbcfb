package api

import (
	"net/http"
	"os"
	"testing"

	"github.com/gavv/httpexpect"
)

func getDomain() string {
	if os.Getenv("INTEGRATION_BASE_URL") == "" {
		return "http://localhost:8080"
	}
	return os.Getenv("INTEGRATION_BASE_URL")
}

func createRequest(t *testing.T) *httpexpect.Request {
	return httpexpect.New(t, getDomain()).GET("/api/v1/fizzbuzz")
}

func TestApiValid(t *testing.T) {
	createRequest(t).
		WithQuery("string1", "fizz").
		WithQuery("string2", "buzz").
		WithQuery("int1", "3").
		WithQuery("int2", "5").
		WithQuery("limit", "5").
		Expect().
		Status(http.StatusOK).
		JSON().Array().Elements("1", "2", "fizz", "4", "buzz")
}

func TestApiInvalidQueryParamName(t *testing.T) {
	createRequest(t).
		WithQuery("string", "fizz").
		Expect().
		Status(http.StatusBadRequest)
}

func TestApiMissingParam(t *testing.T) {
	createRequest(t).
		WithQuery("string2", "buzz").
		WithQuery("int1", "3").
		WithQuery("int2", "5").
		WithQuery("limit", "5").
		Expect().
		Status(http.StatusBadRequest)
}

func TestApiZeroInt1(t *testing.T) {
	createRequest(t).
		WithQuery("string1", "fizz").
		WithQuery("string2", "buzz").
		WithQuery("int1", "0").
		WithQuery("int2", "5").
		WithQuery("limit", "5").
		Expect().
		Status(http.StatusBadRequest)
}

func TestApiZeroInt2(t *testing.T) {
	createRequest(t).
		WithQuery("string1", "fizz").
		WithQuery("string2", "buzz").
		WithQuery("int1", "3").
		WithQuery("int2", "0").
		WithQuery("limit", "5").
		Expect().
		Status(http.StatusBadRequest)
}

func TestApiInvalidLimit(t *testing.T) {
	createRequest(t).
		WithQuery("string1", "fizz").
		WithQuery("string2", "buzz").
		WithQuery("int1", "3").
		WithQuery("int2", "0").
		WithQuery("limit", "-5").
		Expect().
		Status(http.StatusBadRequest)
}
