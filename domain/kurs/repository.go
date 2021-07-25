package kurs

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	mongoDB *mongo.Database
}

func (r *Repository) InsertERates(symbol string, eRate ERate, date string) (err error) {
	data := bson.M{"symbol": symbol, "jual": eRate.ERateSell, "beli": eRate.ERateBuy, "date": date}

	_, err = r.mongoDB.Collection("e_rates").InsertOne(context.Background(), data)

	return err
}

func (r *Repository) InsertTTCounter(symbol string, tt TTCounter, date string) (err error) {
	data := bson.M{"symbol": symbol, "jual": tt.TTSell, "beli": tt.TTBuy, "date": date}
	_, err = r.mongoDB.Collection("tt_counter").InsertOne(context.Background(), data)

	return
}

func (r *Repository) InsertBankNotes(symbol string, note BankNote, date string) (err error) {
	data := bson.M{"symbol": symbol, "jual": note.BNSell, "beli": note.BNBuy, "date": date}
	_, err = r.mongoDB.Collection("bank_notes").InsertOne(context.Background(), data)

	return
}

func (r *Repository) IsExistKurs(data DataKurs) (countBN, countTT, countERates int64, err error) {
	countBN, err = r.mongoDB.Collection("bank_notes").CountDocuments(context.Background(), bson.M{"symbol": data.Symbol, "date": data.Date})
	if err != nil {
		return
	}

	countTT, err = r.mongoDB.Collection("tt_counter").CountDocuments(context.Background(), bson.M{"symbol": data.Symbol, "date": data.Date})
	if err != nil {
		return
	}

	countERates, err = r.mongoDB.Collection("e_rates").CountDocuments(context.Background(), bson.M{"symbol": data.Symbol, "date": data.Date})
	if err != nil {
		return
	}

	return
}

// NewRepository to create new repository instance
func NewRepository(mongoDB *mongo.Database) *Repository {
	return &Repository{mongoDB: mongoDB}
}
