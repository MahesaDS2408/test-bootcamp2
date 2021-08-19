package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/markbates/goth"
	gf "github.com/markbates/goth/providers/google"
)

const (
	key = "101667299166-h2pqpe8gkt84eaoknbq228er0uf91li4.apps.googleusercontent.com"
	sec = "rH2dauqgONNKM6feTSjt_3z2"
)

func main() {
	goth.UseProviders(
		gf.New(key, sec, "http://localhost:3000/auth/google/callback"),
	)

	app := fiber.New()

	app.Get("/auth/:provider/callback", func(ctx *fiber.Ctx) error {
		user, err := gf.CompleteUserAuth(ctx)
		if err != nil {
			return err
		}
		ctx.JSON(user)
		return nil
	})

	app.Get("/logout/:provider", func(ctx *fiber.Ctx) error {
		gf.Logout(ctx)
		ctx.Redirect("/")
		return nil
	})

	app.Get("/auth/:provider", func(ctx *fiber.Ctx) error {
		if gothUser, err := gf.CompleteUserAuth(ctx); err == nil {
			ctx.JSON(gothUser)
		} else {
			gf.BeginAuthHandler(ctx)
		}
		return nil
	})

	app.Get("/", func(ctx *fiber.Ctx) error {
		ctx.Format("<p><a href='/auth/google'>google</a></p>")
		return nil
	})

	log.Fatal(app.Listen(":3000"))

}
