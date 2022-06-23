package controllers

import (
	"auth-api/database"
	"auth-api/models"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	// リクエストデータをパース
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// パスワードのマッチチェック
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Passwords do not match!",
		})
	}

	// パスワードのエンコード
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		Password:  password,
	}

	// データ登録
	// CreateはGORMのメソッド
	database.DB.Create(&user)

	return c.JSON(user)
}
