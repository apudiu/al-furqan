package hs

import (
	"golang.org/x/crypto/bcrypt"
)

func HashStr(str string, round uint8) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), int(round))
	return string(bytes), err
}

type ResData struct {
	// is the operation successful
	Status bool `json:"status"`
	// any message
	Msg string `json:"msg"`
	// any data
	D any `json:"d"`
}

// Res formats response
func Res(d ResData) AnyMap {
	rd := AnyMap{
		"status": d.Status,
		"msg":    d.Msg,
		"d":      d.D,
	}

	return rd
}

type AnyMap = map[string]interface{}
