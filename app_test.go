package gildedrose_test

import (
	"github.com/jamesjoshuahill/gildedrose"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App", func() {
	It("updates the quality of normal items", func() {
		items := []gildedrose.Item{gildedrose.NewItem("some-item", 1, 1)}
		app := gildedrose.NewApp(items)

		app.UpdateQuality()

		Expect(app.Items).To(HaveLen(1))
		Expect(app.Items[0].SellIn()).To(Equal(0))
		Expect(app.Items[0].Quality()).To(Equal(0))
	})
})
