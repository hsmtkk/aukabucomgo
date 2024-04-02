package base

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const CONTENT_TYPE = "Content-Type"
const APPLICATION_JSON = "application/json"
const X_API_KEY = "X-API-KEY"

type Client interface {
	Get(path string, params map[string]string) ([]byte, error)
	Post(path string, params map[string]string) ([]byte, error)
}

type clientImpl struct {
	port  int
	token string
}

type Environment int

const (
	PRODUCTION Environment = iota
	TEST
)

func New(env Environment, apiPassword string) (Client, error) {
	if env == PRODUCTION {
		return newProduction(apiPassword)
	} else if env == TEST {
		return newTest(apiPassword)
	} else {
		return nil, fmt.Errorf("unknown environment: %v", env)
	}
}

func newTest(apiPassword string) (Client, error) {
	clt := &clientImpl{port: TEST_PORT}
	if err := clt.auth(apiPassword); err != nil {
		return nil, err
	}
	return clt, nil
}

func newProduction(apiPassword string) (Client, error) {
	clt := &clientImpl{port: PRODUCTION_PORT}
	if err := clt.auth(apiPassword); err != nil {
		return nil, err
	}
	return clt, nil
}

func (clt *clientImpl) auth(apiPassword string) error {
	url, err := clt.makeURL("/token")
	if err != nil {
		return err
	}
	type requestSchema struct {
		APIPassword string `json:"APIPassword"`
	}
	body := requestSchema{APIPassword: apiPassword}
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bodyJSON))
	if err != nil {
		return fmt.Errorf("failed to make new HTTP request: %w", err)
	}
	req.Header.Add(CONTENT_TYPE, APPLICATION_JSON)
	resp, err := clt.doHTTP(req)
	if err != nil {
		return err
	}
	type responseSchema struct {
		ResultCode int    `json:"ResultCode"`
		Token      string `json:"token"`
	}
	decoded := responseSchema{}
	if err := json.Unmarshal(resp, &decoded); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	clt.token = decoded.Token
	return nil
}

func (clt *clientImpl) makeURL(path string) (string, error) {
	base := fmt.Sprintf("http://localhost:%d", clt.port)
	url, err := url.JoinPath(base, BASE_PATH, path)
	if err != nil {
		return "", fmt.Errorf("failed to make URL: %w", err)
	}
	return url, nil
}

func (clt *clientImpl) doHTTP(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP status code %d, status %s", resp.StatusCode, resp.Status)
	}
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read HTTP response: %w", err)
	}
	return result, nil
}

func (clt *clientImpl) Get(path string, params map[string]string) ([]byte, error) {
	url, err := clt.makeURL(path)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	query := req.URL.Query()
	for k, v := range params {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Add(X_API_KEY, clt.token)
	return clt.doHTTP(req)
}

func (clt *clientImpl) Post(path string, params map[string]string) ([]byte, error) {
	url, err := clt.makeURL(path)
	if err != nil {
		return nil, err
	}
	body, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %w", err)
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	req.Header.Add(CONTENT_TYPE, APPLICATION_JSON)
	req.Header.Add(X_API_KEY, clt.token)
	return clt.doHTTP(req)
}
