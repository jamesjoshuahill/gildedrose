package gildedrose_test

import (
	"github.com/jamesjoshuahill/gildedrose"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Gilded Rose", func() {
	It("foo", func() {
		items := []gildedrose.Item{{Name: "foo", SellIn: 0, Quality: 0}}
		app := gildedrose.New(items)

		app.UpdateQuality()

		Expect(items[0].Name).To(Equal("fixme"))
	})
})
