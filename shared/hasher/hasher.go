package hasher

import "golang.org/x/crypto/bcrypt"

type HashService struct{}

type IHashService interface {
	Hash(pass string) (string, error)
	Compare(pass, hash string) bool
}

func (hasherService *HashService) Hash(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}

func (hasherService *HashService) Compare(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
