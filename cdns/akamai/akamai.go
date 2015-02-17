package akamai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const APIEndpoint = "https://api.ccu.akamai.com"

type API struct {
	User     string
	Password string

	BaseURL string
}

func NewAPI(user, password string) *API {
	return &API{
		User:     user,
		Password: password,
		BaseURL:  APIEndpoint,
	}
}

type APIPurgeReq struct {
	Objects []string `json:"objects"`
}

type APIPurgeResponse struct {
	Detail           string `json:"detail"`
	EstimatedSeconds int    `json:"estimatedSeconds"`
	HTTPStatus       int    `json:"httpStatus"`
	PingAfterSeconds int    `json:"pingAfterSeconds"`
	ProgressURI      string `json:"progressUri"`
	PurgeID          string `json:"purgeId"`
	SupportID        string `json:"supportId"`
}

func (a *API) addHeaders(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(a.User, a.Password)
}

func (a *API) Purge(url string) (*APIPurgeResponse, error) {
	body, err := json.Marshal(APIPurgeReq{
		Objects: []string{url},
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", a.BaseURL+"/ccu/v2/queues/default", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	a.addHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 201 {
		return nil, fmt.Errorf("Status Code: %d, Body: %s", resp.StatusCode, respBody)
	}

	apiResp := new(APIPurgeResponse)
	err = json.Unmarshal(respBody, apiResp)
	if err != nil {
		return nil, fmt.Errorf("json decode error: %s, Body: |%s|", err, respBody)
	}

	return apiResp, nil
}
