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

func NewItem(name string, sellIn, quality int) Item {
	switch name {
	case agedBrie:
		return &AgedBrie{sellIn: NewSellIn(sellIn), quality: NewQuality(quality)}
	case backstagePasses:
		return &BackstagePasses{sellIn: NewSellIn(sellIn), quality: NewQuality(quality)}
	case sulfuras:
		return Sulfuras{sellIn: NewSellIn(sellIn), quality: NewQuality(quality)}
	default:
		return &MagicItem{
			name:    name,
			sellIn:  NewSellIn(sellIn),
			quality: NewQuality(quality),
		}
	}
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

type Quality struct {
	value int
}

func NewQuality(value int) *Quality {
	return &Quality{value: value}
}

func (q *Quality) Update(amount int) {
	v := q.value + amount

	if v < 0 {
		v = 0
	}
	if v > 50 {
		v = 50
	}

	q.value = v
}

type MagicItem struct {
	name    string
	sellIn  *SellIn
	quality *Quality
}

func (m MagicItem) Name() string {
	return m.name
}

func (m MagicItem) SellIn() int {
	return m.sellIn.value
}

func (m MagicItem) Quality() int {
	return m.quality.value
}

func (m *MagicItem) Update() {
	var change int
	if m.sellIn.value < 1 {
		change = -2
	} else {
		change = -1
	}

	m.sellIn.Tick()
	m.quality.Update(change)
}

type Sulfuras struct {
	sellIn  *SellIn
	quality *Quality
}

func (s Sulfuras) Name() string {
	return sulfuras
}

func (s Sulfuras) SellIn() int {
	return s.sellIn.value
}

func (s Sulfuras) Quality() int {
	return s.quality.value
}

func (s Sulfuras) Update() {
	return
}

type AgedBrie struct {
	sellIn  *SellIn
	quality *Quality
}

func (a AgedBrie) Name() string {
	return agedBrie
}

func (a AgedBrie) SellIn() int {
	return a.sellIn.value
}

func (a AgedBrie) Quality() int {
	return a.quality.value
}

func (a *AgedBrie) Update() {
	var change int
	if a.sellIn.value < 1 {
		change = 2
	} else {
		change = 1
	}

	a.sellIn.Tick()
	a.quality.Update(change)
}

type BackstagePasses struct {
	sellIn  *SellIn
	quality *Quality
}

func (b BackstagePasses) Name() string {
	return backstagePasses
}

func (b BackstagePasses) SellIn() int {
	return b.sellIn.value
}

func (b BackstagePasses) Quality() int {
	return b.quality.value
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
		change = -b.quality.value
	}

	b.sellIn.Tick()
	b.quality.Update(change)
}
