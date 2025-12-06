package utils

import "golang.org/x/crypto/bcrypt"

// password hashing
func Hashing(Password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(Password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePassword
func HashCompare(HashedPass string, InputPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(HashedPass), []byte(InputPass)); 
	if err != nil {
		return err
	}
	return nil
}