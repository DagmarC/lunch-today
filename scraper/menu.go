package scraper

import (
	"errors"
	"regexp"
	"strings"

	"github.com/DagmarC/lunch-today/restaurants"
	"github.com/DagmarC/lunch-today/utils"
	"github.com/PuerkitoBio/goquery"
)

func Menu(restaurant string) []restaurants.WeekLunch {

	allLunches := make([]restaurants.WeekLunch, 0)

	switch restaurant {
	case restaurants.All:
		fallthrough

	case restaurants.MonteBu:
		weekLunch := parseMenu(restaurants.UrlMB, restaurants.MonteBu, restaurants.SelectorDayMB, restaurants.SelectorSoupMB,
			restaurants.SelectorMealMB, true)

		allLunches = append(allLunches, *weekLunch)
		if restaurant != restaurants.All {
			break
		}
		fallthrough
	case restaurants.Viva:
		weekLunch := parseMenu(restaurants.UrlV, restaurants.Viva, restaurants.SelectorDayV, restaurants.SelectorMealV,
			restaurants.SelectorMealV, false)

		allLunches = append(allLunches, *weekLunch)
		if restaurant != restaurants.All {
			break
		}
		fallthrough
	case restaurants.Suzies:
		weekLunch := parseMenu(restaurants.UrlS, restaurants.Suzies, restaurants.SelectorDayS, restaurants.SelectorMealS,
			restaurants.SelectorMealS, false)

		allLunches = append(allLunches, *weekLunch)
	}

	return allLunches
}

func parseMenu(restaurantUrl, restaurant, daySelector, soupSelector, mealSelector string, soupSelect bool) *restaurants.WeekLunch {

	doc := loadDoc(restaurantUrl)

	weekLunch := restaurants.CreateWeekLunch(restaurant)

	days := findBySelector(daySelector, doc)
	err := parseDay(days, weekLunch)
	utils.Check(err)

	if soupSelect {
		soups := findBySelector(soupSelector, doc)
		err = parseSoup(soups, weekLunch)
		utils.Check(err)
	}

	meals := findBySelector(mealSelector, doc)
	err = parseMeal(meals, weekLunch, !soupSelect)
	utils.Check(err)

	return weekLunch

}

func findBySelector(selector string, doc *goquery.Document) []string {
	result := make([]string, 0)

	doc.Find(selector).Each(func(i int, item *goquery.Selection) {
		text := strings.TrimSpace(item.Text())
		if !strings.EqualFold(text, "") {
			result = append(result, strings.Trim(text, ""))
		}
	})
	return result
}

func parseDay(days []string, weekLunch *restaurants.WeekLunch) error {
	if len(days) == 0 {
		return errors.New("no day available")
	}

	regDay := regexp.MustCompile(`([P|p]ondělí|[Ú|ú]terý|[S|s]tředa|[Č|č]tvrtek|[P|p]átek).*`)
	for _, day := range days {
		match := regDay.FindStringSubmatch(day)

		if len(match) == 2 {
			lunch := restaurants.CreateDayLunch()
			lunch.SetDay(match[1])
			weekLunch.AddLunch(lunch)
		}
	}
	return nil
}

func parseSoup(soups []string, weekLunch *restaurants.WeekLunch) error {
	lunch := weekLunch.Lunch()

	if len(soups) == 0 {
		return errors.New("invalid input for soup")
	}
	for i, soup := range soups {
		lunch[i].SetSoup(strings.TrimSpace(soup))
	}
	return nil
}

func parseMeal(meals []string, weekLunch *restaurants.WeekLunch, soup bool) error {
	mealsInDay := len(meals) / len(weekLunch.Lunch()) // Days are already saved to determine the current length of the week.
	day := 0
	m := 0 // Counter for meal on a day.
	lunch := weekLunch.Lunch()

	for _, meal := range meals {

		if soup && m == 0 {
			lunch[day].SetSoup(strings.TrimSpace(meal))
		} else {
			lunch[day].AddMeal(strings.TrimSpace(meal))
		}
		m++

		if m == mealsInDay { // Multiple meals at one day. When reaches the meals in day, go to next day and reset counter.
			m = 0
			day++
		}
	}
	return nil
}
