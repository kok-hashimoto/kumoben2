package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		log.Fatalf("failed to parse %s as int; %s", portStr, err.Error())
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", personHandler)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}

type responseSchema struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

func personHandler(c echo.Context) error {
	idStr := c.QueryParam("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Errorf("failed to parse %s as int; %w", idStr, err)
	}
	resp := responseSchema{ID: id}
	switch id % 3 {
	case 0:
		resp.Name = "Alice"
		resp.City = "Sapporo"
	case 1:
		resp.Name = "Bravo"
		resp.City = "Tokyo"
	default:
		resp.Name = "Charlie"
		resp.City = "Naha"
	}
	return c.JSON(http.StatusOK, resp)
}
