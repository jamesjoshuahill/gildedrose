package inventory_test

import (
	"github.com/jamesjoshuahill/gildedrose/inventory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ItemBuilder", func() {
	It("builds standard items", func() {
		b := inventory.ItemBuilder{}

		item := b.Build("some-name", 1, 2)

		Expect(item).To(BeAssignableToTypeOf(&inventory.MagicItem{}))
		Expect(item.Name()).To(Equal("some-name"))
		Expect(item.SellIn()).To(Equal(1))
		Expect(item.Quality()).To(Equal(2))
	})

	It("builds Sulfuras", func() {
		b := inventory.ItemBuilder{}

		item := b.Build("Sulfuras, Hand of Ragnaros", 10, 80)

		Expect(item).To(BeAssignableToTypeOf(inventory.Sulfuras{}))
		Expect(item.Name()).To(Equal("Sulfuras, Hand of Ragnaros"))
		Expect(item.SellIn()).To(Equal(10))
		Expect(item.Quality()).To(Equal(80))
	})

	It("builds AgedBrie", func() {
		b := inventory.ItemBuilder{}

		item := b.Build("Aged Brie", 2, 0)

		Expect(item).To(BeAssignableToTypeOf(&inventory.AgedBrie{}))
		Expect(item.Name()).To(Equal("Aged Brie"))
		Expect(item.SellIn()).To(Equal(2))
		Expect(item.Quality()).To(Equal(0))
	})

	It("builds Backstage Passes", func() {
		b := inventory.ItemBuilder{}

		item := b.Build("Backstage passes to a TAFKAL80ETC concert", 15, 20)

		Expect(item).To(BeAssignableToTypeOf(&inventory.BackstagePasses{}))
		Expect(item.Name()).To(Equal("Backstage passes to a TAFKAL80ETC concert"))
		Expect(item.SellIn()).To(Equal(15))
		Expect(item.Quality()).To(Equal(20))
	})
})
