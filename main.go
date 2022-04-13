package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})
	app.Use(compress.New())
	app.Use(recover.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${locals:requestid} ${status} - ${latency} ${method} ${path}\n",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/os", func(c *fiber.Ctx) error {
		myOS, myArch := runtime.GOOS, runtime.GOARCH
		inContainer := "inside"
		if _, err := os.Lstat("/.dockerenv"); err != nil && os.IsNotExist(err) {
			inContainer = "outside"
		}

		str := "I'm running " + inContainer + " a container.\n" + "I'm running on " + myOS + "/" + myArch + "."
		return c.SendString(str)
	})

	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                                     // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGUSR2) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	log.Println("Running cleanup tasks...")

	// Your cleanup tasks go here e.g.
	// db.Close()
	// redisConn.Close()
	log.Println("Fiber was successful shutdown.")
}
