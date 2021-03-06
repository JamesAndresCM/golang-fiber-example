package main

import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/logger" 
  "github.com/JamesAndresCM/golang-fiber-example/routes"
  "github.com/JamesAndresCM/golang-fiber-example/configuration"
)

func main() {
  app := fiber.New()
  app.Use(logger.New())

  configuration.GetConnection()
  routes.MovieRoutes(app)
  app.Get("/", func(c *fiber.Ctx) error {
    return c.Status(fiber.StatusOK).JSON(fiber.Map{
      "success": true,
      "message": "First Endpoint",
    })
  })
  err := app.Listen(":8000")

  if err != nil {
    panic(err)
  }
}
