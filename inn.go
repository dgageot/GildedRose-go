package main

import "fmt"

type Inn struct {
	items *Items
}

func NewInn() *Inn {
	inn := &Inn{
		items: NewItems(),
	}
	inn.items.Add(&Item{"+5 Dexterity Vest", 10, 20})
	inn.items.Add(&Item{"Aged Brie", 2, 0})
	inn.items.Add(&Item{"Elixir of the Mongoose", 5, 7})
	inn.items.Add(&Item{"Sulfuras, Hand of Ragnaros", 0, 80})
	inn.items.Add(&Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20})
	inn.items.Add(&Item{"Conjured Mana Cake", 3, 6})

	return inn
}

func (in *Inn) UpdateQuality() {
	for i := 0; i < in.items.Size(); i++ {
		if !(in.items.Get(i).GetName() == "Aged Brie") && !(in.items.Get(i).GetName() == "Backstage passes to a TAFKAL80ETC concert") {
			if (in.items.Get(i).GetQuality() > 0) {
				if !(in.items.Get(i).GetName() == "Sulfuras, Hand of Ragnaros") {
					in.items.Get(i).SetQuality(in.items.Get(i).GetQuality() - 1)
				}
			}
		} else {
			if (in.items.Get(i).GetQuality() < 50) {
				in.items.Get(i).SetQuality(in.items.Get(i).GetQuality() + 1)

				if (in.items.Get(i).GetName() == "Backstage passes to a TAFKAL80ETC concert") {
					if (in.items.Get(i).GetSellIn() < 11) {
						if (in.items.Get(i).GetQuality() < 50) {
							in.items.Get(i).SetQuality(in.items.Get(i).GetQuality() + 1)
						}
					}

					if (in.items.Get(i).GetSellIn()) < 6 {
						if (in.items.Get(i).GetQuality() < 50) {
							in.items.Get(i).SetQuality(in.items.Get(i).GetQuality() + 1)
						}
					}
				}
			}
		}

		if !(in.items.Get(i).GetName() == "Sulfuras, Hand of Ragnaros") {
			in.items.Get(i).SetSellIn(in.items.Get(i).GetSellIn() - 1)
		}

		if (in.items.Get(i).GetSellIn() < 0) {
			if !(in.items.Get(i).GetName() == "Aged Brie") {
				if !(in.items.Get(i).GetName() == "Backstage passes to a TAFKAL80ETC concert") {
					if (in.items.Get(i).GetQuality() > 0) {
						if !(in.items.Get(i).GetName() == "Sulfuras, Hand of Ragnaros") {
							in.items.Get(i).SetQuality(in.items.Get(i).GetQuality() - 1)
						}
					}
				} else {
					in.items.Get(i).SetQuality(in.items.Get(i).GetQuality() - in.items.Get(i).GetQuality())
				}
			} else {
				if (in.items.Get(i).GetQuality() < 50) {
					in.items.Get(i).SetQuality(in.items.Get(i).GetQuality() + 1)
				}
			}
		}
	}

}

func main() {
	fmt.Println("OMGHAI!")
	NewInn().UpdateQuality()
}
