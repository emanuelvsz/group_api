package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type animal struct {
	id            int
	name          string
	cientificName string
	pictureURL    string
}

var animals = []animal{
	{
		id: 1,
		name: "",
		cientificName: "",
		pictureURL: "",
	},
	{
		id: 2,
		name: "",
		cientificName: "",
		pictureURL: "",
	},
	{
		id: 3,
		name: "",
		cientificName: "",
		pictureURL: "",
	},
	{
		id: 4,
		name: "",
		cientificName: "",
		pictureURL: "",
	},
}

func main() {
	e := echo.New()

	jsonData, err := json.Marshal(animals)
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	e.GET("/", hello)

	e.Start(":8080")
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Olá, mundo!")
}
