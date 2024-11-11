package routes

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func SetupCountRoutes(app *fiber.App, rdb *redis.Client) {
	Increment := func(ctx context.Context, key string) (string, error) {
		count, err := rdb.Incr(ctx, key).Result()
		if err != nil {
			return "", err
		}
		return strconv.FormatInt(count, 10), nil
	}

	app.Get("/api/parking-space", func(c *fiber.Ctx) error {
		count, err := Increment(c.Context(), "space_views")
		if err != nil {
			return c.Status(500).SendString("Internal Server Error")
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"count": count,
		})
	})

	app.Get("/api/parking-space/count", func(c *fiber.Ctx) error {
		count, err := rdb.Get(c.Context(), "space_views").Result()
		if err == redis.Nil {
			return c.SendString("Current page views: 0")
		} else if err != nil {
			return c.Status(500).SendString("Internal Server Error")
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"count": count,
		})
	})

	app.Get("/api/parking-lots", func(c *fiber.Ctx) error {
		count, err := Increment(c.Context(), "lot_views")
		if err != nil {
			return c.Status(500).SendString("Internal Server Error")
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"count": count,
		})
	})

	app.Get("/api/parking-lots/count", func(c *fiber.Ctx) error {
		count, err := rdb.Get(c.Context(), "lot_views").Result()
		if err == redis.Nil {
			return c.SendString("Current page views: 0")
		} else if err != nil {
			return c.Status(500).SendString("Internal Server Error")
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"count": count,
		})
	})
}
