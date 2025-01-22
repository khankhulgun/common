package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khankhulgun/common/models"
	"github.com/lambda-platform/lambda/DB"
)

func AboutUs(c *fiber.Ctx) error {
	var aboutus []models.AboutUs

	if err := DB.DB.Raw(`SELECT title, description, image, list_order FROM about_us`).Scan(&aboutus).Error; err != nil {
		return c.Status(500).SendString("Failed")
	}

	return c.JSON(aboutus)
}
