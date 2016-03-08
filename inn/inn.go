package inn

type Inn struct {
	Items []*Item
}

func NewInn() *Inn {
	return &Inn{
		Items: []*Item{
			{"+5 Dexterity Vest", 10, 20},
			{"Aged Brie", 2, 0},
			{"Elixir of the Mongoose", 5, 7},
			{"Sulfuras, Hand of Ragnaros", 0, 80},
			{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
			{"Conjured Mana Cake", 3, 6},
		},
	}
}

func (inn *Inn) UpdateQuality() {
	for _, item := range inn.Items {
		item.Update()
	}
}
