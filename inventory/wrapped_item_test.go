package inventory_test

import (
	"github.com/jamesjoshuahill/gildedrose/inventory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Wrapped Item", func() {
	Context("when a standard item is updated", func() {
		It("reduces sell in by one", func() {
			w := inventory.WrappedItem{Item: inventory.Item{SellIn: 1}}

			updated := w.Update()

			Expect(updated.SellIn).To(Equal(0))
		})
	})

	Context("when Sulfuras is updated", func() {
		It("does not reduce sell in", func() {
			w := inventory.WrappedItem{Item: inventory.Item{Name: "Sulfuras, Hand of Ragnaros", SellIn: 1}}

			updated := w.Update()

			Expect(updated.SellIn).To(Equal(1))
		})

		It("does not reduce in quality", func() {
			w := inventory.WrappedItem{Item: inventory.Item{Name: "Sulfuras, Hand of Ragnaros", Quality: 80}}

			updated := w.Update()

			Expect(updated.Quality).To(Equal(80))
		})
	})
})
