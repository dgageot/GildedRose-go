package inn

type Item struct {
	Name    string
	SellIn  int
	Quality int
}

func (item *Item) Update() {
	if item.Name == "Sulfuras, Hand of Ragnaros" {
		return
	}

	item.SellIn -= 1

	switch item.Name {
	case "Aged Brie":
		item.increaseQuality()

		if item.SellIn < 0 {
			item.increaseQuality()
		}
	case "Backstage passes to a TAFKAL80ETC concert":
		item.increaseQuality()

		if item.SellIn < 10 {
			item.increaseQuality()
		}

		if item.SellIn < 5 {
			item.increaseQuality()
		}

		if item.SellIn < 0 {
			item.Quality = 0
		}
	default:
		item.decreaseQuality()

		if item.SellIn < 0 {
			item.decreaseQuality()
		}
	}
}

func (item *Item) increaseQuality() {
	if item.Quality < 50 {
		item.Quality += 1
	}
}

func (item *Item) decreaseQuality() {
	if item.Quality > 0 {
		item.Quality -= 1
	}
}
