package helpers

import "golang.org/x/crypto/bcrypt"

func HashStr(str string, round uint8) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), int(round))
	return string(bytes), err
}
