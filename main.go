package main

import(
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"net/http"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func RenderIndex(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func RenderCollections(c *fiber.Ctx) error {
	service := os.Getenv("COLLECTION_SERVICE_DOMAIN")
	response, err := http.Get(service)
	if err != nil {
		fmt.Println("Error in calling Colleciton domain")
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return c.Send(responseData)
}

func main() {
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	app.Static("/", "./static")

	app.Get("/", RenderIndex)
	app.Get("/collections", RenderCollections)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	log.Fatal(app.Listen("0.0.0.0:"+port))
}
