package main

import (
	"fmt"

	"github.com/DagmarC/lunch-today/restaurants"
	"github.com/DagmarC/lunch-today/scraper"
)

func main() {
	weekLunches := scraper.Menu(restaurants.All)

	for _, w := range weekLunches {
		fmt.Println("_______________________________________")
		fmt.Println("_______________________________________")
		w.PrintLunch()
	}
}
