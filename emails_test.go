package resend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mux    *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	client = NewClient("")
	url, _ := url.Parse(server.URL)
	client.BaseURL = url
}

func teardown() {
	server.Close()
}

func TestSendEmail(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/emails", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		ret := &SendEmailResponse{
			Id: "1923781293",
		}
		err := json.NewEncoder(w).Encode(&ret)
		if err != nil {
			panic(err)
		}
	})

	req := &SendEmailRequest{
		To: []string{"d@e.com"},
	}
	resp, err := client.Emails.Send(req)
	fmt.Println("RESPONSE")
	fmt.Printf("%v", resp)
	if err != nil {
		t.Errorf("Emails.Send returned error: %v", err)
	}
	assert.Equal(t, resp.Id, "1923781293")
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if expected != r.Method {
		t.Errorf("Request method = %v, expected %v", r.Method, expected)
	}
}
