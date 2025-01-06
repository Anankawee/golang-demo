package main

import (
	"fmt"

	"github.com/few-enersys/golang-demo/pkg/calculate"
	"github.com/few-enersys/golang-demo/pkg/personal"

	"github.com/gofiber/fiber/v2"
)

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
