package test

import (
	"quasar/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuasarImpl_GetMessageInitialData(t *testing.T) {
	quasar := services.NewQuasarImpl()
	array01 := []string{"este", "", "", "mensaje", ""}
	array02 := []string{"", "es", "", "", "secreto"}
	array03 := []string{"este", "", "un", "", ""}

	expected := "este es un mensaje secreto"

	assert.Equal(t, expected, quasar.GetMessage(array01, array02, array03))
}

func TestQuasarImpl_GetMessageOnlyOneArray(t *testing.T) {
	quasar := services.NewQuasarImpl()
	array01 := []string{"este", "", "", "mensaje", ""}

	expected := "este mensaje"
	assert.Equal(t, expected, quasar.GetMessage(array01))
}

func TestQuasarImpl_GetMessageSizeDiffArray(t *testing.T) {
	quasar := services.NewQuasarImpl()
	array01 := []string{"este", "", "", "mensaje", "", "nuevo"}
	array02 := []string{"", "es", "", "", "secreto"}
	array03 := []string{"este", "un", "", ""}

	expected := "este es un mensaje secreto nuevo"
	assert.Equal(t, expected, quasar.GetMessage(array01, array02, array03))
}

func TestQuasarImpl_GetMessageEmptyArray(t *testing.T) {
	quasar := services.NewQuasarImpl()
	array01 := []string{"", "", "", "", "", ""}
	var array02 []string

	expected01 := ""
	assert.Equal(t, expected01, quasar.GetMessage(array01))

	expected02 := ""
	assert.Equal(t, expected02, quasar.GetMessage(array02))
}

func TestQuasarImpl_GetLocationInitialData(t *testing.T) {
	quasar := services.NewQuasarImpl()
	array01 := []float32{100, 115.5, 142.7}
	var expectedX float32 = 301.39245
	var expectedY float32 = -258.77872

	x, y := quasar.GetLocation(array01...)

	assert.Equal(t, expectedX, x)
	assert.Equal(t, expectedY, y)
}

func TestQuasarImpl_GetLocationOnlyOneDistance(t *testing.T) {
	quasar := services.NewQuasarImpl()
	array01 := []float32{100}

	x, y := quasar.GetLocation(array01...)

	assert.NotEqual(t, x, float32(0))
	assert.NotEqual(t, y, float32(0))
}

func TestQuasarImpl_GetLocationEmptyDistance(t *testing.T) {
	quasar := services.NewQuasarImpl()
	var array01 []float32

	x, y := quasar.GetLocation(array01...)

	assert.Equal(t, x, float32(0))
	assert.Equal(t, y, float32(0))
}

func TestQuasarImpl_GetLocationOnlyZeroDistance(t *testing.T) {
	quasar := services.NewQuasarImpl()
	array01 := []float32{0, 0, 0}

	x, y := quasar.GetLocation(array01...)

	assert.Equal(t, x, float32(0))
	assert.Equal(t, y, float32(0))
}
