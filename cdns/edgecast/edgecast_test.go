package edgecast

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPurge(t *testing.T) {
	token := "00000000-0000-0000-0000-000000000000"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "TOK:"+token {
			t.Error("Authorization header is incorrect.")
		}

		if r.URL.Path != "/v2/mcc/customers/0000/edge/purge" {
			t.Errorf("URL was not properly constructed: %s", r.URL.Path)
		}

		fmt.Fprint(w, `{"Id": "success"}`)
	}))
	defer ts.Close()

	api := NewAPI("0000", token)
	api.BaseURL = ts.URL + "/v2"
	id, err := api.Purge("https://testurl/testPurge")
	if err != nil {
		t.Fatal(err)
	}
	if id != "success" {
		t.Error("purge did not return success")
	}
}
