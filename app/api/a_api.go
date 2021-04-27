package api

import (
	"github.com/quangdangfit/go-backend/app/interfaces"
	"github.com/quangdangfit/go-backend/pkg/validation"
)

// API handle customer api
type API struct {
	service   interfaces.IService
	validator validation.Validation
}

// NewAPI new API
func NewAPI(service interfaces.IService, validator validation.Validation) *API {
	return &API{
		service:   service,
		validator: validator,
	}
}
