package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gustialfian/go-simple-payment/pkg/config"
	"github.com/gustialfian/go-simple-payment/pkg/database"
	"github.com/gustialfian/go-simple-payment/pkg/order"

	_ "github.com/lib/pq"
)

func main() {
	log.Println("order insertion service...")

	cfg := config.New()

	db, err := database.New(cfg.ConnectionDB)
	if err != nil {
		return
	}
	defer db.Close()

	app := fiber.New()
	app.Use(logger.New())
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON("hello world")
	})

	appV0 := app.Group("/v0")
	orderSvc := order.NewHandler(db)
	orderSvc.RegisterRoutes(appV0)

	log.Fatal(app.Listen(":3000"))
}
