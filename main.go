package main

import(
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func RenderIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Static("/", "./static")

	app.Get("/", RenderIndex)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen("0.0.0.0:"+port))
}
