package inventory_test

import (
	"github.com/jamesjoshuahill/gildedrose/inventory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Inventory", func() {
	It("lists items", func() {
		items := []inventory.Item{inventory.NewItem("some-item", 1, 1)}
		i := inventory.New(items)

		list := i.List()

		Expect(list).To(Equal(items))
	})

	It("updates items", func() {
		items := []inventory.Item{inventory.NewItem("some-item", 1, 1)}
		i := inventory.New(items)

		i.Update()

		Expect(i.List()).To(HaveLen(1))
		Expect(i.List()[0].SellIn()).To(Equal(0))
		Expect(i.List()[0].Quality()).To(Equal(0))
	})
})
