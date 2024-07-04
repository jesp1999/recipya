package scraper_test

import (
	"testing"
  "time"

	"github.com/reaper47/recipya/internal/models"
)

func TestScraper_N(t *testing.T) {
	testcases := []testcase{
		{
			name: "natashaskitchen.com",
			in:   "https://natashaskitchen.com/taco-seasoning-recipe",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "Condiments"},
				CookingMethod: nil,
				Cuisine:       &models.Cuisine{Value: "Mexican"},
				DatePublished: "2022-07-05T11:14:41+00:00",
				Description: &models.Description{
					Value: "Once you try this homemade taco seasoning blend, you&#039;ll never want to use store-bought again. This DIY Taco seasoning is so versatile for tacos, taco soup, and taco salad, as a seasoning for veggies, you name it! This recipe makes 3 Tbsp of seasoning.",
				},
				Keywords: &models.Keywords{Values: "homemade taco seasoning, taco seasoning, taco seasoning recipe"},
				Image:    &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"1 Tbsp chili powder",
						"2 tsp ground cumin",
						"1 tsp paprika (preferably smoked paprika)",
						"1 tsp fine sea salt",
						"1 tsp ground black pepper",
						"1/2 tsp garlic powder",
						"1/2 tsp onion powder",
						"1/2 tsp dried oregano",
						"Pinch red pepper flakes (optional)",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "In a small bowl, add all of the taco seasoning ingredients. Stir to combine and store homemade taco seasoning in an airtight container for up to 6 months."},
					},
				},
				Name: "Homemade Taco Seasoning Recipe",
				NutritionSchema: &models.NutritionSchema{
					Calories:       "7",
					Carbohydrates:  "1",
					Cholesterol:    "",
					Fat:            "0.3",
					Fiber:          "1",
					Protein:        "0.3",
					SaturatedFat:   "0.04",
					Servings:       "1",
					Sodium:         "274",
					Sugar:          "0.1",
					TransFat:       "",
					UnsaturatedFat: "0.2",
				},
				PrepTime:  "PT5M",
				TotalTime: "PT5M",
				Tools:     nil,
				Yield:     &models.Yield{Value: 9},
				URL:       "https://natashaskitchen.com/taco-seasoning-recipe",
			},
		},
		{
			name: "nigella.com",
			in:   "https://www.nigella.com/recipes/moonblush-tomatoes",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "uncategorized"},
				CookingMethod: &models.CookingMethod{},
				Cuisine:       &models.Cuisine{},
				Description: &models.Description{
					Value: "You may have come across Sunblush Tomatoes in delis; these are my homespun version. I halve cherry tomatoes, sprinkle them with herbs, a little salt and a pinch of sugar and drizzle over oil, put in a hot oven which I turn off immediately, leaving the tomatoes to cook in the residual heat overnight, hence “moonblush”. The taste of even disappointing tomatoes becomes, transformed this way, bright and intense. Think very essence of tomato. Eat in salads, alongside cheese, in pasta or however you wish.",
				},
				Image: &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"500 grams cherry tomatoes (on the vine or other baby tomatoes)",
						"2 teaspoons Maldon salt (or 1 teaspoon table salt)",
						"¼ teaspoon white sugar",
						"1 teaspoon dried thyme",
						"2 tablespoons olive oil",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "Preheat the oven to 220°C/200°C Fan/gas mark 7/450ºF."},
						{Type: "HowToStep", Text: "Cut the tomatoes in half and sit them cut side up in an ovenproof dish. Sprinkle with the salt, sugar, thyme and olive oil."},
						{Type: "HowToStep", Text: "Put them in the oven, and immediately turn it off. Leave the tomatoes in the oven overnight or for a day without opening the door."},
					},
				},
				Keywords:        &models.Keywords{Values: "Vegetarian, Gluten Free, Nibbly, Impatient, Express, Summer, Dairy Free, Egg Free, Nut Free, Vegan,"},
				Name:            "Moonblush Tomatoes",
				NutritionSchema: &models.NutritionSchema{},
				ThumbnailURL:    &models.ThumbnailURL{},
				Tools:           &models.Tools{Values: []models.HowToItem{}},
				Yield:           &models.Yield{Value: 80},
				URL:             "https://www.nigella.com/recipes/moonblush-tomatoes",
				Video:           &models.Videos{},
			},
		},
		{
			name: "ninjatestkitchen.eu",
			in:   "https://ninjatestkitchen.eu/recipe/dirt-worm-brownies",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "Halloween Inspiration"},
				CookingMethod: &models.CookingMethod{},
				Cuisine:       &models.Cuisine{},
				DatePublished: "2023-10-30T13:35:27+0000",
				Description: &models.Description{
					Value: "These Vegan Dirt and Worm Brownies are a labour of love but well worth the effort for a deliciously show stopping Halloween dessert. Packed full of healthier, nutritious ingredients to conventional brownies, and naturally gluten and grain-free they'll be perfect for most dietary requirements.",
				},
				Keywords: &models.Keywords{Values: "BN800UK, DT200UK, Halloween, Vegan"},
				Image:    &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"2 tbsp milled flax or chia",
						"6 tbsp cold water",
						"125g ground almonds",
						"125g coconut sugar",
						"30g cacao/cocoa powder",
						"1 tsp baking powder",
						"1/2 tsp sea salt",
						"60ml plant-based milk",
						"Optional Black food colouring or 1 tbsp activated charcoal",
						"75g coconut oil", "melted",
						"50g 85% dark chocolate", "melted",
						"400ml tin full fat coconut milk",
						"200ml grape juice",
						"200ml pomegranate juice",
						"4 tbsp agave nectar or maple syrup",
						"1 1/2 tbsp agar agar powder",
						"1 tsp cacao/cocoa powder",
						"100 plastic bendy straws",
						"120g coconut butter/manna",
						"80g 85% dark chocolate",
						"50g agave nectar",
						"125ml warm water (freshly boiled)",
						"1 tbsp cocoa/cacao powder",
						"¼ tsp vanilla extract",
						"Pinch sea salt",
						"Optional Black food colouring or 1 tbsp activated charcoal",
						"100g macadamia nuts",
						"50g almond flour/ground almonds",
						"1 tsp vanilla essence",
						"4 tbsp coconut sugar",
						"2 tbsp cocoa powder",
						"1 tbsp almond butter",
						"1 tbsp coconut flour",
						"1 tbsp activated charcoal (or black gel food colouring)",
						"50g 85% dark chocolate",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "Pre-heat oven to 160°C"},
						{Type: "HowToStep", Text: "Grease Ninja Foodi Zerostick Medium 20cm square tin and line with parchment paper with a little excess to help lift out the brownies later. Set to one side"},
						{Type: "HowToStep", Text: "Place milled flax in small bowl, gradually whisk in the cold water until combined then leave to one side for 5 minutes"},
						{Type: "HowToStep", Text: "Add all brownie ingredients except for the milk, food colouring, melted chocolate and coconut oil into a medium bowl and stir until fully combined. Stir in the food colouring into the milk followed by the flax mixture and beat into the dry ingredients. Finally, stir in the melted chocolate and coconut oil until combined"},
						{Type: "HowToStep", Text: "Pour mixture into prepared Ninja Foodi Zerostick Medium 20cm square tin. Bake for 15 to 20 minutes in centre of oven. Once baked leave to cool in tin"},
						{Type: "HowToStep", Text: "Whilst brownies are baking begin to prepare the worms. Extend all 100 plastic straws, tap down onto a surface and secure with elastic band then carefully place into a pint glass with the bendy ends at the bottom. Set to one side"},
						{Type: "HowToStep", Text: "Place all of the ingredients for the ‘worms’ into a small or medium pan and whisky until incorporated. Place onto a low to medium heat and bring to a rolling boil stirring occasionally. Allow to boil for 3 to 5 minutes then remove from the heat and set to one side to cool for 5 to 10 minutes before pouring into the straws. Place glass into the refrigerator and leave for 2 hours to set"},
						{Type: "HowToStep", Text: "Place all of the ingredients for the frosting into a small blender and blend until smooth. Immediately remove the lid and pour mixture over the brownies"},
						{Type: "HowToStep", Text: "Now place all dirt ingredients into a food processor and blend until the mixture becomes the texture of dirt. Evenly sprinkle the dirt over the brownies, reserving a couple of tablespoons for late, and gently pat down to allow the dirt to adhere. Place brownie tin into the refrigerator for 1 to 2 hours until set"},
						{Type: "HowToStep", Text: "Once the worms have set, carefully remove straws and push from the end opposite to the bendy side until the worm slides out. Repeat for all straws. You may find some break but don’t worry as you’ll be hiding them within the dirt"},
						{Type: "HowToStep", Text: "Remove brownies from the refrigerator and using a hot, damp knife, cut the brownies into 16 squares, wiping the knife between cuts"},
						{Type: "HowToStep", Text: "Add the worms to the top of the brownies using your creative flare, then sprinkle over a little dirt to hide any imperfections. Store in an airtight container for 3 to 5 days"},
					},
				},
				Name:            "‘Dirt’ & ‘Worm’ Brownies",
				NutritionSchema: &models.NutritionSchema{},
				PrepTime:        "PT1H",
				ThumbnailURL:    &models.ThumbnailURL{},
				Tools:           &models.Tools{Values: []models.HowToItem{}},
				Yield:           &models.Yield{Value: 16},
				URL:             "https://ninjatestkitchen.eu/recipe/dirt-worm-brownies",
				Video:           &models.Videos{},
			},
		},
		{
			name: "nourishedbynutrition.com",
			in:   "https://nourishedbynutrition.com/fudgy-gluten-free-tahini-brownies/",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "uncategorized"},
				CookTime:      "PT-475409H4M47S",
				DatePublished: "2022-02-09",
				Description: &models.Description{
					Value: "Rich and fudgy gluten-free tahini brownies that just happen to be also be grain-free and nut-free! These tahini brownies make for the perfect healthier chocolate dessert!",
				},
				Image: &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"1 cup tahini",
						"½ cup maple syrup",
						"2 eggs",
						"1 teaspoon vanilla",
						"⅓ cup cocoa powder",
						"½ teaspoon baking soda",
						"¼ teaspoon salt",
						"⅓ cup chocolate chips",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "Preheat oven to 350°F. Line an 8×8-inch baking pan with parchment paper."},
						{Type: "HowToStep", Text: "In a large bowl, combine tahini, maple syrup, eggs, and vanilla; whisk well to combine. The mixture will thicken quite a bit. Add cocoa powder, baking soda, and salt. Continue to mix until the mixture is smooth."},
						{Type: "HowToStep", Text: "Melt the chocolate chips in the microwave for 90 seconds, stopping every 30 seconds to mix (this can also be done on the stovetop). Add the melted chocolate to the batter and mix to combine."},
						{Type: "HowToStep", Text: "Transfer mixture to prepared baking pan. Bake for 23 to 25 minutes, or until a toothpick inserted in the center comes out mostly clean. Sprinkle with flaky salt."},
						{Type: "HowToStep", Text: "Let brownies cool completely in the pan. Lift parchment to remove the brownies from the pan. Cut into 12-16 squares."},
					},
				},
				Name:      "Fudgy Tahini Brownies",
				PrepTime:  "PT10M",
				TotalTime: "PT14M",
				Yield:     &models.Yield{Value: 12},
				URL:       "https://nourishedbynutrition.com/fudgy-gluten-free-tahini-brownies/",
			},
		},
		{
			name: "nosalty.hu",
			in:   "https://www.nosalty.hu/recept/kelt-pizzateszta",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "pizzatészta"},
				CookingMethod: &models.CookingMethod{},
				CookTime:      "PT12M",
				Cuisine:       &models.Cuisine{Value: "olasz"},
				DatePublished: "2014-09-30T05:45:00+00:00",
				Description: &models.Description{
					Value: "Az igazi pizzatészta friss élesztőből, BL &#39;00&#39;-ás finomlisztből és extra szűz olívaolajból az igazi. A tisztított víz pedig nem kötelező, de itt leírtam, hogy miért jobb választás.",
				},
				Keywords: &models.Keywords{
					Values: "buli receptek, ebéd, vacsora, ovo-lakto vegetáriánus, lakto vegetáriánus, ovo vegetáriánus, vegetáriánus, tejmentes, tojásmentes, laktózmentes, vegán, Finomliszt, Élesztő, Víz, Cukor, Finomliszt, Olívaolaj, Só, közepes, gyors",
				},
				Image: &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"500 g Finomliszt", "3 dkg Élesztő",
						"3 dl Víz lehetőség szerint tisztított", "0 teáskanál Cukor",
						"2 evőkanál Finomliszt", "4 ek Olívaolaj", "0 ízlés szerint Só",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "A langyos víz felébe belemorzsoljuk az élesztőt, elkavarjuk benne, majd hozzáadunk két evőkanál lisztet, és negyed ór\xc3\xa1ig érni hagyjuk a kovászt, amíg habos nem lesz a teteje."},
						{Type: "HowToStep", Text: "A liszthez öntjük a kovászt, a maradék vizet és az olívaolajat, majd alaposan összedolgozzuk a tésztát."},
						{Type: "HowToStep", Text: "Ha a tészta összeállt, akkor lisztezett deszkára tesszük, és további 10 percig dagasztjuk."},
						{Type: "HowToStep", Text: "A tésztát 2-3 részre szedjük (attól függően, mekkora és hány féle pizzát szeretnénk készíteni), majd zsemléket formálunk belőlük, és 45 percig letakarva kelesztjük őket."},
						{Type: "HowToStep", Text: "Miután megkeltek a tésztáink, egyenként átgyúrjuk őket, és újabb 45 percig másodszorra is megkelesztjük."},
						{Type: "HowToStep", Text: "A kétszer kelt tésztákat kinyújtjuk, megszórjuk a kívánt feltétekkel, és 220-250 fokra előmelegített sütőben barnára sütjük."},
					},
				},
				Name: "Kelt pizzatészta",
				NutritionSchema: &models.NutritionSchema{
					Calories:       "2253.255",
					Carbohydrates:  "412.3725",
					Cholesterol:    "0",
					Fat:            "37.764",
					Fiber:          "16.74",
					Protein:        "57.269",
					Servings:       "894.5 g",
					Sodium:         "32.265",
					Sugar:          "3.926",
					TransFat:       "",
					UnsaturatedFat: "",
				},
				PrepTime: "PT20M",
				ThumbnailURL: &models.ThumbnailURL{
					Value: "https://image-api.nosalty.hu/nosalty/images/recipes/Fu/LX/kelt-pizzateszta.jpeg?w=64&h=64&fit=crop&s=0fff8f5722bcf7e92685e39cf4898aef",
				},
				Tools:     &models.Tools{Values: []models.HowToItem{}},
				TotalTime: "PT32M",
				Yield:     &models.Yield{Value: 6},
				URL:       "https://www.nosalty.hu/recept/kelt-pizzateszta",
				Video: &models.Videos{
					Values: []models.VideoObject{
						{
							AtType:      "VideoObject",
							ContentURL:  "https://cdn.jwplayer.com/manifests/fM0VLO5e.m3u8",
							Description: "Az igazi pizzatészta friss élesztőből, BL '00'-ás finomlisztből és extra szűz olívaolajból az igazi. FELIRATKOZÁS: http://bit.ly/NoSalty-feliratkozas Recept: https://www.nosalty.hu/recept/kelt-pizzateszta Nosalty Magazin - http://nosalty.hu Facebook - https://www.facebook.com/NoSalty Instagram - https://www.instagram.com/nosalty Twitter - https://twitter.com/NoSalty Pinterest - https://hu.pinterest.com/nosalty",
							Duration:    "PT1M13S",
							EmbedURL:    "https://cdn.jwplayer.com/manifests/fM0VLO5e.m3u8",
							Name:        "Kelt pizzatészta",
							UploadDate:  time.Date(2018, 8, 17, 0, 0, 0, 0, time.FixedZone("CEST", 2*60*60)),
						},
					},
				},
			},
		},
		{
			name: "nrk.no",
			in:   "https://www.nrk.no/mat/japansk-omelett-_tamagoyaki_-1.16435297",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "Forrett/smårett"},
				CookingMethod: &models.CookingMethod{},
				Cuisine:       &models.Cuisine{Value: "Japan"},
				DateModified:  "2023-12-19T10:49:36+01:00",
				DatePublished: "2023-06-14T11:18:23+02:00",
				Description: &models.Description{
					Value: "Tamagoyaki er japansk rullet omelett. Den serveres oftest som smakfull siderett eller til frokost. Den lages ved å brette sammen tynne lag med krydret egg.",
				},
				Keywords: &models.Keywords{Values: "Panne/wok, Egg"},
				Image:    &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"8 egg", "30 g sukker", "1 ss mirin", "1 ss sake", "180 ml bonito dashi",
						"litt salt", "olje",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "Pisk eggene sammen med de andre ingrediensene."},
						{Type: "HowToStep", Text: "Ha oljen klar i en liten bolle. Dypp et stykke tørkepapir i oljen og smør pannen grundig."},
						{Type: "HowToStep", Text: "Ha litt av eggemassen på den varme pannen og la den stivne (Laget i en firkantet panne)."},
						{Type: "HowToStep", Text: "Brett omeletten to ganger, og skyv den over på motsatt side av pannen."},
						{Type: "HowToStep", Text: "Hell litt mer eggeblanding i pannen og gjenta prosessen. Tanken er at du ender opp med en rektangulær «omelettrulade», som s\xc3\xa5 kan skjæres i skiver."},
						{Type: "HowToStep", Text: "Pynt gjerne med vårløk."},
					},
				},
				Name:            "Japansk omelett (Tamagoyaki)",
				NutritionSchema: &models.NutritionSchema{},
				ThumbnailURL:    &models.ThumbnailURL{},
				Tools:           &models.Tools{Values: []models.HowToItem{}},
				URL:             "https://www.nrk.no/mat/japansk-omelett-_tamagoyaki_-1.16435297",
				Video:           &models.Videos{},
				Yield:           &models.Yield{Value: 1},
			},
		},
		{
			name: "number-2-pencil.com",
			in:   "https://www.number-2-pencil.com/creamy-one-pot-pumpkin-alfredo/",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "Dinner"},
				CookTime:      "PT15M",
				Cuisine:       &models.Cuisine{Value: "American"},
				DatePublished: "2023-10-15T10:46:53+00:00",
				Description:   &models.Description{},
				Keywords:      &models.Keywords{Values: "Pumpkin Alfredo"},
				Image:         &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"4 tablespoons butter", "2 cloves garlic (minced)",
						"2 cups low-sodium chicken broth", "1 cup heavy cream",
						"3/4 cup pumpkin (canned pure pumpkin)", "8 oz pasta (uncooked)",
						"1 cup parmesan cheese (freshly shredded)", "black pepper (to taste)",
						"kosher salt (to taste)", "freshly grated nutmeg (to taste)",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "In a large skillet, melt butter over medium heat and sauté garlic for 1-2 minutes."},
						{Type: "HowToStep", Text: "Add chicken broth, heavy cream, pumpkin, and pasta to the pot and stir to combine. Increase heat to medium-high until it starts to simmer. Stir again and reduce heat to medium-low. Cover and cook for 12-15 minutes, or until pasta is tender."},
						{Type: "HowToStep", Text: "Remove from heat and stir in parmesan cheese. Adjust amount of cheese, salt, pepper and nutmeg as needed."},
					},
				},
				Name: "Creamy One Pot Pumpkin Alfredo",
				NutritionSchema: &models.NutritionSchema{
					Calories:       "638",
					Carbohydrates:  "48",
					Cholesterol:    "114",
					Fat:            "41",
					Fiber:          "2",
					Protein:        "21",
					SaturatedFat:   "25",
					Servings:       "1",
					Sodium:         "457",
					Sugar:          "4",
					TransFat:       "0.5",
					UnsaturatedFat: "13",
				},
				PrepTime:  "PT15M",
				TotalTime: "PT30M",
				Yield:     &models.Yield{Value: 4},
				URL:       "https://www.number-2-pencil.com/creamy-one-pot-pumpkin-alfredo/",
			},
		},
		{
			name: "nutritionfacts.org",
			in:   "https://nutritionfacts.org/recipe/cinnamon-roll-oatmeal/",
			want: models.RecipeSchema{
				AtContext:     "https://schema.org",
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "Breakfast"},
				DatePublished: "2023-09-30T09:50:12+00:00",
				Description: &models.Description{
					Value: "This Cinnamon Roll Oatmeal is a naturally sweet and delicious way to start your day. It’s a date-sweetened oatmeal paired with a creamy cashew drizzle that makes for a fancy, yet simple breakfast. Less than 3 percent of Americans meet the daily recommended fiber intake, despite research suggesting that high-fiber foods, such as whole grains, can affect the progression of coronary heart disease. The soluble fiber of oatmeal forms a gel in the stomach, delaying stomach emptying, making one feel full for a longer period.",
				},
				Image: &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"4 cups water", "8 pitted dates", "2 teaspoons cinnamon",
						"2 teaspoons vanilla extract or powder", "2 cups rolled oats",
						"Raw pecans (optional garnish, as desired)", "½ cup raw cashews",
						"½ teaspoon vanilla extract", "2 pitted dates", "¾-1 cup water",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "To make the oatmeal, in a high-speed blender, combine the water, dates, cinnamon, and vanilla. Blend until smooth and pour into a saucepan with the rolled oats. Cook until the oats reach your desired consistency."},
						{Type: "HowToStep", Text: "To make the drizzle, combine all of the ingredients into a high-speed blender. Blend until smooth and creamy."},
						{Type: "HowToStep", Text: "Divide the oatmeal into bowls, top with the drizzle, and garnish with pecans, as desired."},
					},
				},
				Name:  "Cinnamon Roll Oatmeal",
				Yield: &models.Yield{Value: 4},
				URL:   "https://nutritionfacts.org/recipe/cinnamon-roll-oatmeal/",
			},
		},
		{
			name: "nytimes.com",
			in:   "https://cooking.nytimes.com/recipes/8357-spaghetti-with-fried-eggs",
			want: models.RecipeSchema{
				AtContext:     atContext,
				AtType:        &models.SchemaType{Value: "Recipe"},
				Category:      &models.Category{Value: "dinner, easy, for two, quick, weeknight, pastas, main course"},
				CookingMethod: &models.CookingMethod{Value: ""},
				Cuisine:       &models.Cuisine{Value: "italian"},
				Description: &models.Description{
					Value: "Here's a quick and delicious pasta dish to make when you have little time, and even less food in the house. All you need is a box of spaghetti, four eggs, olive oil and garlic (Parmesan is a delicious, but optional, addition).",
				},
				Keywords: &models.Keywords{Values: "egg, spaghetti, fall, vegetarian"},
				Image:    &models.Image{Value: anUploadedImage.String()},
				Ingredients: &models.Ingredients{
					Values: []string{
						"Salt",
						"1/2 pound thin spaghetti",
						"6 tablespoons extra virgin olive oil or lard",
						"2 large cloves garlic, lightly smashed and peeled",
						"4 eggs",
						"Freshly ground black pepper",
						"Freshly grated Parmesan or pecorino cheese, optional",
					},
				},
				Instructions: &models.Instructions{
					Values: []models.HowToItem{
						{Type: "HowToStep", Text: "Bring a pot of salted water to the boil. Start the sauce in the next step, and start cooking the pasta when the water boils."},
						{Type: "HowToStep", Text: "Combine garlic and 4 tablespoons of the oil in a small skillet over medium-low heat. Cook the garlic, pressing it into the oil occasionally to release its flavor; it should barely color on both sides. Remove the garlic, and add the remaining oil."},
						{Type: "HowToStep", Text: "Fry the eggs gently in the oil, until the whites are just about set and the yolks still quite runny. Drain the pasta, and toss with the eggs and oil, breaking up the whites as you do. (The eggs will finish cooking in the heat of the pasta.) Season to taste, and serve immediately, with cheese if you like."},
					},
				},
				Name: "Spaghetti With Fried Eggs",
				NutritionSchema: &models.NutritionSchema{
					Calories:       "607",
					Carbohydrates:  "58",
					Cholesterol:    "",
					Fat:            "34",
					Fiber:          "3",
					Protein:        "17",
					SaturatedFat:   "6",
					Sodium:         "381",
					Sugar:          "2",
					TransFat:       "0",
					UnsaturatedFat: "26",
				},
				ThumbnailURL: &models.ThumbnailURL{},
				Tools:        &models.Tools{Values: []models.HowToItem{}},
				TotalTime:    "PT20M",
				Yield:        &models.Yield{Value: 2},
				URL:          "https://cooking.nytimes.com/recipes/8357-spaghetti-with-fried-eggs",
				Video:        &models.Videos{},
			},
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			test(t, tc)
		})
	}
}
