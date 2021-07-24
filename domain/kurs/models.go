package kurs

type DataKurs struct {
	Symbol    string `json:"symbol"`
	ERate     `json:"e_rate"`
	TTCounter `json:"tt_counter"`
	BankNote  `json:"bank_notes"`
	Date      string `json:"date"`
}

type ERate struct {
	ERateBuy  float64 `json:"jual"`
	ERateSell float64 `json:"beli"`
}

type TTCounter struct {
	TTBuy  float64 `json:"jual"`
	TTSell float64 `json:"beli"`
}

type BankNote struct {
	BNBuy  float64 `json:"jual"`
	BNSell float64 `json:"beli"`
}
