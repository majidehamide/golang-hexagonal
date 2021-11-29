package main

import (
	"golang-hexagonal/app"
)

type Customer struct {
	Name    string `json:"name" xml:"name`
	City    string `json:"city" xml:"city`
	ZipCode string `json:"zip_code" xml:"zip_code`
}

func main() {
	app.Start()

}
