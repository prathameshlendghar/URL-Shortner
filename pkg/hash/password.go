package hash

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// 14 is the "cost". The higher the number, the slower it is to compute,
	// which protects against brute-force attacks. 14 is a great standard.
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePasswordHash(storedHashPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedHashPassword), []byte(password))
	return err == nil
}
