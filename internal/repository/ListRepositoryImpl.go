package repository

import (
	"errors"
	"quasar/internal/contracts"
	"quasar/internal/models"
)

type ListRepositoryImpl struct{}

func NewListRepositoryImpl() contracts.IRepository {
	return &ListRepositoryImpl{}
}
func (r *ListRepositoryImpl) Save(model models.Storage) error {
	for index, v := range models.DataBase {
		if v.Satellite.Name == model.Satellite.Name {
			models.DataBase[index] = model
			return nil
		}
	}
	models.DataBase = append(models.DataBase, model)
	return nil
}
func (r *ListRepositoryImpl) GetForName(name string) (*models.Storage, error) {
	for _, item := range models.DataBase {
		if item.Satellite.Name == name {
			return &item, nil
		}
	}
	return nil, errors.New("Satellite not found")
}
