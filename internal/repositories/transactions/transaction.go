package transactions

type ITransaction interface {
	CreateTransaction(trans *Transaction) (Transaction, error)
	GetAllTransactions() ([]Transaction, error)
}
type DefaultTransactionService struct {
	Transaction ITransaction
}

var TransactionRepo DefaultTransactionService
