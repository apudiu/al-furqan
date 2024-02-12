package userservice

import (
	"crypto/md5"
	"github.com/apudiu/alfurqan/database"
	"github.com/apudiu/alfurqan/internal/helpers"
	"github.com/apudiu/alfurqan/internal/model"
)

func Create(u *model.User) error {
	res := database.DB.Create(u)

	if res.Error != nil {
		return res.Error
	}
	return nil
}

func GetById(id uint) (user model.User, err error) {
	res := database.DB.First(&user, id)
	if res.Error != nil {
		err = res.Error
	}
	return
}

func GetByEmail(email string) (user model.User, err error) {
	res := database.DB.Where("email = ?", email).First(&user)
	if res.Error != nil {
		err = res.Error
	}
	return
}

// UpdateTokenHash updates user last invoked tokens hash. @token will be hashed before insert into db
func UpdateTokenHash(user *model.User, token string) bool {
	hashedToken := ""

	// calculate token sum only if token is available
	if len(token) > 1 {
		tokenSum := md5.Sum([]byte(token))
		ht, err := helpers.HashStr(string(tokenSum[:]), 10)
		if err != nil {
			return false
		}
		hashedToken = ht
	}

	res := database.DB.Model(user).Update("token", hashedToken)
	return res.Error == nil
}
