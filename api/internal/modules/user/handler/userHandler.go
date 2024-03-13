package userhandler

import (
	"github.com/apudiu/alfurqan/database"
	"github.com/apudiu/alfurqan/internal/hs"
	"github.com/apudiu/alfurqan/internal/model"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User

	// find all users in the database
	db.Find(&users)

	// If no user is present return an error
	if len(users) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(hs.Res(hs.ResData{
			Msg: "No users present",
		}))
	}

	// Else return users
	return c.JSON(hs.Res(hs.ResData{
		Status: true,
		Msg:    "Users found",
		D:      users,
	}))
}

func CreateUsers(c *fiber.Ctx) error {
	db := database.DB
	user := new(model.User)

	// Store the body in the user and return error if encountered
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(hs.Res(hs.ResData{
			Msg: "Review your input",
			D:   err,
		}))
	}

	// Create the User and return error if encountered
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(hs.Res(hs.ResData{
			Msg: "Could not create user",
			D:   err,
		}))
	}

	// Return the created user
	return c.JSON(hs.Res(hs.ResData{
		Status: true,
		Msg:    "Created user",
		D:      user,
	}))
}

func GetUser(c *fiber.Ctx) error {
	db := database.DB
	var user model.User

	// Read the param userId
	id := c.Params("userId")

	// Find the user with the given Id
	db.Find(&user, "id = ?", id)

	// If no such user present return an error
	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(hs.Res(hs.ResData{
			Msg: "No user present",
		}))
	}

	// Return the user with the Id
	return c.JSON(hs.Res(hs.ResData{
		Status: true,
		Msg:    "Users found",
		D:      user,
	}))
}

func UpdateUser(c *fiber.Ctx) error {
	type updateUser struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	db := database.DB
	var user model.User

	// Read the param userId
	id := c.Params("userId")

	// Find the user with the given Id
	db.Find(&user, "id = ?", id)

	// If no such user present return an error
	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(hs.Res(hs.ResData{
			Msg: "No user present",
		}))
	}

	// Store the body containing the updated data and return error if encountered
	var updateUserData updateUser
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(hs.Res(hs.ResData{
			Msg: "Review your input",
			D:   err,
		}))
	}

	// Edit the user
	user.Name = updateUserData.Name
	user.Email = updateUserData.Email

	// Save the Changes
	db.Save(&user)

	// Return the updated user
	return c.JSON(hs.Res(hs.ResData{
		Status: true,
		Msg:    "Users found",
		D:      user,
	}))
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.DB
	var user model.User

	// Read the param userId
	id := c.Params("userId")

	// Find the user with the given Id
	db.Find(&user, "id = ?", id)

	// If no such user present return an error
	if user.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(hs.Res(hs.ResData{
			Msg: "No user present",
		}))
	}

	// Delete the user and return error if encountered
	err := db.Delete(&user, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(hs.Res(hs.ResData{
			Msg: "Failed to delete user",
		}))
	}

	// Return success message
	return c.JSON(hs.Res(hs.ResData{
		Status: true,
		Msg:    "Deleted User",
	}))
}
