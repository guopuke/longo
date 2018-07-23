package auth

import "golang.org/x/crypto/bcrypt"

// Encrypt encrypts the plain text with bcrypt.
func Encrypt(source string) (string, error) {
	bytes, e := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(bytes), e
}

// Compare compares the encrypted text with the plain text if it's the same.
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
