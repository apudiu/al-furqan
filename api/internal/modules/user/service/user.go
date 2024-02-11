package userservice

import (
	"github.com/apudiu/alfurqan/database"
	"github.com/apudiu/alfurqan/internal/model"
)

func Create(u *model.User) error {
	res := database.DB.Omit(
		"id", "created_at", "updated_at", "updated_at",
	).Create(u)

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
