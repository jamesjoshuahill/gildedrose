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

		Expect(item.Name()).To(Equal("some-name"))
		Expect(item.SellIn()).To(Equal(1))
		Expect(item.Quality()).To(Equal(2))
	})

	It("builds Sulfuras", func() {
		b := inventory.ItemBuilder{}

		item := b.Build("Sulfuras, Hand of Ragnaros", 10, 80)

		Expect(item.Name()).To(Equal("Sulfuras, Hand of Ragnaros"))
		Expect(item.SellIn()).To(Equal(10))
		Expect(item.Quality()).To(Equal(80))
		updated := item.Update()
		Expect(item).To(Equal(updated))
	})
})
