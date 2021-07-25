package kurs

import (
	"database/sql"
)

type Repository struct {
	pg *sql.DB
}


func (r *Repository) InsertERates(symbol string, eRate ERate, date string) (err error) {
	query := `
   			INSERT INTO kurs.e_rates (symbol, buy, sell, indexing_date)    
   			VALUES($1, $2, $3, $4)
   `
	_, err = r.pg.Exec(query, symbol, eRate.ERateBuy, eRate.ERateSell, date)
	return
}

func (r *Repository) InsertTTCounter(symbol string, tt TTCounter, date string) (err error) {
	query := `
   			INSERT INTO kurs.tt_counter (symbol, buy, sell, indexing_date)    
   			VALUES($1, $2, $3, $4)
   `
	_, err = r.pg.Exec(query, symbol, tt.TTBuy, tt.TTSell, date)
	return
}

func (r *Repository) InsertBankNotes(symbol string, note BankNote, date string) (err error) {
	query := `
   			INSERT INTO kurs.bank_notes (symbol, buy, sell, indexing_date)    
   			VALUES($1, $2, $3, $4)
   `
	_, err = r.pg.Exec(query, symbol, note.BNBuy, note.BNSell, date)
	return
}

func (r *Repository) IsExistKurs(data DataKurs) (countBN, countTT, countERates int, err error) {
	queryBN := `
		SELECT 
			COUNT(*) as countBN 
		FROM
			kurs.bank_notes
		WHERE
			symbol = $1 and indexing_date = $2
	`

	err = r.pg.QueryRow(queryBN, data.Symbol, data.Date).Scan(&countBN)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			countBN = 0
		}
		return
	}

	queryTT := `
		SELECT 
			COUNT(*) as countTT 
		FROM
			kurs.tt_counter
		WHERE
			symbol = $1 and indexing_date = $2
	`

	err = r.pg.QueryRow(queryTT, data.Symbol, data.Date).Scan(&countTT)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			countTT = 0
		}
		return
	}

	queryERates := `
		SELECT 
			COUNT(*) as countTT 
		FROM
			kurs.e_rates
		WHERE
			symbol = $1 and indexing_date = $2
	`

	err = r.pg.QueryRow(queryERates, data.Symbol, data.Date).Scan(&countERates)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			countERates = 0
		}
		return
	}

	return
}

// NewRepository to create new repository instance
func NewRepository(db *sql.DB) *Repository {
	return &Repository{pg: db}
}
