package transactions

import (
	log "Challenge/internal/adapters/Logger"
	Db "Challenge/internal/adapters/db"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

var db, _ = Db.NewDatabaseConnection()

func (s *DefaultTransactionService) GetAllTransactions() ([]Transaction, error) {
	var trans = []Transaction{}
	db.Find(&trans)
	if trans != nil {
		log.Info.Printf("c31f9477-7360-4668-a70a-b2e3dfbef262 Get transactions Scucessfully, %v", trans)
		return trans, nil
	} else {
		log.Error.Printf("d4cd8940-5296-415b-a0e6-257d68709e9f Error occur while getting all transactions %v", trans)
		return nil, errors.New("no transactions found")
	}
}

func (s *DefaultTransactionService) CreateTransaction(trans *Transaction) (Transaction, error) {
	if trans.Amount < LowAmount || trans.Amount > HighAmount {
		return Transaction{}, errors.New("the amount must be between 0 and 100000")
	}
	if len(trans.Currency) > NumberOfAvailableCharacter {
		return Transaction{}, errors.New("currency iso should be 3 char")
	}
	trans.Id = uuid.New()
	trans.Createdat = time.Now().String()
	trans.Status = false
	if err := db.Create(&trans).Error; err != nil {
		log.Info.Printf("5da33b37-e397-41c6-9276-5e71e4d55cab create transaction Scucessfully, %v", trans)
		return *trans, err
	}

	//Produce Kafa Event after create transaction
	Produce(trans)

	return *trans, nil
}

func (s *DefaultTransactionService) UpdateTransaction(JsonTransaction string) (bool, error) {
	var transaction Transaction
	json.Unmarshal([]byte(JsonTransaction), &transaction)
	transaction.Status = true
	if err := db.Model(&transaction).Update("status", true).Error; err != nil {
		return true, err
	}
	return false, nil
}
