package gildedrose_test

import (
	"github.com/jamesjoshuahill/gildedrose"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App", func() {
	const initialQuality = 10

	It("updates quality of all items", func() {
		items := []gildedrose.Item{
			{Name: "some normal item", SellIn: 1, Quality: initialQuality},
			{Name: "Aged Brie", SellIn: 0, Quality: initialQuality},
		}
		app := gildedrose.NewApp(items)

		app.UpdateQuality()

		Expect(items).To(ConsistOf(
			gildedrose.Item{Name: "some normal item", SellIn: 0, Quality: initialQuality - 1},
			gildedrose.Item{Name: "Aged Brie", SellIn: -1, Quality: initialQuality + 2},
		))
	})
})
