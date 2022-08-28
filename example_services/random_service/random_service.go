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
	// minMsg, _ := input.ValueFor("min")
	// maxMsg, _ := input.ValueFor("max")

	// out := service.DoRandomThings(minMsg.Value.(int), maxMsg.Value.(int))
	out := service.DoRandomThings(0, 100)

	output.Set("output", shared_entities.NumberMessage(out))
	return nil
}

// Generates a pseudo-random number between 1 and 100
func (service *RandomService) DoRandomThings(min int, max int) int {
	randomNumber := rand.Intn(max-min) + min
	return randomNumber
}
