package services

import (
	"github.com/quangdangfit/go-backend/app/interfaces"
)

// Service struct
type Service struct {
	repo interfaces.IRepository
}

// NewService new customer service
func NewService(repo interfaces.IRepository) interfaces.IService {
	return &Service{
		repo: repo,
	}
}
