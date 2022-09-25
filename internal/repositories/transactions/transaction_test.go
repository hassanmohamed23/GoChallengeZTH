package transactions

import (
	"testing"
)

func TestCreateTransaction(t *testing.T) {
	testTrans := Transaction{
		Amount:   150,
		Currency: "USD",
	}
	_, err := TransactionRepo.CreateTransaction(&testTrans)
	if err != nil {
		t.Errorf("The t failed and return unexpected Error: %v", err)
	}
}

func TestGetAllTrans(t *testing.T) {
	_, err := TransactionRepo.GetAllTransactions()
	if err != nil {
		t.Errorf("Error in getting all transactions: %v", err)
	}
}
