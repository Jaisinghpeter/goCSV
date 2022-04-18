package router

import(
	"github.com/gofiber/fiber/v2"
	"producer/controller/csv"
)

func StartRouter(){
	app := fiber.New()
	api := app.Group("/api/v1")
	api.Post("/addEmployee", addEmployeeToDB)
	app.Listen(":3000")
}

func addEmployeeToDB(context *fiber.Ctx) error {
	err := csv.AddCSVToDB(context)
	return err
}
