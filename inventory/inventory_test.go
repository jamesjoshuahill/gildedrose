package inventory_test

import (
	"github.com/jamesjoshuahill/gildedrose/inventory"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Inventory", func() {
	var builder inventory.ItemBuilder

	It("lists items", func() {
		items := []inventory.Item{builder.Build("some-item", 1, 1)}
		i := inventory.New(items)

		list := i.List()

		Expect(list).To(Equal(items))
	})

	It("updates items", func() {
		items := []inventory.Item{builder.Build("some-item", 1, 1)}
		i := inventory.New(items)

		i.Update()

		Expect(i.List()).To(ConsistOf(builder.Build("some-item", 0, 0)))
	})
})
