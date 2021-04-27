package repositories

import (
	"github.com/quangdangfit/go-backend/app/interfaces"
)

// Repository struct
type Repository struct {
	db             interfaces.IDatabase
	collectionName string
}

// NewRepository new customer repository
func NewRepository(db interfaces.IDatabase) interfaces.IRepository {
	return &Repository{
		db:             db,
		collectionName: "models",
	}
}
