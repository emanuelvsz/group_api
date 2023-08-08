package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type animal struct {
	ID            int    `json:"ID"`
	Name          string `json:"Name"`
	CientificName string `json:"cientific_Name"`
	PictureURL    string `picture_url`
}

func main() {
	e := echo.New()

	e.GET("/", hello)
	e.GET("/animals", getAnimals)

	e.Start(":8080")
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Olá, mundo!")
}

func getAnimals(c echo.Context) error {
	animals := []animal{
		{
			ID:            1,
			Name:          "Sapo-comum-do-sul",
			CientificName: "Bufo regularis",
			PictureURL:    "https://images.app.goo.gl/Fk811Bd3obN6qvt26",
		},
		{
			ID:            2,
			Name:          "Boto-cor-de-rosa",
			CientificName: "Inia geoffrensis",
			PictureURL:    "https://images.app.goo.gl/YSur1sdEdCrp6Rr18",
		},
		{
			ID:            3,
			Name:          "Camarão-pistola",
			CientificName: "AlpheIDae ",
			PictureURL:    "https://images.app.goo.gl/xBuM5SJCG2ybM3E47",
		},
		{
			ID:            4,
			Name:          "Urubu",
			CientificName: "Coragyps atratus",
			PictureURL:    "https://images.app.goo.gl/pxpGT2VrQqS1Bq1P7",
		},
	}

	fmt.Println(animals)

	return c.JSON(http.StatusOK, animals)
}
