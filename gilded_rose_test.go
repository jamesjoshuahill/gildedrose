package gildedrose_test

import (
	"github.com/jamesjoshuahill/gildedrose"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Inventory", func() {
	It("updates items", func() {
		items := []gildedrose.Item{{Name: "some-item", SellIn: 1, Quality: 1}}
		app := gildedrose.New(items)

		app.UpdateQuality()

		Expect(app.Items).To(HaveLen(1))
		Expect(app.Items[0].SellIn()).To(Equal(0))
		Expect(app.Items[0].Quality()).To(Equal(0))
	})
})
