package services

import "golang.org/x/crypto/bcrypt"

type DefaultPasswordHasher struct{}

func (h *DefaultPasswordHasher) Hash(pw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(bytes), err
}

func (h *DefaultPasswordHasher) CheckHash(pw string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	return err == nil
}
