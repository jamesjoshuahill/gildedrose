package gildedrose_test

import (
	"github.com/jamesjoshuahill/gildedrose"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Inventory", func() {
	It("updates items", func() {
		items := []gildedrose.MagicItem{gildedrose.NewItem("some-item", 1, 1)}
		i := gildedrose.New(items)

		i.Update()

		Expect(items).To(HaveLen(1))
		Expect(items[0].SellIn()).To(Equal(0))
		Expect(items[0].Quality()).To(Equal(0))
	})
})
