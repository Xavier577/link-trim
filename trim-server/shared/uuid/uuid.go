package uuid

import "github.com/aidarkhanov/nanoid/v2"

type UUIDService struct{}

type IUUIDService interface {
	Generate() (string, error)
}

func (uuidService *UUIDService) Generate() (string, error) {
	return nanoid.GenerateString(nanoid.DefaultAlphabet, 16)
}
