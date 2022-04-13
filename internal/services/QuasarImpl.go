package services

import (
	"math"
	"quasar/internal/contracts"
	"quasar/internal/models"
	"strings"
)

type QuasarImpl struct {
}

func NewQuasarImpl() contracts.IQuasar { return &QuasarImpl{} }

func (quasar *QuasarImpl) GetMessage(messages ...[]string) (msg string) {
	var words []string
	var lastWord string = ""
	var sizeMax int = 0
	if len(messages) > 1 {
		for _, item := range messages {
			if sizeMax < len(item) {
				sizeMax = len(item)
			}
		}
		for index, item := range messages {
			array := make([]string, sizeMax)
			copy(array, item)
			messages[index] = array
		}
		for index := 0; index < sizeMax; index++ {
			if index == 0 && messages[0][index] != "" {
				lastWord = messages[0][index]
				words = append(words, lastWord)
			}
			if messages[0][index] != "" && messages[0][index] != lastWord {
				lastWord = messages[0][index]
				words = append(words, lastWord)
			}
			if messages[1][index] != "" && messages[1][index] != lastWord {
				lastWord = messages[1][index]
				words = append(words, lastWord)
			}
			if messages[2][index] != "" && messages[2][index] != lastWord {
				lastWord = messages[2][index]
				words = append(words, lastWord)
			}
		}
	} else {
		for index := 0; index < len(messages[0]); index++ {
			if messages[0][index] != "" {
				lastWord = messages[0][index]
				words = append(words, lastWord)
			}
		}
	}
	return strings.Join(words, " ")
}
func (quasar *QuasarImpl) GetLocation(distances ...float32) (x, y float32) {
	satellites := models.NewSatellitesList()
	sum := 0.0
	for _, distance := range distances {
		sum += float64(distance)
	}
	if len(distances) > 0 && sum > 0 {
		list := make([]float32, 3)
		copy(list, distances)
		AB := math.Sqrt(math.Pow(float64(satellites[1].X-satellites[0].X), 2) + math.Pow(float64(satellites[1].Y-satellites[0].Y), 2))
		i := satellites[2].X - 0
		j := satellites[2].Y - 0
		x = float32((math.Pow(float64(list[0]), 2) - math.Pow(float64(list[1]), 2) + math.Pow(float64(AB), 2)) / (2 * AB))
		y = float32((math.Pow(float64(list[0]), 2) - math.Pow(float64(list[2]), 2) - math.Pow(float64(x), 2) + math.Pow(float64(x-i), 2) + math.Pow(float64(j), 2)) / float64(2*j))
	}
	return x, y
}
