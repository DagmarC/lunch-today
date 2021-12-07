package restaurants

import (
	"fmt"
	"strings"
)

type WeekLunch struct {
	restaurant string
	lunch      []*DayLunch
}

func CreateWeekLunch(name string) *WeekLunch {
	return &WeekLunch{restaurant: name, lunch: make([]*DayLunch, 0)}
}

func (w *WeekLunch) AddLunch(lunch *DayLunch) {
	w.lunch = append(w.lunch, lunch)
}

func (w *WeekLunch) Lunch() []*DayLunch {
	return w.lunch
}

func (w *WeekLunch) PrintLunch() {
	fmt.Printf("Restaurant: %v\n", w.restaurant)
	for _, lunch := range w.lunch {
		fmt.Println("---------------------------------")
		fmt.Println("Day: \t", (*lunch).day)
		fmt.Println("Soup: \t", (*lunch).soup)
		lunch.PrintMeal()
	}
}

type DayLunch struct {
	day  string
	soup string
	meal []string
}

func CreateDayLunch() *DayLunch {
	return &DayLunch{meal: make([]string, 0)}
}

func (d *DayLunch) SetDay(day string) {
	d.day = day
}

func (d *DayLunch) AddMeal(meal string) {
	d.meal = append(d.meal, meal)
}

func (d *DayLunch) SetSoup(soup string) {
	d.soup = soup
}

func (d *DayLunch) PrintMeal() {
	for _, m := range d.meal {
		fmt.Println("****")
		fmt.Println("\t", strings.TrimSpace(m))
	}
}
