package inventory_test

import (
	"github.com/jamesjoshuahill/gildedrose/inventory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Inventory", func() {
	It("lists items", func() {
		items := []inventory.Item{{Name: "some-item", SellIn: 1, Quality: 1}}
		i := inventory.New(items)

		list := i.List()

		Expect(list).To(Equal(items))
	})

	It("updates items", func() {
		items := []inventory.Item{{Name: "some-item", SellIn: 1, Quality: 1}}
		i := inventory.New(items)

		i.Update()

		Expect(i.List()).To(ConsistOf(inventory.Item{Name: "some-item", SellIn: 0, Quality: 0}))
	})
})
