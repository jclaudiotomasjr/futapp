package security

import "golang.org/x/crypto/bcrypt"

func Hash(key string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)
}

func CheckPassword(keyWithHash, keyString string) error {
	return bcrypt.CompareHashAndPassword([]byte(keyWithHash), []byte(keyString))
}
