package controllers

import (
	"io"
	"net/http"
	"github.com/gofiber/fiber/v2"
)

func GetUserReplays(c *fiber.Ctx) error {
	var user = c.Query("user")
	if user == "" {
		return c.Status(400).JSON(fiber.Map{"message": "No user provided"})
	}
	
	url := "https://replay.pokemonshowdown.com/search.json?user" + user
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
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error reading response"})
	}

	return c.JSON(fiber.Map{"message": "Received replay", "data": body})
}

func GetUsersReplays(c *fiber.Ctx) error {
	var user1 = c.Query("user")
	var user2 = c.Query("user2")
	if user1 == "" || user2 == "" {
		return c.Status(400).JSON(fiber.Map{"message": "No user provided"})
	}

	url := "https://replay.pokemonshowdown.com/search.json?user" + user1 + "&user2=" + user2
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
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error reading response"})
	}

	return c.JSON(fiber.Map{"message": "Received replays", "data": body})
}

func GetFormatReplays(c *fiber.Ctx) error {
	var format = c.Query("format")
	if format == "" {
		return c.Status(400).JSON(fiber.Map{"message": "No format provided"})
	}

	url := "https://replay.pokemonshowdown.com/search.json?format" + format
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
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error reading response"})
	}

	return c.JSON(fiber.Map{"message": "Received replays", "data": body})
}

func GetUsersFormatReplays(c *fiber.Ctx) error {
	var user = c.Query("user")
	var format = c.Query("format")
	if user == "" || format == "" {
		return c.Status(400).JSON(fiber.Map{"message": "No user or format provided"})
	}

	url := "https://replay.pokemonshowdown.com/search.json?user" + user + "&format=" + format
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
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Error reading response"})
	}

	return c.JSON(fiber.Map{"message": "Received replays", "data": body})
}