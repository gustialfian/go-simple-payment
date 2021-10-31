package order

import (
	"github.com/gofiber/fiber/v2"
)

func (h *OrderHandler) RegisterRoutes(router fiber.Router) {
	or := router.Group("/order")
	or.Post("/", h.Insert())
}
