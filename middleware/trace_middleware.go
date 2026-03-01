package middleware

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func TraceMiddleware(c fiber.Ctx) error {
	traceID := uuid.New().String()

	// Simpan di header
	c.Set("X-Trace-ID", traceID)

	// Simpan di locals
	c.Locals("trace_id", traceID)
	c.Locals("timestamp", time.Now().UTC().Format(time.RFC3339))

	return c.Next()
}
