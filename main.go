package main

import (
	"fiber-mongo-api/configs"
	"fiber-mongo-api/routes"

	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmfiber"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := fiber.New()
	tracer, err := apm.NewTracer("my-fiber-app", "")
	if err != nil {
		// Lida com erros de inicialização do APM
		panic(err)
	}
	defer tracer.Close()

	app.Use(apmfiber.Middleware(apmfiber.WithTracer(tracer)))

	configs.ConnectDB()

	routes.UserRoute(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Hello from Fiber & mongoDB"})
	})

	app.Listen(":6000")
}
