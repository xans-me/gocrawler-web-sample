package kurs

// IKursService interface
type IKursService interface {
	IndexingKurs() (ResultIndexing []DataKurs, err error)
	InsertDataKurs(kurs DataKurs) (err error)
}

// IKursRepository interface
type IKursRepository interface {
	InsertERates(symbol string, eRate ERate, date string) (err error)
	InsertTTCounter(symbol string, tt TTCounter, date string) (err error)
	InsertBankNotes(symbol string, note BankNote, date string) (err error)
	IsExistKurs(data DataKurs) (countBN, countTT, countERates int, err error)
}
