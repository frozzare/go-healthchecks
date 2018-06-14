package healthchecks

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthchecks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`ok`))
	}))

	defer server.Close()

	c := NewClient(nil, server.URL)

	if err := c.Success(context.Background(), "test-id"); err != nil {
		t.Fatal(err)
	}

	if err := c.Fail(context.Background(), "test-id"); err != nil {
		t.Fatal(err)
	}
}
