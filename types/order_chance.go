package types

type OrderChance struct {
	BidFee string       `json:"bid_fee"`
	AskFee string       `json:"ask_fee"`
	Market ChanceMarket `json:"market"`
	BidAccount Account `json:"bid_account"`
	AskAccount Account `json:"ask_account"`
}

type ChanceMarket struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	OrderTypes []string `json:"order_types"`
	OrderSide  []string `json:"order_side"`
	Bid        Bid      `json:"bid"`
	Ask        Bid      `json:"ask"`
	MaxTotal   string   `json:"max_total"`
	State      string   `json:"state"`
}

type Bid struct {
	Currency  string `json:"currency"`
	PriceUnit string `json:"price_unit"`
	MinTotal  int64  `json:"min_total"`
}

type Account struct {
	Currency     string `json:"currency"`
	Balance      string `json:"balance"`
	Locked       string `json:"locked"`
	AvgBuyPrice  string `json:"avg_buy_price"`
	Modified     bool   `json:"avg_buy_price_modified"`
	UnitCurrency string `json:"unit_currency"`
}
