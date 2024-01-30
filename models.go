package pay

type OnRampExperience string

type Credentials struct {
	ApiKey string `json:"X-CC-Api-Key"`
	AppId  string `json:"CBPAY-APP-ID"`
}

type BuyConfigResponse struct {
	Data *ConfigData `json:"json:data"`
}

type ConfigData struct {
	Countries []Countries `json:"countries"`
}

type Countries struct {
	Id             string           `json:"id"`
	Subdivisions   []string         `json:"subdivisions"`
	PaymentMethods []PaymentMethods `json:"payment_methods"`
}

type PaymentMethods struct {
	Id string `json:"id"`
}

type BuyOptionsRequest struct {
	Country     string  `json:"country"`
	Subdivision *string `json:"subdivision,omitempty"`
}

type BuyOptionsResponse struct {
	PaymentCurrencies  []Currencies         `json:"payment_currencies"`
	PurchaseCurrencies []PurchaseCurrencies `json:"purchase_currencies"`
}

type PurchaseCurrencies struct {
	Id       string     `json:"id"`
	Name     string     `json:"name"`
	Symbol   string     `json:"symbol"`
	Networks []Networks `json:"networks"`
}

type Currencies struct {
	Id                  string   `json:"id"`
	PaymentMethodLimits []Limits `json:"payment_method_limits"`
}

type Limits struct {
	Id  string `json:"id"`
	Min string `json:"min"`
	Max string `json:"max"`
}

type Networks struct {
	Name            string `json:"name"`
	DisplayName     string `json:"display_name"`
	ChainId         string `json:"chain_id"`
	ContractAddress string `json:"contract_address"`
}

type BuyQuotePayload struct {
	PurchaseCurrency string  `json:"purchase_currency"`
	PurchaseNetwork  *string `json:"purchase_network,omitempty"`
	PaymentAmount    string  `json:"payment_amount"`
	PaymentCurrency  string  `json:"payment_currency"`
	PaymentMethod    string  `json:"payment_method"`
	Country          string  `json:"country"`
	Subdivision      *string `json:"subdivision,omitempty"`
}

type BuyQuoteResponse struct {
	PaymentTotal    Money  `json:"payment_total"`
	PaymentSubtotal Money  `json:"payment_subtotal"`
	PurchaseAmount  Money  `json:"purchase_amount"`
	CoinbaseFee     Money  `json:"coinbase_fee"`
	NetworkFee      Money  `json:"network_fee"`
	QuoteId         string `json:"quote_id"`
}

type Money struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type TransactionRequest struct {
	PartnerUserId string  `json:"partner_user_id"`
	PageKey       *string `json:"page_key,omitempty"`
	PageSize      *int    `json:"page_size,omitempty"`
}

type TransactionResponse struct {
	Transactions []OnrampTransaction `json:"transactions"`
	NextPageKey  string              `json:"next_page_key"`
	TotalCount   string              `json:"total_count"`
}

type OnrampTransaction struct {
	Status           string       `json:"status"`
	PurchaseCurrency string       `json:"purchase_currency"`
	PurchaseNetwork  string       `json:"purchase_network"`
	PurchaseAmount   AmountDetail `json:"purchase_amount"`
	PaymentTotal     AmountDetail `json:"payment_total"`
	PaymentSubtotal  AmountDetail `json:"payment_subtotal"`
	CoinbaseFee      AmountDetail `json:"coinbase_fee"`
	NetworkFee       AmountDetail `json:"network_fee"`
	ExchangeRate     AmountDetail `json:"exchange_rate"`
	TxHash           string       `json:"tx_hash"`
	CreatedAt        string       `json:"created_at"`
	Country          string       `json:"country"`
	UserID           string       `json:"user_id"`
	PaymentMethod    string       `json:"payment_method"`
	TransactionID    string       `json:"transaction_id"`
}

type AmountDetail struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type DestinationWallet struct {
	Address           string    `json:"address"`
	Blockchains       *[]string `json:"blockchains,omitempty"`
	Assets            *[]string `json:"assets,omitempty"`
	SupportedNetworks *[]string `json:"supportedNetworks,omitempty"`
}

type OnRampAppParams struct {
	DestinationWallets    []DestinationWallet `json:"destinationWallets"`
	PresetCryptoAmount    *float64            `json:"presetCryptoAmount,omitempty"`
	PresetFiatAmount      *float64            `json:"presetFiatAmount,omitempty"`
	DefaultNetwork        *string             `json:"defaultNetwork,omitempty"`
	DefaultExperience     *OnRampExperience   `json:"defaultExperience,omitempty"`
	HandlingRequestedUrls bool                `json:"handlingRequestedUrls,omitempty"`
	PartnerUserId         *string             `json:"partnerUserId,omitempty"`
}

type OnRampAggregatorAppParams struct {
	QuoteId              string  `json:"quoteId"`
	DefaultAsset         string  `json:"defaultAsset"`
	DefaultNetwork       *string `json:"defaultNetwork,omitempty"`
	DefaultPaymentMethod string  `json:"defaultPaymentMethod"`
	PresetFiatAmount     float64 `json:"presetFiatAmount"`
	FiatCurrency         string  `json:"fiatCurrency"`
}

type GenerateOnRampUrlOptions struct {
	AppId           string          `json:"appId"`
	Host            *string         `json:"host,omitempty"`
	OnRampAppParams OnRampAppParams `json:"generateOnRampUrlParameters"` //fix json naming??? does it need a name?
}

type ApiError struct {
	Code    int       `json:"code"`
	Message string    `json:"message"`
	Details *[]string `json:"details,omitempty"`
}

type Token struct {
	Token     string `json:"token"`
	ChannelId string `json:"channel_id"`
}
