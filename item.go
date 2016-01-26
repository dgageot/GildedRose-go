package main

type Item struct {
	name    string
	sellIn  int
	quality int
}

func (i *Item) GetName() string {
	return i.name
}

func (i *Item) GetSellIn() int {
	return i.sellIn
}

func (i *Item) GetQuality() int {
	return i.quality
}

func (i *Item) SetSellIn(sellIn int) {
	i.sellIn = sellIn
}

func (i *Item) SetQuality(quality int) {
	i.quality = quality
}

type Items struct {
	list []*Item
}

func NewItems() *Items {
	return &Items{
		list: []*Item{},
	}
}

func (is *Items) Add(item *Item) {
	is.list = append(is.list, item)
}

func (is *Items) Get(index int) *Item {
	return is.list[index]
}

func (is *Items) Size() int {
	return len(is.list)
}
