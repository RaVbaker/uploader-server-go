package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getResponseBody(ts *httptest.Server, t *testing.T) (string, error) {
	resp, err := http.Get(ts.URL) // test from the outside
	if err != nil {
		t.Error(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	return string(body), err
}

func TestPing(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(statusHandler))
	defer ts.Close()

	if body, _ := getResponseBody(ts, t); body != "ok" {
		t.Error("expected", "ok", "got", body)
	}
}
