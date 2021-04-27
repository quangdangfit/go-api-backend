package schema

import (
	"github.com/quangdangfit/go-backend/pkg/paging"
)

// Response schema
type Response struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
}

// ListResult schema
type ListResult struct {
	Paging *paging.Paging `json:"paging"`
	Data   interface{}    `json:"data"`
}
