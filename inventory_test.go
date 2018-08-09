package gildedrose_test

import (
	"github.com/jamesjoshuahill/gildedrose"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Inventory", func() {
	It("lists items", func() {
		items := []gildedrose.Item{gildedrose.NewItem("some-item", 1, 1)}
		i := gildedrose.New(items)

		list := i.List()

		Expect(list).To(Equal(items))
	})

	It("updates items", func() {
		items := []gildedrose.Item{gildedrose.NewItem("some-item", 1, 1)}
		i := gildedrose.New(items)

		i.Update()

		Expect(i.List()).To(HaveLen(1))
		Expect(i.List()[0].SellIn()).To(Equal(0))
		Expect(i.List()[0].Quality()).To(Equal(0))
	})
})
