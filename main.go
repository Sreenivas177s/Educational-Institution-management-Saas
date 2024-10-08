package main

import (
	apiframework "chat-server/api-framework"
	"chat-server/utils"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func main() {
	// load env variables
	if err := godotenv.Load(); err != nil {
		panic("Error while loading env file")
	}
	// initializing data base
	// database.InitDataBase()

	app := initAppInstance()
	accessLogger := utils.RegisterAccessLogger(app)
	apiframework.HandleApiCall(app.Group(`/api/:version<regex(v\d{1,2})>`))
	// ws := app.Group("/ws/v1")

	// return not found
	app.All("/*", utils.ServeNotFoundHTML)
	defer accessLogger.Close()
	url := fmt.Sprintf("%s:%s", os.Getenv("DEV_HOST"), os.Getenv("DEV_PORT"))
	log.Fatal(app.Listen(url))
}

func initAppInstance() *fiber.App {
	app := fiber.New()
	app.Static(`/ui`, `./static`)
	log.Info("Server Initialized ", time.Now().Format(time.DateTime), " ....")
	return app
}
