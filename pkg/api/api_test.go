package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	postResponse string = "Updated list of websites: www.google.com,www.facebook.com,www.fakewebsite1.com,thisisnotaurl"
)

func TestPostWebsites(t *testing.T) {

	requestBody, err := json.Marshal(postRequestBody{Websites: []string{"www.google.com", "www.facebook.com", "www.fakewebsite1.com", "thisisnotaurl"}})
	if err != nil {
		t.Errorf(err.Error())
	}
	request := httptest.NewRequest(http.MethodPost, "/websites", bytes.NewReader(requestBody))
	response := httptest.NewRecorder()
	PostWebsites(response, request)
	res := response.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != postResponse {
		t.Errorf("expected `%s` got %v", postResponse, string(data))
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code `%d` got %d", http.StatusOK, res.StatusCode)
	}
}

func TestGetWebsites(t *testing.T) {

	var tests = []struct {
		queryParam     string
		wantedResponse string
	}{
		{"", "{\"thisisnotaurl\":\"INVALID_URL\",\"www.facebook.com\":\"INIT\",\"www.fakewebsite1.com\":\"INIT\",\"www.google.com\":\"INIT\"}"},
		{"?name=www.google.com", "{\"www.google.com\":\"INIT\"}"},
		{"?name=www.facebook.com,www.fakewebsite1.com", "{\"www.facebook.com\":\"INIT\",\"www.fakewebsite1.com\":\"INIT\"}"},
		{"?name=www.youtube.com", "{\"www.youtube.com\":\"DOES_NOT_EXIST\"}"},
	}

	for i, tt := range tests {
		testName := fmt.Sprintf("Test %d", i)
		t.Run(testName, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, "/websites"+tt.queryParam, nil)
			response := httptest.NewRecorder()
			GetWebsites(response, request)
			res := response.Result()
			defer res.Body.Close()
			data, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			if string(data) != tt.wantedResponse {
				t.Errorf("expected `%s` got %v", tt.wantedResponse, string(data))
			}
			if res.StatusCode != http.StatusOK {
				t.Errorf("expected status code `%d` got %d", http.StatusOK, res.StatusCode)
			}
		})
	}
}
