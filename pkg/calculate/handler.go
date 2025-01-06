package calculate

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoute(a *fiber.App) {
	app := a.Group("calculate")
	app.Post("/", calculateHandler)
}

func calculator(value1, value2 int, op string) (result int, err error) {
	if op == "sum" {
		result = value1 + value2
	} else if op == "sub" {
		result = value1 - value2

	} else if op == "multiple" {
		result = value1 * value2

	} else if op == "div" {
		result = value1 / value2
	} else {
		return 0, fmt.Errorf("%v", "Operation is invalids.")

	}

	if result == 20 {
		return 0, fmt.Errorf("%v", "The result is 20.")
	}

	return result, nil
}

func calculateHandler(c *fiber.Ctx) error {
	rc := new(requestCal)

	if err := c.BodyParser(rc); err != nil {
		return err
	}

	result, err := calculator(rc.Value1, rc.Value2, rc.Operation)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}

	s1 := strconv.FormatInt(int64(result), 10)
	return c.Status(fiber.StatusOK).SendString(s1)
}
