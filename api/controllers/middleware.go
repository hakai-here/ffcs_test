package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

func (r Repository) AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("session_id")
		sess, err := r.Rdb.GetAuthCache(cookie)
		if err != nil && err == redis.Nil {
			return c.Next()
		}
		if sess.Valid {
			return c.Status(http.StatusOK).JSON(message.Error("session already authenticated"))
		}
		return c.Next()
	}
}

func (r Repository) AdminMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("session_id")
		sess, err := r.Rdb.GetAuthCache(cookie)
		if err == redis.Nil {
			return c.Status(http.StatusUnauthorized).JSON(message.Error("Session  is not authenticated"))
		}
		if sess.IsAdmin && sess.Valid {
			return c.Next()
		}
		return c.Status(http.StatusForbidden).JSON(message.Error("user is not authorized to execute these functions"))
	}
}

func (r Repository) ApiMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		cookie := c.Cookies("session_id")
		sess, err := r.Rdb.GetAuthCache(cookie)
		if err == redis.Nil {
			return c.Status(http.StatusUnauthorized).JSON(message.Error("Session  is not authenticated"))
		}
		c.Locals("userid", sess.UserId)
		c.Locals("branch", sess.Branch)
		if sess.Valid {
			return c.Next()
		}
		return c.Status(http.StatusUnauthorized).JSON(message.Error("Session  is not authenticated"))
	}
}
