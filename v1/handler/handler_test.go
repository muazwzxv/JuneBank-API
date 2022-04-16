package handler

import (
	"bytes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetAccountById(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		method       string
		expectedCode int
		id           int
	}{
		{
			description:  "Get All Accounts status 200",
			route:        fmt.Sprintf("http://localhost:8080/api/account/%d", 3),
			method:       "GET",
			expectedCode: fiber.StatusOK,
			id:           3,
		},
	}

	for _, test := range tests {
		req, err := http.NewRequest(test.method, test.route, nil)
		if err != nil {
			t.Fatal(err)
		}

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		data, _ := ioutil.ReadAll(resp.Body)
		t.Logf("%s", data)

		assert.Equal(t, test.expectedCode, resp.StatusCode)
	}
}

func TestGetAllAccount(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		method       string
		expectedCode int
	}{
		{
			description:  "Get All Accounts status 200",
			route:        "http://localhost:8080/api/account",
			method:       "GET",
			expectedCode: fiber.StatusOK,
		},
	}

	for _, test := range tests {

		req, err := http.NewRequest(test.method, test.route, nil)
		if err != nil {
			t.Fatal(err)
		}

		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		data, _ := ioutil.ReadAll(resp.Body)
		t.Logf("%s", data)
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}

func TestPostAccount(t *testing.T) {
	tests := []struct {
		description  string
		route        string
		method       string
		expectedCode int
		payload      []byte
	}{
		{
			description:  "Post User status 200",
			route:        "http://localhost:8080/api/account",
			method:       "POST",
			expectedCode: fiber.StatusCreated,
			payload: []byte(
				`{
					"owner": "Fatin Athirah Binti Nordin",
					"balance": 0,
					"currency": "US Dollar"
				}`,
			),
		},
	}

	for _, test := range tests {

		req, err := http.NewRequest(test.method, test.route, bytes.NewBuffer(test.payload))
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}

		data, _ := ioutil.ReadAll(resp.Body)
		t.Logf("%s", data)
		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
