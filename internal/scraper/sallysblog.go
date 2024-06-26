package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/reaper47/recipya/internal/models"
	"strconv"
	"strings"
)

func scrapeSallysblog(root *goquery.Document) (models.RecipeSchema, error) {
	rs := models.NewRecipeSchema()

	rs.Description.Value, _ = root.Find("meta[name='description']").Attr("content")
	rs.Name = strings.ToLower(root.Find("h1").First().Text())

	prep := root.Find("p:contains('Zubereitungszeit')").Next().Text()
	split := strings.Split(prep, " ")
	isMin := strings.Contains(strings.ToLower(prep), "min")
	for i, s := range split {
		_, err := strconv.ParseInt(s, 10, 64)
		if err == nil && isMin {
			prep = split[i]
		}
	}
	rs.PrepTime = prep

	nodes := root.Find(".recipe-description").Next().Find(".hidden").First().Prev().Find("div.text-lg")
	rs.Ingredients.Values = make([]string, 0, nodes.Length())
	nodes.Each(func(_ int, sel *goquery.Selection) {
		s := sel.Text()
		s = strings.TrimSpace(s)
		if s != "" {
			rs.Ingredients.Values = append(rs.Ingredients.Values, s)
		}
	})

	nodes = root.Find(".recipe-description div p")
	rs.Instructions.Values = make([]models.HowToItem, 0, nodes.Length())
	nodes.Each(func(_ int, sel *goquery.Selection) {
		rs.Instructions.Values = append(rs.Instructions.Values, models.NewHowToStep(sel.Text()))
	})

	return rs, nil
}
