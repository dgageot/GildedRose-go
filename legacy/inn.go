package legacy

import "fmt"

type Inn struct {
	Items *Items
}

func NewInn() *Inn {
	inn := &Inn{
		Items: NewItems(),
	}
	inn.Items.Add(&Item{"+5 Dexterity Vest", 10, 20})
	inn.Items.Add(&Item{"Aged Brie", 2, 0})
	inn.Items.Add(&Item{"Elixir of the Mongoose", 5, 7})
	inn.Items.Add(&Item{"Sulfuras, Hand of Ragnaros", 0, 80})
	inn.Items.Add(&Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20})
	inn.Items.Add(&Item{"Conjured Mana Cake", 3, 6})

	return inn
}

func (in *Inn) UpdateQuality() {
	for i := 0; i < in.Items.Size(); i++ {
		if !(in.Items.Get(i).GetName() == "Aged Brie") && !(in.Items.Get(i).GetName() == "Backstage passes to a TAFKAL80ETC concert") {
			if (in.Items.Get(i).GetQuality() > 0) {
				if !(in.Items.Get(i).GetName() == "Sulfuras, Hand of Ragnaros") {
					in.Items.Get(i).SetQuality(in.Items.Get(i).GetQuality() - 1)
				}
			}
		} else {
			if (in.Items.Get(i).GetQuality() < 50) {
				in.Items.Get(i).SetQuality(in.Items.Get(i).GetQuality() + 1)

				if (in.Items.Get(i).GetName() == "Backstage passes to a TAFKAL80ETC concert") {
					if (in.Items.Get(i).GetSellIn() < 11) {
						if (in.Items.Get(i).GetQuality() < 50) {
							in.Items.Get(i).SetQuality(in.Items.Get(i).GetQuality() + 1)
						}
					}

					if (in.Items.Get(i).GetSellIn()) < 6 {
						if (in.Items.Get(i).GetQuality() < 50) {
							in.Items.Get(i).SetQuality(in.Items.Get(i).GetQuality() + 1)
						}
					}
				}
			}
		}

		if !(in.Items.Get(i).GetName() == "Sulfuras, Hand of Ragnaros") {
			in.Items.Get(i).SetSellIn(in.Items.Get(i).GetSellIn() - 1)
		}

		if (in.Items.Get(i).GetSellIn() < 0) {
			if !(in.Items.Get(i).GetName() == "Aged Brie") {
				if !(in.Items.Get(i).GetName() == "Backstage passes to a TAFKAL80ETC concert") {
					if (in.Items.Get(i).GetQuality() > 0) {
						if !(in.Items.Get(i).GetName() == "Sulfuras, Hand of Ragnaros") {
							in.Items.Get(i).SetQuality(in.Items.Get(i).GetQuality() - 1)
						}
					}
				} else {
					in.Items.Get(i).SetQuality(in.Items.Get(i).GetQuality() - in.Items.Get(i).GetQuality())
				}
			} else {
				if (in.Items.Get(i).GetQuality() < 50) {
					in.Items.Get(i).SetQuality(in.Items.Get(i).GetQuality() + 1)
				}
			}
		}
	}

}

func main() {
	fmt.Println("OMGHAI!")
	NewInn().UpdateQuality()
}
