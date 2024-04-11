package domain

// Entity
type Template struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// UseCase (interface)
type TemplateUseCase interface {
	GetAll() ([]Template, error)
}

// Repository (interface)
type TemplateRepository interface {
	GetAll() ([]Template, error)
}
