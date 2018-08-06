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
	Update()
}

type SellIn struct {
	value int
}

func NewSellIn(value int) *SellIn {
	return &SellIn{value: value}
}

func (s *SellIn) Tick() {
	s.value = s.value - 1
}

type MagicItem struct {
	name    string
	sellIn  *SellIn
	quality int
}

func (m MagicItem) Name() string {
	return m.name
}

func (m MagicItem) SellIn() int {
	return m.sellIn.value
}

func (m MagicItem) Quality() int {
	return m.quality
}

func (m *MagicItem) Update() {
	var change int
	if m.sellIn.value < 1 {
		change = -2
	} else {
		change = -1
	}

	m.sellIn.Tick()
	m.quality = normaliseQuality(m.quality, change)
}

type Sulfuras struct {
	sellIn  *SellIn
	quality int
}

func (s Sulfuras) Name() string {
	return sulfuras
}

func (s Sulfuras) SellIn() int {
	return s.sellIn.value
}

func (s Sulfuras) Quality() int {
	return s.quality
}

func (s Sulfuras) Update() {
	return
}

type AgedBrie struct {
	sellIn  *SellIn
	quality int
}

func (a AgedBrie) Name() string {
	return agedBrie
}

func (a AgedBrie) SellIn() int {
	return a.sellIn.value
}

func (a AgedBrie) Quality() int {
	return a.quality
}

func (a *AgedBrie) Update() {
	var change int
	if a.sellIn.value < 1 {
		change = 2
	} else {
		change = 1
	}

	a.sellIn.Tick()
	a.quality = normaliseQuality(a.quality, change)
}

type BackstagePasses struct {
	sellIn  *SellIn
	quality int
}

func (b BackstagePasses) Name() string {
	return backstagePasses
}

func (b BackstagePasses) SellIn() int {
	return b.sellIn.value
}

func (b BackstagePasses) Quality() int {
	return b.quality
}

func (b *BackstagePasses) Update() {
	var change int
	if b.sellIn.value > 10 {
		change = 1
	}
	if b.sellIn.value <= 10 && b.sellIn.value > 5 {
		change = 2
	}
	if b.sellIn.value <= 5 && b.sellIn.value > 0 {
		change = 3
	}
	if b.sellIn.value <= 0 {
		change = -b.quality
	}

	b.sellIn.Tick()
	b.quality = normaliseQuality(b.quality, change)
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
