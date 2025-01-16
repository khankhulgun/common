package khanmap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khankhulgun/khanmap/database/migrations"
	"github.com/khankhulgun/khanmap/database/seeds"
	"github.com/lambda-platform/lambda/config"
)

func Set(app *fiber.App) {

	if config.Config.App.Migrate == "true" {
		migrations.Migrate()
	}
	if config.Config.App.Seed == "true" {
		seeds.Seed()
	}
}
