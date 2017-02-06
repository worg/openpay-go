package openpay

import (
	"errors"
)

var (
	// ErrMissingAPIKey is returned when the API Key is absent
	ErrMissingAPIKey = errors.New(`MISSING_API_KEY`)
	// ErrMissingMerchantID is returned when the Merchant ID is absent
	ErrMissingMerchantID = errors.New(`MISSING_MERCHANT_ID`)
)
