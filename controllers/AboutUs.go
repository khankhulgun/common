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

func TableColumns(c *fiber.Ctx) error {
	schema := c.Params("schema")
	tableName := c.Params("table")
	var columns []models.TableColumn

	if err := DB.DB.Raw(`SELECT column_name, data_type 
		FROM information_schema.columns 
		WHERE table_schema = ? AND table_name = ?`, schema, tableName).Scan(&columns).Error; err != nil {
		return c.Status(500).SendString("Failed")
	}

	return c.JSON(columns)
}
