package services

import (
	"superstar/models"
)

// UserService handles CRUID operations of a admin datamodel,
// it depends on a admin repository for its actions.
// It's here to decouple the data source from the higher level compoments.
// As a result a different repository type can be used with the same logic without any aditional changes.
// It's an interface and it's used as interface everywhere
// because we may need to change or try an experimental different domain logic at the future.
type SuperstarService interface {
	GetAll() []models.StarInfo
	Get(id int) models.StarInfo
	Delete(id int) bool

	Update(user *models.StarInfo) error

	Create(user *models.StarInfo) error

	Search(country string) []models.StarInfo
}

