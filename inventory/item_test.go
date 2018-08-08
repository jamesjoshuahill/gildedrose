package inventory_test

import (
	"github.com/jamesjoshuahill/gildedrose/inventory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Item", func() {
	It("creates normal items", func() {
		item := inventory.NewItem("some-name", 1, 2)

		Expect(item.Name()).To(Equal("some-name"))
		Expect(item.SellIn()).To(Equal(1))
		Expect(item.Quality()).To(Equal(2))
	})

	It("creates Sulfuras", func() {
		item := inventory.NewItem("Sulfuras, Hand of Ragnaros", 10, 80)

		Expect(item.Name()).To(Equal("Sulfuras, Hand of Ragnaros"))
		Expect(item.SellIn()).To(Equal(10))
		Expect(item.Quality()).To(Equal(80))
	})

	It("creates AgedBrie", func() {
		item := inventory.NewItem("Aged Brie", 2, 0)

		Expect(item.Name()).To(Equal("Aged Brie"))
		Expect(item.SellIn()).To(Equal(2))
		Expect(item.Quality()).To(Equal(0))
	})

	It("creates Backstage Passes", func() {
		item := inventory.NewItem("Backstage passes to a TAFKAL80ETC concert", 15, 20)

		Expect(item.Name()).To(Equal("Backstage passes to a TAFKAL80ETC concert"))
		Expect(item.SellIn()).To(Equal(15))
		Expect(item.Quality()).To(Equal(20))
	})

	It("creates Conjured items", func() {
		item := inventory.NewItem("Conjured Mana Cake", 3, 6)

		Expect(item.Name()).To(Equal("Conjured Mana Cake"))
		Expect(item.SellIn()).To(Equal(3))
		Expect(item.Quality()).To(Equal(6))
	})

	Context("when a magic item is updated", func() {
		It("reduces sell in by one", func() {
			i := inventory.NewItem("", 1, 0)

			i.Update()

			Expect(i.SellIn()).To(Equal(0))
		})
	})

	Context("when Sulfuras is updated", func() {
		It("does not reduce sell in", func() {
			i := inventory.NewItem("Sulfuras, Hand of Ragnaros", 1, 80)

			i.Update()

			Expect(i.SellIn()).To(Equal(1))
		})

		It("does not reduce in quality", func() {
			i := inventory.NewItem("Sulfuras, Hand of Ragnaros", 1, 80)

			i.Update()

			Expect(i.Quality()).To(Equal(80))
		})
	})
})
