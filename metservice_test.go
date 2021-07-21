package metservice

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// setup sets up a test HTTP server along with a Client that is configured to
// talk to that test server. Tests should register handlers on mux which
// provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(mux)

	// client is the Client being tested and is configured to use test server.
	localURL, _ := url.Parse(server.URL + "/")
	localClient := &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    localURL.String(),
	}

	return localClient, mux, server.Close
}

// Test whether the marshaling of v produces JSON that corresponds to the want
// string.
func testJSONMarshal(t *testing.T, v interface{}, want string) {
	// Unmarshal the wanted JSON, to verify its correctness, and marshal it
	// back to sort the keys.
	u := reflect.New(reflect.TypeOf(v)).Interface()
	if err := json.Unmarshal([]byte(want), &u); err != nil {
		t.Errorf("Unable to unmarshal JSON for %v: %v", want, err)
	}
	w, err := json.Marshal(u)
	if err != nil {
		t.Errorf("Unable to marshal JSON for %#v", u)
	}

	// Marshal the target value.
	j, err := json.Marshal(v)
	if err != nil {
		t.Errorf("Unable to marshal JSON for %#v", v)
	}

	if diff := cmp.Diff(string(w), string(j)); diff != "" {
		t.Errorf("json.Marshal(%q):\n%s", v, diff)
	}
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}
