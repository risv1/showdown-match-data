package controllers

import (
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetDexData(c *fiber.Ctx) error {
	
	url := "https://pokemonshowdown.com/data/pokedex.json"

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error creating request"})
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error sending request"})
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error reading response"})
	}

	return c.JSON(fiber.Map{"message": "Received dex data", "data": body})
}

func GetMovesData (c *fiber.Ctx) error {

	url := "https://pokemonshowdown.com/data/moves.json"

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error creating request"})
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error sending request"})
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error reading response"})
	}

	return c.JSON(fiber.Map{"message": "Received moves data", "data": body})
}