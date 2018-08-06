package inventory_test

import (
	"github.com/jamesjoshuahill/gildedrose/inventory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Item", func() {
	var builder inventory.ItemBuilder

	Context("when a magic item is updated", func() {
		It("reduces sell in by one", func() {
			i := builder.Build("", 1, 0)

			i.Update()

			Expect(i.SellIn()).To(Equal(0))
		})
	})

	Context("when Sulfuras is updated", func() {
		It("does not reduce sell in", func() {
			i := builder.Build("Sulfuras, Hand of Ragnaros", 1, 80)

			i.Update()

			Expect(i.SellIn()).To(Equal(1))
		})

		It("does not reduce in quality", func() {
			i := builder.Build("Sulfuras, Hand of Ragnaros", 1, 80)

			i.Update()

			Expect(i.Quality()).To(Equal(80))
		})
	})
})
