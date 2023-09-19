package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//os me permite interactuar con el sistema (leer variables de entorno, esbalecer variables de entorno, leer archivos, etc)

	app := fiber.New()

	app.Use(cors.New()) //para cambiar las politicas e CORS, CORS evita que nuestra web se conecte a otros puertos que no tengan que ver con la pagina

	app.Static("/", "../client/dist")

	app.Get("/users", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"data": "usuario desde el backend",
		}) //le envio algo dentro del contexto.    fiber.Map es un struct
	})

	app.Listen(":3000")
	fmt.Println("server on port 3000")
}
