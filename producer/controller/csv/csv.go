package csv

import(
	"github.com/gofiber/fiber/v2"
	"producer/model"
	"producer/constants"
	"log"
)

func AddCSVToDB(context *fiber.Ctx) error{
	filePath := new(model.FilePath)
	err := context.BodyParser(filePath)
	if err != nil {
		log.Println(err)
		context.Status(constants.BAD_REQUEST)
		context.JSON(buildStatusMap(false, string("")))
		return err
	}
	go readCsvFile(filePath.FilePath)

	if err != nil {
		context.Status(constants.INTERNAL_SERVER_ERROR).JSON(buildStatusMap(true, "Error Parsing File"))
		return err
	}
	err = context.JSON(buildStatusMap(true, "File uploaded Successfully"))
	return err
}

func buildStatusMap(status bool, reply string) fiber.Map{
	return fiber.Map{
		"success": status,
		"message": reply,
	}
}
