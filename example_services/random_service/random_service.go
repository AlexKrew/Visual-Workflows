package random_service

import (
	"math/rand"
	"workflows/shared/nodes"
	"workflows/shared/shared_entities"
)

type RandomService struct{}

func NewRandomService() *RandomService {
	return &RandomService{}
}

func (service *RandomService) DoRandomThingsAdapter(input *nodes.NodeInput, output *nodes.NodeOutput) error {
	out := service.DoRandomThings()

	output.Set("output", shared_entities.NumberMessage(out))
	return nil
}

// Generates a pseudo-random number between 1 and 100
func (service *RandomService) DoRandomThings() int {
	min := 1
	max := 100
	randomNumber := rand.Intn(max-min) + min
	return randomNumber
}
