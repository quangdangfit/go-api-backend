package interfaces

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// IDatabase interface
type IDatabase interface {
	GetInstance() *mongo.Database
}
