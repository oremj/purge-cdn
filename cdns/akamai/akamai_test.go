package akamai

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPurge(t *testing.T) {
	api := NewAPI("testuser", "testpass")

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok {
			t.Error("Basic auth was not set.")
		}
		if user != "testuser" {
			t.Errorf("user is not testuser: %s", user)
		}
		if pass != "testpass" {
			t.Errorf("pass is not testpass: %s", pass)
		}

		if r.URL.Path != "/ccu/v2/queues/default" {
			t.Errorf("URL was not properly constructed: %s", r.URL.Path)
		}

		w.WriteHeader(201)

		fmt.Fprint(w, `{
			"httpStatus" : 201,
			"detail" : "Request accepted.",
			"estimatedSeconds" : 420,
			"purgeId" : "95b5a092-043f-4af0-843f-aaf0043faaf0",
			"progressUri" : "/ccu/v2/purges/95b5a092-043f-4af0-843f-aaf0043faaf0",
			"pingAfterSeconds" : 420,
			"supportId" : "17PY1321286429616716-211907680"
		}`)
	}))
	defer ts.Close()

	api.BaseURL = ts.URL

	resp, err := api.Purge("https://testurl/testPurge")
	if err != nil {
		t.Fatal(err)
	}

	if resp.PurgeID != "95b5a092-043f-4af0-843f-aaf0043faaf0" {
		t.Error("purgeId was not correct", resp.PurgeID)
	}
}
