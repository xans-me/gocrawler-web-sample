package kurs

type DataKurs struct {
	Symbol    string `json:"symbol" bson:"symbol"`
	ERate     `json:"e_rate" bson:"e_rate"`
	TTCounter `json:"tt_counter" bson:"tt_counter"`
	BankNote  `json:"bank_notes" bson:"bank_note"`
	Date      string `json:"date" bson:"date"`
}

type ERate struct {
	ERateBuy  float64 `json:"beli" bson:"beli"`
	ERateSell float64 `json:"jual" bson:"jual"`
}

type TTCounter struct {
	TTBuy  float64 `json:"beli" bson:"beli"`
	TTSell float64 `json:"jual" bson:"jual"`
}

type BankNote struct {
	BNBuy  float64 `json:"beli" bson:"beli"`
	BNSell float64 `json:"jual" bson:"jual"`
}
