package contracts

import "quasar/internal/models"

type IRepository interface {
	Save(model models.Storage) error
	GetForName(name string) (*models.Storage, error)
}
