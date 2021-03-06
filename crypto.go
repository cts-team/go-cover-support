package go_cover_support

import "golang.org/x/crypto/bcrypt"

func PasswordHash(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

func PasswordHashAssign(password []byte, param *string) error {
	pwd, err := PasswordHash(password)
	if err != nil {
		return err
	}
	*param = string(pwd)
	return nil
}

func PasswordVerify(hashedPassword, password []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	return err == nil
}
