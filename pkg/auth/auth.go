package auth

import "golang.org/x/crypto/bcrypt"

func Encrypt(source string) (string, error) {
	bytes, e := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(bytes), e
}
