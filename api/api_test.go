package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pkbhowmick/go-rest-api/auth"

	"github.com/magiconair/properties/assert"
)

type Test struct {
	Method             string
	Url                string
	Body               io.Reader
	ExpectedStatusCode int
}

type TestWithID struct {
	Method             string
	Url                string
	Body               io.Reader
	ExpectedStatusCode int
	UserID             string
}

func TestCreateUser(t *testing.T) {
	Init()
	tests := []Test{
		{
			"POST",
			"/api/users",
			bytes.NewReader([]byte(`{
    "id": "6",
    "firstName": "test",
    "lastName": "test",
    "repositories": [
        {
            "id": "10006",
            "name": "http-server",
            "visibility": "public",
            "star": 10
        }
    ]
}`)),
			201,
		},
		{
			"POST",
			"/api/users",
			bytes.NewReader([]byte(`{
    "id": "7",
    "firstName": "test",
    "lastName": "test",
    "repositories": []
}`)),
			201,
		},
		{
			"POST",
			"/api/users",
			bytes.NewReader([]byte(`{"id": "10"}`)),
			400,
		},
	}
	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Url, test.Body)
		if err != nil {
			t.Fatal(err)
		}
		token, err := auth.GenerateToken("test")
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, res.Result().StatusCode, test.ExpectedStatusCode)
	}
}

func TestDeleteUser(t *testing.T) {
	Init()
	tests := []TestWithID{
		{"DELETE", "/api/users/%s", nil, 200, "1"},
		{"DELETE", "/api/users/%s", nil, 404, "1"},
		{"DELETE", "/api/users/%s", nil, 404, "6"},
	}
	for _, test := range tests {
		url := fmt.Sprintf(test.Url, test.UserID)
		req, err := http.NewRequest(test.Method, url, test.Body)
		if err != nil {
			t.Fatal(err)
		}
		token, err := auth.GenerateToken("test")
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, res.Result().StatusCode, test.ExpectedStatusCode)
	}
}

func TestGetUser(t *testing.T) {
	Init()
	tests := []TestWithID{
		{"GET", "/api/users/%s", nil, 200, "1"},
		{"GET", "/api/users/%s", nil, 404, "8"},
		{"GET", "/api/users/%s", nil, 200, "5"},
	}
	for _, test := range tests {
		url := fmt.Sprintf(test.Url, test.UserID)
		req, err := http.NewRequest(test.Method, url, test.Body)
		if err != nil {
			t.Fatal(err)
		}
		token, err := auth.GenerateToken("test")
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, res.Result().StatusCode, test.ExpectedStatusCode)
	}
}

func TestGetUsers(t *testing.T) {
	Init()
	tests := []Test{
		{"GET", "/api/users", nil, 200},
	}
	for _, test := range tests {
		req, err := http.NewRequest(test.Method, test.Url, test.Body)
		if err != nil {
			t.Fatal(err)
		}
		token, err := auth.GenerateToken("test")
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, res.Result().StatusCode, test.ExpectedStatusCode)
	}
}

func TestUpdateUser(t *testing.T) {
	Init()
	tests := []TestWithID{
		{
			"PUT",
			"/api/users/%s",
			bytes.NewReader([]byte(`{"firstName": "test","lastName": "test"}`)),
			200,
			"1",
		},
		{
			"PUT",
			"/api/users/%s",
			bytes.NewReader([]byte(`{"firstName": "test",
    "lastName": "test",
    "repositories": [
        {
            "id": "10006",
            "name": "http-server",
            "visibility": "public",
            "star": 10
        }
    ]}`)),
			200,
			"5",
		},
		{
			"PUT",
			"/api/users/%s",
			bytes.NewReader([]byte(``)),
			400,
			"2",
		},
	}

	for _, test := range tests {
		url := fmt.Sprintf(test.Url, test.UserID)
		req, err := http.NewRequest(test.Method, url, test.Body)
		if err != nil {
			t.Fatal(err)
		}
		token, err := auth.GenerateToken("test")
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Authorization", "Bearer "+token)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)
		assert.Equal(t, res.Result().StatusCode, test.ExpectedStatusCode)
	}
}
