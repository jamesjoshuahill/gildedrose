package inventory

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Update", func() {
	Context("when there is a standard item", func() {
		It("reduces sell in by one", func() {
			i := inventory{items: []Item{{sellIn: 1}}}

			i.Update()

			Expect(i.items[0].sellIn).To(Equal(0))
		})
	})

	Context("when the item is Sulfuras", func() {
		It("does not reduce sell in", func() {
			i := inventory{items: []Item{{name: "Sulfuras, Hand of Ragnaros", sellIn: 1}}}

			i.Update()

			Expect(i.items[0].sellIn).To(Equal(1))
		})

		It("does not reduce in quality", func() {
			i := inventory{items: []Item{{name: "Sulfuras, Hand of Ragnaros", quality: 80}}}

			i.Update()

			Expect(i.items[0].quality).To(Equal(80))
		})
	})
})
