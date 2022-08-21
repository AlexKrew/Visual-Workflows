package random_service

import "math/rand"

type RandomService struct{}

func NewRandomService() *RandomService {
	return &RandomService{}
}

// Generates a pseudo-random number between 1 and 100
func (service *RandomService) DoRandomThings() int {
	min := 1
	max := 100
	randomNumber := rand.Intn(max-min) + min
	return randomNumber
}
