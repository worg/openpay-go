// Package openpay provides wrappers around OpenPay API
package openpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

type (
	// OpenPay type allows us to interact with the api
	OpenPay struct {
		client     *http.Client
		baseURL    *url.URL
		merchantID string
		apiKey     string
		Charges    *chargeClient
	}

	// Error holds API error data
	Error struct {
		Category    string `json:"category"`
		Description string `json:"description"`
		Code        int    `json:"error_code"`
		StatusCode  int    `json:"http_code"`
		RequestID   string `json:"request_id"`
	}

	// HookMessage holds information received from OpenPay WebHook
	HookMessage struct {
		EventDate   string  `json:"event_date,omitempty"`
		Transaction *Charge `json:"transaction,omitempty"`
		Type        string  `json:"type,omitempty"`
		Code        string  `json:"verification_code,omitempty"`
	}

	timeStamp string
)

const (
	jsonTimeF = `"` + time.RFC3339 + `"`
)

var (
	// Production base URL
	productionURL = `https://api.openpay.mx`
	// Sandbox base URL
	sandboxURL = `https://sandbox-api.openpay.mx`
	// API version
	apiVersion = `/v1/`
	// Default private api key
	defaultAPIKey = ``
	// Default merchant id
	defaultMerchantID = ``
)

func init() {
	defaultAPIKey = os.Getenv(`OPENPAY_PRIVATE_KEY`)
	defaultMerchantID = os.Getenv(`OPENPAY_MERCHANT_ID`)
}

// NewClient returns a new initialized OpenPay API Client given a valid
// Merchant Id and Key, productionReady determines the environment
// to make calls
func NewClient(merchantID, key string, prodictionReady bool) *OpenPay {
	if len(key) < 1 && len(defaultAPIKey) > 0 {
		key = defaultAPIKey
	}

	if len(merchantID) < 1 && len(defaultMerchantID) > 0 {
		merchantID = defaultMerchantID
	}

	// Set base url
	urlStr := sandboxURL
	if prodictionReady {
		urlStr = productionURL
	}

	baseURL, _ := url.Parse(urlStr + apiVersion)

	c := &OpenPay{
		client:     http.DefaultClient,
		merchantID: merchantID,
		apiKey:     key,
		baseURL:    baseURL,
	}

	c.Charges = newChargeClient(c)

	return c
}

func (c *OpenPay) doRequest(method, path string, responseData, requestData interface{}) error {
	req, err := c.prepare(method, path, requestData)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	if err := checkAPIError(resp); err != nil {
		return err
	}

	if responseData != nil {
		if err := json.NewDecoder(resp.Body).Decode(responseData); err != nil {
			return err
		}
	}

	return nil
}

func (c *OpenPay) prepare(method, path string, body interface{}) (*http.Request, error) {
	var (
		data = new(bytes.Buffer)
	)
	// Check that API key and Merchant ID are set
	switch {
	case len(c.apiKey) < 1:
		return nil, ErrMissingAPIKey
	case len(c.merchantID) < 1:
		return nil, ErrMissingMerchantID
	}

	trail, err := url.Parse(c.merchantID + `/` + path)
	if err != nil {
		return nil, err
	}

	url := c.baseURL.ResolveReference(trail)

	err = json.NewEncoder(data).Encode(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url.String(), data)
	if err != nil {
		return nil, err
	}

	req.Header.Add(`Content-type`, `application/json`)
	req.SetBasicAuth(c.apiKey, ``)

	return req, nil
}

// GetBaseURL returns the current environment base url
func (c *OpenPay) GetBaseURL() string {
	return c.baseURL.String()
}

func (e Error) Error() string {
	return fmt.Sprintf("Status Code: %d Error Code: %d Request id: %q Description: %q Category: %q",
		e.StatusCode,
		e.Code,
		e.RequestID,
		e.Description,
		e.Category)
}

func checkAPIError(r *http.Response) error {
	if r.StatusCode == http.StatusOK {
		return nil
	}

	var apiErr Error
	if err := json.NewDecoder(r.Body).Decode(&apiErr); err != nil {
		return err
	}

	return apiErr
}

func (t *timeStamp) Time() time.Time {
	time, _ := time.Parse(jsonTimeF, t.String())
	return time
}

func (t timeStamp) String() string {
	return string(t)
}
