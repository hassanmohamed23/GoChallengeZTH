package api

import (
	Log "Challenge/internal/adapters/Logger"
	TransactionService "Challenge/internal/repositories/transactions"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

//Get all transactions
func GetAll(w http.ResponseWriter, r *http.Request) {

	transactions, err := TransactionService.TransactionRepo.GetAllTransactions()
	if err != nil {
		Log.Error.Printf("e408cfac-208b-4ec2-bc03-bf7ea3a03e94 Error Occur during Get All transactions, %v", err)
		http.Error(w, fmt.Sprint(err), 400)
		return
	}
	json.NewEncoder(w).Encode(transactions)
	w.WriteHeader(200)
}

//create new transaction
func CreateTransaction(w http.ResponseWriter, r *http.Request) {

	trans := TransactionService.Transaction{}
	json.NewDecoder(r.Body).Decode(&trans)
	_, err := TransactionService.TransactionRepo.CreateTransaction(&trans)
	if err != nil {
		Log.Error.Printf("bc005f4e-9f4e-42cf-a3c8-4ea6e2089e18 Error Occur during create transactions,%v", trans)
		http.Error(w, fmt.Sprintf("Can't Add Transaction %s", err), 200)
		return
	}
	w.Write([]byte("The transaction added successfully\n"))
	json.NewEncoder(w).Encode(trans)
	w.WriteHeader(200)
}

func HandleRequest() {
	Router := chi.NewRouter()
	Router.Get("/transactions", GetAll)
	Router.Post("/transactions/create", CreateTransaction)
	log.Fatal(http.ListenAndServe(":8080", Router))
}
