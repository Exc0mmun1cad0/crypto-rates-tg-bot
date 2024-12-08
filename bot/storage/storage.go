package storage

type Storage interface {
	Add(username string, tokenList string) error
	GetAll(username string) ([]string, error)
	Delete(username string, tokenListr string) error
	DeleteAll(username string) error
}

// type Recording struct {
// 	TokenList string
// 	Username  string
// }
