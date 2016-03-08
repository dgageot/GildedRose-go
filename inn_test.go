package main

import (
	"testing"

	"github.com/dgageot/gildedrose/legacy"
)

func Test_GildedRose(t *testing.T) {
	inn := NewInn()
	legacyInn := legacy.NewInn()

	for day := 0; day < 1000; day++ {
		for i := 0; i < len(inn.Items); i++ {
			if len(inn.Items) != legacyInn.Items.Size() {
				t.Fatalf("Expected item count %d. Got %d", legacyInn.Items.Size(), len(inn.Items))
			}

			item := inn.Items[i]
			legacyItem := legacyInn.Items.Get(i)

			if item.Name != legacyItem.GetName() {
				t.Errorf("Expected name %s. Got %s", legacyItem.GetName(), item.Name)
			}
			if item.Quality != legacyItem.GetQuality() {
				t.Errorf("Expected quality %d. Got %d", legacyItem.GetQuality(), item.Quality)
			}
			if item.SellIn != legacyItem.GetSellIn() {
				t.Errorf("Expected sellIn %d. Got %d", legacyItem.GetSellIn(), item.SellIn)
			}
		}

		inn.UpdateQuality()
		legacyInn.UpdateQuality()
	}
}
