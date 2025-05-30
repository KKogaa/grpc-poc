package customerrors

import "errors"

var (
	ErrTransactionNotFound   = errors.New("transaction not found")
	ErrTransactionRepository = errors.New("transaction repository error")
)
