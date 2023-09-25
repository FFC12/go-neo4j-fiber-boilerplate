package api

import (
	controllers "main/controllers/user"
	"main/database"
	"main/models"
	"main/utils"

	"github.com/gofiber/fiber/v2"
)

// SignUp godoc
// @Summary User registration
// @Description It sign up users to database
// @Tags User
// @Accept application/json
// @Param data body models.UserSignUp true "SignUp input data"
// @Success 200
// @Router /user/signup [post]
func SignUpHandler(ctx *fiber.Ctx) error {
	user := new(models.UserSignUp)

	if err := ctx.BodyParser(user); err != nil {
		utils.ErrorLogger.Println("Error: ", err)
		return ctx.SendStatus(422)
	}

	if !controllers.UserSignUp(*database.Context, *user) {
		return ctx.JSON(map[string]string{
			"error": "Invalid email or password",
		})
	} else {
		return ctx.SendStatus(200)
	}
}

// Login godoc
// @Summary User login
// @Description User can login and get the JWT token
// @Tags User
// @Accept application/json
// @Param data body models.UserLogin true "Login input data"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/login [post]
func LoginHandler(ctx *fiber.Ctx) error {
	user := new(models.UserLogin)

	if err := ctx.BodyParser(user); err != nil {
		utils.ErrorLogger.Println("Error: ", err)
		return ctx.SendStatus(422)
	}

	// Check if it's email or username
	if !utils.ValidEmail(user.Email) {
		// If you wanna use username or email, handle here.
		utils.WarningLogger.Println("Email is no valid: ", user.Email)
	}

	data, err := controllers.UserLogin(*database.Context, *user)

	if err != nil {
		utils.ErrorLogger.Println("Error: ", err)

		return ctx.JSON(map[string]string{
			"error": "Invalid email or password",
		})
	} else {
		username := data["username"]
		id := data["id"]
		email := data["email"]

		token, err := utils.GenerateJWT(map[string]any{
			"id":       id,
			"username": username.(string),
			"email":    email.(string),
		}, 1)

		if err != nil {
			utils.ErrorLogger.Println("Error: ", err)

			return ctx.JSON(map[string]string{
				"error": "An unexpected error occurred",
			})
		}

		return ctx.JSON(map[string]string{
			"jwt": token,
		})

	}
}
