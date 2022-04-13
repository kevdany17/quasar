package test

import (
	"quasar/internal/models"
	"quasar/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestRepositorySuite struct {
	suite.Suite
	object models.Storage
}

func (suite *TestRepositorySuite) SetupTest() {
	suite.object = models.Storage{
		Satellite: models.SatelliteRequest{
			Name:     "Sato",
			Distance: 100,
			Message: []string{
				"Hello",
				"World",
				"Quasar",
			},
		},
		Position: models.Position{
			X: 100,
			Y: 75.5,
		}}

}
func (suite *TestRepositorySuite) TestListRepositoryAddSatellite() {
	respository := repository.NewListRepositoryImpl()

	size := len(models.DataBase)
	err := respository.Save(suite.object)

	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), size+1, len(models.DataBase))
}

func (suite *TestRepositorySuite) TestListRepositoryAddSatelliteNotDuplicate() {
	respository := repository.NewListRepositoryImpl()

	size := len(models.DataBase)

	err := respository.Save(suite.object)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), size, len(models.DataBase))
}

func (suite *TestRepositorySuite) TestListRepositoryGetForName() {
	respository := repository.NewListRepositoryImpl()

	suite.object = models.Storage{
		Satellite: models.SatelliteRequest{
			Name:     "skywalker",
			Distance: 114,
			Message: []string{
				"Hello",
				"World",
				"skywalker",
			},
		},
		Position: models.Position{
			X: -500,
			Y: -200,
		}}

	err := respository.Save(suite.object)
	satellite, err := respository.GetForName("Sato")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "Sato", satellite.Satellite.Name)
}

func (suite *TestRepositorySuite) TestListRepositoryGetForNameNotFound() {
	respository := repository.NewListRepositoryImpl()

	_, err := respository.GetForName("location")

	assert.EqualError(suite.T(), err, "Satellite not found")
}

func (suite *TestRepositorySuite) TestListRepositoryUpdate() {
	respository := repository.NewListRepositoryImpl()

	suite.object = models.Storage{
		Satellite: models.SatelliteRequest{
			Name:     "Sato",
			Distance: 500,
			Message: []string{
				"Hello",
				"World",
				"skywalker",
			},
		},
		Position: models.Position{
			X: -110,
			Y: 220,
		}}

	err := respository.Save(suite.object)
	assert.Nil(suite.T(), err)

	satellite, err := respository.GetForName("Sato")

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), float32(500), satellite.Satellite.Distance)
}

func TestRepositorySuiteInit(t *testing.T) {
	suite.Run(t, new(TestRepositorySuite))
}
