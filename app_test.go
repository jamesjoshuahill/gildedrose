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

		Expect(app.Items).To(ConsistOf(
			gildedrose.NewItem("some normal item", 0, initialQuality-1),
			gildedrose.NewItem("Aged Brie", -1, initialQuality+2),
		))
	})
})
