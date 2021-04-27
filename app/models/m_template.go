package models

import (
	"github.com/quangdangfit/go-backend/app/schema"
	"github.com/quangdangfit/go-backend/pkg/logger"
	"github.com/quangdangfit/go-backend/pkg/utils"
)

// Template model
type Template struct {
	Model `bson:",inline" json:",inline"`
}

// ToSchema convert Template model to Template Schema
func (m *Template) ToSchema() *schema.Template {
	var sc schema.Template
	err := utils.Copy(&sc, m)
	if err != nil {
		logger.Errorf("Failed to convert customer to schema, id: %s : %s", m.ID, err)
		return nil
	}

	sc.CreatedAt = utils.TimeToString(m.CreatedAt)
	sc.UpdatedAt = utils.TimeToString(m.UpdatedAt)
	return &sc
}

// Templates type
type Templates []*Template

// ToSchema Templates
func (a Templates) ToSchema() []*schema.Template {
	list := make([]*schema.Template, len(a))
	for i, item := range a {
		list[i] = item.ToSchema()
	}
	return list
}

// Get customer have index i
func (a Templates) Get(i int) *Template {
	if i < 0 || i > len(a) {
		return nil
	}
	return a[i]
}
