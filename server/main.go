package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	piston "github.com/milindmadhukar/go-piston"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())
	app.Post("/code", func(c *fiber.Ctx) error {
		// for parsing the body
		var body = map[string]interface{}{}
		c.BodyParser(&body)

		// getting the code and language from the body in string format
		code := body["code"].(string)
		lang := body["lang"].(string)

		// creating a piston client
		client := piston.CreateDefaultClient()
		execution, err := client.Execute(lang, "",
			[]piston.Code{
				{Content: code},
			},
		)
		if err != nil {
			return c.SendString(err.Error())
		}

		// getting the output from the execution
		output := execution.GetOutput()
		// formtting the string to be sent as a response
		_, errFound := c.WriteString(fmt.Sprintf("<pre class='w-3/4 bg-gray-800 text-gray-300 p-2 rounded overflow-auto' id='output'>%s</pre>", output))
		if errFound != nil {
			return c.SendString(errFound.Error())
		}
		return c.SendStatus(fiber.StatusOK)
	})

	log.Println("\nServer is listening on port", ":8080")
	log.Fatal(app.Listen(":8080"))
}
