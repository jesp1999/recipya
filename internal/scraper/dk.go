package scraper

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/reaper47/recipya/internal/models"
	"strconv"
	"strings"
)

func scrapeDk(root *goquery.Document) (models.RecipeSchema, error) {
	rs := models.NewRecipeSchema()

	content := root.Find("section[itemtype='http://schema.org/Recipe']")

	yieldStr, _ := content.Find("section[itemprop='recipeYield']").Attr("content")
	yield, _ := strconv.ParseInt(yieldStr, 10, 16)
	rs.Yield.Value = int16(yield)

	nodes := content.Find("li[itemprop='recipeIngredient']")
	rs.Ingredients.Values = make([]string, 0, nodes.Length())
	nodes.Each(func(_ int, s *goquery.Selection) {
		v := strings.TrimSpace(s.Text())
		v = strings.ReplaceAll(v, "\n", "")
		rs.Ingredients.Values = append(rs.Ingredients.Values, strings.Join(strings.Fields(v), " "))
	})

	rs.Instructions.Values = make([]models.HowToItem, 0)
	content.Find("div[itemprop='recipeInstructions'] h3,div[itemprop='recipeInstructions'] li").Each(func(i int, s *goquery.Selection) {
		if i > 0 && s.Nodes[0].Data == "h3" {
			rs.Instructions.Values = append(rs.Instructions.Values, models.NewHowToStep("\n"))
		}

		v := strings.ReplaceAll(s.Text(), "\n", "")
		v = strings.ReplaceAll(v, "\u00a0", "")
		rs.Instructions.Values = append(rs.Instructions.Values, models.NewHowToStep(strings.TrimSpace(v)))
	})

	description := content.Find("p[itemprop='description']").Text()
	rs.Description.Value = strings.TrimSpace(strings.ReplaceAll(description, "\n", ""))

	rs.Image.Value, _ = content.Find("meta[itemprop='url']").Attr("content")
	rs.DatePublished, _ = content.Find("meta[itemprop='datePublished']").Attr("content")
	rs.Name = content.Find("h1[itemprop='name']").Text()

	return rs, nil
}
