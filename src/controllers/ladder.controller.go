package controllers

import (
	"io"
	"net/http"
	"github.com/gofiber/fiber/v2"
)

func GetFormatLadderData(c *fiber.Ctx) error {
	var format = c.Query("format")
	if format == "" {
		return c.Status(400).JSON(fiber.Map{"message": "No format provided"})
	}

	url := "https://replay.pokemonshowdown.com/ladder/" + format + ".json"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error creating request"})
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error sending request"})
	}

	defer res.Body.Close()
	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error reading response"})
	}

	return c.JSON(fiber.Map{"message": "Received ladder data", "data": responseBody})
}