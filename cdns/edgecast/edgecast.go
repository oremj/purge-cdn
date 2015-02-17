package edgecast

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var APIEndpoint = "https://api.edgecast.com/v2/"

const (
	MediaTypeFlash = 2
	MediaTypeLarge = 3
	MediaTypeSmall = 8
	MediaTypeADN   = 14
)

type PurgeResponse struct {
	Id string
}

type API struct {
	AccountId string
	Token     string
}

func (a *API) addHeaders(req *http.Request) {
	req.Header.Add("Authorization", "TOK:"+a.Token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}

func (a *API) Purge(url string) (string, error) {
	body, err := json.Marshal(map[string]interface{}{
		"MediaPath": url,
		"MediaType": MediaTypeLarge,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("PUT", APIEndpoint+"mcc/customers/"+a.AccountId+"/edge/purge", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	a.addHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Status Code: %d, Body: %s", resp.StatusCode, respBody)
	}

	purgeResp := new(PurgeResponse)
	err = json.Unmarshal(respBody, purgeResp)
	if err != nil {
		return "", fmt.Errorf("json decode error: %s, Body: |%s|", err, respBody)
	}

	return purgeResp.Id, nil
}
