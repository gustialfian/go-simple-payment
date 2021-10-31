package order

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	db *sql.DB
}

type response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewHandler(db *sql.DB) *OrderHandler {
	return &OrderHandler{
		db: db,
	}
}

func (h *OrderHandler) Insert() fiber.Handler {

	orderRepo := NewOrderRepository(h.db)

	return func(c *fiber.Ctx) error {
		order := new(Order)

		if err := c.BodyParser(order); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		if !order.IsValid() {
			return c.Status(400).JSON(response{"ERROR", "Sum != 0"})
		}

		order.Proccess(orderRepo)

		log.Println("{}", order)

		return c.JSON(response{"SUCCESS", "OK"})
	}
}
