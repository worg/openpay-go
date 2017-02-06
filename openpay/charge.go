package openpay

type (
	chargeClient struct {
		client *OpenPay
		path   string
	}

	// Transaction base type
	Transaction struct {
		ID              string       `json:"id,omitempty"`
		Amount          float32      `json:"amount"`
		Authorization   string       `json:"authorization,omitempty"`
		BankAccount     *BankAccount `json:"bank_account,omitempty"`
		CreationDate    string       `json:"creation_date,omitempty"`
		Currency        string       `json:"currency,omitempty"`
		CustomerID      string       `json:"customer_id,omitempty"`
		Description     string       `json:"description,omitempty"`
		ErrorMessage    interface{}  `json:"error_message,omitempty"`
		Method          string       `json:"method,omitempty"`
		OperationType   string       `json:"operation_type,omitempty"`
		OrderID         string       `json:"order_id,omitempty"`
		Status          string       `json:"status,omitempty"`
		TransactionType string       `json:"transaction_type,omitempty"`
	}

	// Charge structure is used to create and retrieve charges
	Charge struct {
		// Include Transaction fields
		ID              string       `json:"id,omitempty"`
		Amount          float32      `json:"amount"`
		Authorization   string       `json:"authorization,omitempty"`
		BankAccount     *BankAccount `json:"bank_account,omitempty"`
		CreationDate    string       `json:"creation_date,omitempty"`
		Currency        string       `json:"currency,omitempty"`
		CustomerID      string       `json:"customer_id,omitempty"`
		Description     string       `json:"description,omitempty"`
		ErrorMessage    interface{}  `json:"error_message,omitempty"`
		Method          string       `json:"method,omitempty"`
		OperationType   string       `json:"operation_type,omitempty"`
		OrderID         string       `json:"order_id,omitempty"`
		Status          string       `json:"status,omitempty"`
		TransactionType string       `json:"transaction_type,omitempty"`
		// Start of charge fields
		Card          *Card          `json:"card,omitempty"`
		DueDate       *timeStamp     `json:"due_date,omitempty"`
		OperationDate string         `json:"operation_date,omitempty"`
		PaymentMethod *PaymentMethod `json:"payment_method,omitempty"`
		Refund        *Transaction   `json:"refund,omitempty"`
		Metadata      interface{}    `json:"metadata,omitempty"`
	}

	// Card holds Credit/Debit card information
	Card struct {
		BankAccount
		ID              string   `json:"id"`
		Address         *Address `json:"address,omitempty"`
		AllowsCharges   bool     `json:"allows_charges,omitempty"`
		AllowsPayouts   bool     `json:"allows_payouts,omitempty"`
		Brand           string   `json:"brand,omitempty"`
		CardNumber      string   `json:"card_number,omitempty"`
		CustomerID      string   `json:"customer_id,omitempty"`
		ExpirationMonth string   `json:"expiration_month,omitempty"`
		ExpirationYear  string   `json:"expiration_year,omitempty"`
		Type            string   `json:"type,omitempty"`
	}

	// ExchangeRate is used when money conversions are needed
	ExchangeRate struct {
		Date  string  `json:"date,omitempty"`
		From  string  `json:"from,omitempty"`
		To    string  `json:"to,omitempty"`
		Value float32 `json:"value,omitempty"`
	}

	// PaymentMethod holds information of alternative payment
	// options like store payent
	PaymentMethod struct {
		Bank       string `json:"bank,omitempty"`
		BarcodeURL string `json:"barcode_url,omitempty"`
		Clabe      string `json:"clabe,omitempty"`
		Name       string `json:"name,omitempty"`
		Reference  string `json:"reference,omitempty"`
		Type       string `json:"type,omitempty"`
	}

	// BankAccount holds information about bank accounts
	BankAccount struct {
		Alias        string `json:"alias,omitempty"`
		BankCode     string `json:"bank_code,omitempty"`
		BankName     string `json:"bank_name,omitempty"`
		Clabe        string `json:"clabe,omitempty"`
		CreationDate string `json:"creation_date,omitempty"`
		HolderName   string `json:"holder_name,omitempty"`
	}

	// Address holds customer adresses
	Address struct {
		City        string `json:"city,omitempty"`
		CountryCode string `json:"country_code,omitempty"`
		Line1       string `json:"line1,omitempty"`
		Line2       string `json:"line2,omitempty"`
		Line3       string `json:"line3,omitempty"`
		PostalCode  string `json:"postal_code,omitempty"`
		State       string `json:"state,omitempty"`
	}
)

func newChargeClient(c *OpenPay) *chargeClient {
	return &chargeClient{
		client: c,
		path:   `charges`,
	}
}

// Create adds a new charge
func (c *chargeClient) Create(charge Charge) (*Charge, error) {
	res := new(Charge)
	if err := c.client.doRequest(`POST`, c.path, res, &charge); err != nil {
		return nil, err
	}

	return res, nil
}

// Get queries for a single charge
func (c *chargeClient) Get(id string) (*Charge, error) {
	path := c.path + `/` + id
	res := new(Charge)
	if err := c.client.doRequest(`GET`, path, res, nil); err != nil {
		return nil, err
	}

	return res, nil
}
