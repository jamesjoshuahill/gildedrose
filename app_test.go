package gildedrose_test

import (
	"github.com/jamesjoshuahill/gildedrose"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("App", func() {
	const initialQuality = 10

	It("updates quality of all items", func() {
		items := []gildedrose.Item{
			gildedrose.NewItem("some normal item", 1, initialQuality),
			gildedrose.NewItem("Aged Brie", 0, initialQuality),
		}
		app := gildedrose.NewApp(items)

		app.UpdateQuality()

		Expect(app.Items[0].SellIn()).To(Equal(0))
		Expect(app.Items[0].Quality()).To(Equal(initialQuality - 1))
		Expect(app.Items[1].SellIn()).To(Equal(-1))
		Expect(app.Items[1].Quality()).To(Equal(initialQuality + 2))
	})
})
