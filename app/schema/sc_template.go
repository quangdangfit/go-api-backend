package schema

// Template schema
type Template struct {
	ID          string `json:"id"`
	AgentID     string `json:"agent_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	OfficePhone string `json:"office_phone"`
	Mobile      string `json:"mobile"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// TemplateCreateParams schema
type TemplateCreateParams struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	OfficePhone string `json:"office_phone"`
	Mobile      string `json:"mobile"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

// TemplateUpdateParams schema
type TemplateUpdateParams struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	OfficePhone string `json:"office_phone"`
	Mobile      string `json:"mobile"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Description string `json:"description"`
}

// TemplateQueryParams schema
type TemplateQueryParams struct {
	Page  int64 `json:"-" form:"page,omitempty"`
	Limit int64 `json:"-" form:"limit,omitempty"`
}
