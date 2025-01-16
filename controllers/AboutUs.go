package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khankhulgun/common/models"
	"github.com/lambda-platform/lambda/DB"
)

func AboutUs(c *fiber.Ctx) {
	var aboutus []models.AboutUs
	if err := DB.DB.Raw(`SELECT f_table_schema AS schema, f_table_name, f_geometry_column, coord_dimension, srid, type FROM geometry_columns`).Scan(&geometryTables).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(aboutus)
}
