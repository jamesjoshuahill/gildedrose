package inventory_test

import (
	"github.com/jamesjoshuahill/gildedrose/inventory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Update", func() {
	Context("when there is a standard item", func() {
		It("reduces sell in by one", func() {
			i := inventory.New([]inventory.Item{{SellIn: 1}})

			i.Update()

			Expect(i.List()[0].SellIn).To(Equal(0))
		})
	})

	Context("when the item is Sulfuras", func() {
		It("does not reduce sell in", func() {
			i := inventory.New([]inventory.Item{{Name: "Sulfuras, Hand of Ragnaros", SellIn: 1}})

			i.Update()

			Expect(i.List()[0].SellIn).To(Equal(1))
		})

		It("does not reduce in quality", func() {
			i := inventory.New([]inventory.Item{{Name: "Sulfuras, Hand of Ragnaros", Quality: 80}})

			i.Update()

			Expect(i.List()[0].Quality).To(Equal(80))
		})
	})
})
