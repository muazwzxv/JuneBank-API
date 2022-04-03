package handler

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

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
			expectedCode: 200,
			payload: []byte(
				`{
					"owner": "Muaz Bin Wazir",
					"balance": 0.00,
					"currency": US Dollar
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

		assert.Equal(t, test.expectedCode, resp.StatusCode, test.description)

		data, _ := ioutil.ReadAll(resp.Body)
		t.Fatalf("%s", data)
	}
}
