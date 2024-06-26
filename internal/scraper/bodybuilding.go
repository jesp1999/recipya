package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/reaper47/recipya/internal/models"
	"github.com/reaper47/recipya/internal/utils/extensions"
	"strings"
)

func scrapeBodybuilding(root *goquery.Document) (models.RecipeSchema, error) {
	rs := models.NewRecipeSchema()

	rs.DateModified, _ = root.Find("meta[property='og:updated_time']").Attr("content")
	rs.DatePublished, _ = root.Find("meta[property='article:published_time']").Attr("content")
	rs.Description.Value = strings.TrimSpace(root.Find(".BBCMS__content--article-description").Text())

	name, _ := root.Find("meta[property='og:title']").Attr("content")
	before, _, found := strings.Cut(name, "|")
	if found {
		name = strings.TrimSpace(before)
	}
	rs.Name = name

	var n models.NutritionSchema
	nodes := root.Find(".bb-recipe__meta-nutrient-label")
	nodes.Each(func(_ int, sel *goquery.Selection) {
		switch sel.Text() {
		case "Calories":
			n.Calories = sel.Prev().Text() + " kcal"
		case "Carbs":
			n.Carbohydrates = sel.Prev().Text()
		case "Protein":
			n.Protein = sel.Prev().Text()
		case "Fat":
			n.Fat = sel.Prev().Text()
		}
	})
	rs.NutritionSchema = &n

	nodes = root.Find(".bb-recipe__ingredient-list-item")
	rs.Ingredients.Values = make([]string, 0, nodes.Length())
	nodes.Each(func(_ int, sel *goquery.Selection) {
		s := strings.ReplaceAll(sel.Text(), "\n", "")
		s = strings.Join(strings.Fields(s), " ")
		rs.Ingredients.Values = append(rs.Ingredients.Values, s)
	})

	nodes = root.Find(".bb-recipe__directions-list-item")
	rs.Instructions.Values = make([]models.HowToItem, 0, nodes.Length())
	nodes.Each(func(_ int, sel *goquery.Selection) {
		s := strings.ReplaceAll(sel.Text(), "\n", "")
		rs.Instructions.Values = append(rs.Instructions.Values, models.NewHowToStep(strings.TrimSpace(s)))
	})

	node := root.Find(".bb-recipe__directions-timing--prep").Find("time")
	rs.PrepTime, _ = node.Attr("datetime")

	node = root.Find(".bb-recipe__directions-timing--cook").Find("time")
	rs.CookTime, _ = node.Attr("datetime")

	nodes = root.Find(".bb-recipe__topic")
	xk := make([]string, 0, nodes.Length())
	nodes.Each(func(_ int, sel *goquery.Selection) {
		xk = append(xk, sel.Text())
	})

	rs.Keywords.Values = strings.Join(extensions.Unique(xk), ",")
	rs.Yield = &models.Yield{Value: findYield(root.Find(".bb-recipe__meta-servings .bb-recipe__meta-value-text").Text())}
	return rs, nil
}
