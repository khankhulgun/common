package common

import (
	"github.com/gofiber/fiber/v2"
	"github.com/khankhulgun/common/controllers"
	"github.com/khankhulgun/common/database/migrations"
	"github.com/khankhulgun/common/database/seeds"
	"github.com/lambda-platform/lambda/agent/agentMW"
	"github.com/lambda-platform/lambda/config"
)

func Set(app *fiber.App) {

	a := app.Group("/common/api")
	a.Get("/About-Us", agentMW.IsLoggedIn(), controllers.AboutUs)
	a.Get("/table-columns/:schema/:table", agentMW.IsLoggedIn(), controllers.TableColumns)

	if config.Config.App.Migrate == "true" {
		migrations.Migrate()
	}
	if config.Config.App.Seed == "true" {
		seeds.Seed()
	}
}
