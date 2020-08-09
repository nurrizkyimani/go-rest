package main

import ("github.com/gofiber/fiber"
				"go-rest/film"					

)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World 👋!")
}

func setupRoutes(app *fiber.App){
	app.Get("/api/v1/film", film.GetFilms)
	app.Get("/api/v1/film/:id", film.GetFilm)
	app.Post("/api/v1/film", film.NewFilm)
	app.Delete("api/v1/film/:id", film.DelFilm)

}

func main() {
	app := fiber.New()

	setupRoutes(app)

	app.Listen(8080)
}
