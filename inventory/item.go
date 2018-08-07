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

func (s SellIn) Value() int {
	return s.value
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

func (q Quality) Value() int {
	return q.value
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
	return m.sellIn.Value()
}

func (m MagicItem) Quality() int {
	return m.quality.Value()
}

func (m *MagicItem) Update() {
	var change int
	if m.sellIn.Value() < 1 {
		change = -2
	} else {
		change = -1
	}

	m.quality.Update(change)
	m.sellIn.Tick()
}

type Sulfuras struct {
	sellIn  *SellIn
	quality *Quality
}

func (s Sulfuras) Name() string {
	return sulfuras
}

func (s Sulfuras) SellIn() int {
	return s.sellIn.Value()
}

func (s Sulfuras) Quality() int {
	return s.quality.Value()
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
	return a.sellIn.Value()
}

func (a AgedBrie) Quality() int {
	return a.quality.Value()
}

func (a *AgedBrie) Update() {
	var change int
	if a.sellIn.Value() < 1 {
		change = 2
	} else {
		change = 1
	}

	a.quality.Update(change)
	a.sellIn.Tick()
}

type BackstagePasses struct {
	sellIn  *SellIn
	quality *Quality
}

func (b BackstagePasses) Name() string {
	return backstagePasses
}

func (b BackstagePasses) SellIn() int {
	return b.sellIn.Value()
}

func (b BackstagePasses) Quality() int {
	return b.quality.Value()
}

func (b *BackstagePasses) Update() {
	current := b.sellIn.Value()

	var change int
	if current > 10 {
		change = 1
	}
	if current <= 10 && current > 5 {
		change = 2
	}
	if current <= 5 && current > 0 {
		change = 3
	}
	if current <= 0 {
		change = -b.quality.Value()
	}

	b.quality.Update(change)
	b.sellIn.Tick()
}
