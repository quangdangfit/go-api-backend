package models

import "time"

// Model base model
type Model struct {
	ID        string    `bson:"id,omitempty" json:"id,omitempty"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
