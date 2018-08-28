package gildedrose_test

import (
	"github.com/jamesjoshuahill/gildedrose"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("updateableItem", func() {
	const initialQuality = 10

	DescribeTable("updates quality of normal items",
		func(sellIn, quality, expectedQuality int) {
			item := gildedrose.NewUpdateableItem(gildedrose.Item{Name: "some normal item", SellIn: sellIn, Quality: quality})

			updated := item.Update()

			Expect(updated.SellIn).To(Equal(sellIn - 1))
			Expect(updated.Quality).To(Equal(expectedQuality))
		},
		Entry("before sell by date", 1, initialQuality, initialQuality-1),
		Entry("on sell by date", 0, initialQuality, initialQuality-2),
		Entry("after sell by date", -1, initialQuality, initialQuality-2),
		Entry("of zero quality", 1, 0, 0),
	)

	DescribeTable("updates quality of aged brie",
		func(sellIn, quality, expectedQuality int) {
			item := gildedrose.NewUpdateableItem(gildedrose.Item{Name: "Aged Brie", SellIn: sellIn, Quality: quality})

			updated := item.Update()

			Expect(updated.SellIn).To(Equal(sellIn - 1))
			Expect(updated.Quality).To(Equal(expectedQuality))
		},
		Entry("before sell by date", 1, initialQuality, initialQuality+1),
		Entry("on sell by date", 0, initialQuality, initialQuality+2),
		Entry("after sell by date", -1, initialQuality, initialQuality+2),
		Entry("of quality 50", 1, 50, 50),
	)

	DescribeTable("updates quality of sulfuras",
		func(sellIn, quality, expectedQuality int) {
			item := gildedrose.NewUpdateableItem(gildedrose.Item{Name: "Sulfuras, Hand of Ragnaros", SellIn: sellIn, Quality: quality})

			updated := item.Update()

			Expect(updated.SellIn).To(Equal(sellIn))
			Expect(updated.Quality).To(Equal(expectedQuality))
		},
		Entry("before sell by date", 1, 80, 80),
		Entry("on sell by date", 0, 80, 80),
		Entry("after sell by date", -1, 80, 80),
	)

	DescribeTable("updates quality of backstage passes",
		func(sellIn, quality, expectedQuality int) {
			item := gildedrose.NewUpdateableItem(gildedrose.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: sellIn, Quality: quality})

			updated := item.Update()

			Expect(updated.SellIn).To(Equal(sellIn - 1))
			Expect(updated.Quality).To(Equal(expectedQuality))
		},
		Entry("11 days before sell by date", 11, initialQuality, initialQuality+1),
		Entry("10 days before sell by date", initialQuality, initialQuality, initialQuality+2),
		Entry("6 days before sell by date", 6, initialQuality, initialQuality+2),
		Entry("5 days before sell by date", 5, initialQuality, initialQuality+3),
		Entry("1 day before sell by date", 1, initialQuality, initialQuality+3),
		Entry("1 day before sell by date of quality 50", 1, 50, 50),
		Entry("on sell by date", 0, initialQuality, 0),
		Entry("of zero quality", 0, 0, 0),
	)

	DescribeTable("updates quality of conjured items",
		func(sellIn, quality, expectedQuality int) {
			item := gildedrose.NewUpdateableItem(gildedrose.Item{Name: "Conjured Mana Cake", SellIn: sellIn, Quality: quality})

			updated := item.Update()

			Expect(updated.SellIn).To(Equal(sellIn - 1))
			Expect(updated.Quality).To(Equal(expectedQuality))
		},
		Entry("before sell by date", 1, initialQuality, initialQuality-2),
		Entry("on sell by date", 0, initialQuality, initialQuality-4),
		Entry("after sell by date", -1, initialQuality, initialQuality-4),
		Entry("of zero quality", 1, 0, 0),
	)
})
