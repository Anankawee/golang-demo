package main

import (
	"example/hello/pkg/calculate"
	"example/hello/pkg/personal"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// CRUD -> Create, Read, Update and Delete

// Get for get data, trick something good
// Get params
// Ex localhost:3000?name='test'&age='40'
// Ex1 localhost:3000/{{id}}/test?name=test

// Post for creation
// Wrong localhost:3000?name='test'&age='2'
// right localhost/add, payload: {name: 'test', age=2}

// Put or Patch for update or modification
// name address tel email password
//

// Delete for delete data only ?
// right localhost/del/{{id}}

func CheckError(v int) (string, error) {
	if v == 1 {
		return "true", nil
	} else {
		return "", fmt.Errorf("%v", "not 1")
	}

}

func main() {
	app := fiber.New()
	personal.SetUpRoute(app)
	calculate.SetUpRoute(app)
	app.Listen(":3000")
}
