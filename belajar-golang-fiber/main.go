package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Prefork:      true,
	})

	app.Use("/api", func(ctx *fiber.Ctx) error { // -> "/api" adalah prefix, jadi middleware nya cuman jalan di route yg ada "/api"
		fmt.Println("I'm middleware before processing request")
		err := ctx.Next()
		fmt.Println("I'm middleware after processing request")
		return err
	})

	app.Get("/api/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World")
	})

	if fiber.IsChild() {
		fmt.Println("I'm child process")
	} else {
		fmt.Println("I'm parent process")
	}

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
