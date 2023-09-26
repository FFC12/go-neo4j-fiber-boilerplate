package main

import (
	"main/api"
	"main/database"
	_ "main/docs"
	"main/utils"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Fiber Neo4j Swagger API
// @version 2.0
// @description Fiber, Neo4j Boilerplate.
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	// Init config and globals
	utils.InitConfig()

	// Init logger, set True if you want to write the file instead of stdout
	// It will create a 'logs.txt' file and start printing to it.
	utils.InitLogger(false)

	ctx := database.InitDB()
	defer database.CloseDriver(ctx)
	app := fiber.New()

	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	app.Post("/user/signup", api.SignUpHandler)
	app.Post("/user/login", api.LoginHandler)

	app.Listen(":8000")
}
