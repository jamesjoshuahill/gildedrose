package inventory

const (
	agedBrie        = "Aged Brie"
	backstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	sulfuras        = "Sulfuras, Hand of Ragnaros"
)

type Item interface {
	Name() string
	SellIn() int
	Quality() int
	Update() Item
}

type MagicItem struct {
	name            string
	sellIn, quality int
}

func (m MagicItem) Name() string {
	return m.name
}

func (m MagicItem) SellIn() int {
	return m.sellIn
}

func (m MagicItem) Quality() int {
	return m.quality
}

func (m MagicItem) Update() Item {
	var change int
	if m.sellIn < 1 {
		change = -2
	} else {
		change = -1
	}
	updatedQuality := normaliseQuality(m.quality, change)

	return MagicItem{name: m.name, sellIn: m.sellIn - 1, quality: updatedQuality}
}

type Sulfuras struct {
	sellIn, quality int
}

func (s Sulfuras) Name() string {
	return sulfuras
}

func (s Sulfuras) SellIn() int {
	return s.sellIn
}

func (s Sulfuras) Quality() int {
	return s.quality
}

func (s Sulfuras) Update() Item {
	return s
}

type AgedBrie struct {
	sellIn, quality int
}

func (a AgedBrie) Name() string {
	return agedBrie
}

func (a AgedBrie) SellIn() int {
	return a.sellIn
}

func (a AgedBrie) Quality() int {
	return a.quality
}

func (a AgedBrie) Update() Item {
	var change int
	if a.sellIn < 1 {
		change = 2
	} else {
		change = 1
	}
	updatedQuality := normaliseQuality(a.quality, change)

	return AgedBrie{sellIn: a.sellIn - 1, quality: updatedQuality}
}

type BackstagePasses struct {
	sellIn, quality int
}

func (b BackstagePasses) Name() string {
	return backstagePasses
}

func (b BackstagePasses) SellIn() int {
	return b.sellIn
}

func (b BackstagePasses) Quality() int {
	return b.quality
}

func (b BackstagePasses) Update() Item {
	var change int
	if b.sellIn > 10 {
		change = 1
	}
	if b.sellIn <= 10 && b.sellIn > 5 {
		change = 2
	}
	if b.sellIn <= 5 && b.sellIn > 0 {
		change = 3
	}
	if b.sellIn <= 0 {
		change = -b.quality
	}
	updatedQuality := normaliseQuality(b.quality, change)

	return BackstagePasses{sellIn: b.sellIn - 1, quality: updatedQuality}
}

func normaliseQuality(current int, change int) int {
	q := current + change

	if q < 0 {
		q = 0
	}
	if q > 50 {
		q = 50
	}

	return q
}
