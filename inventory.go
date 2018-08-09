package gildedrose

type inventory struct {
	list []Item
}

// New returns a new inventory with a list of items.
func New(list []Item) *inventory {
	return &inventory{list: list}
}

// List returns the items in the inventory.
func (i *inventory) List() []Item {
	return i.list
}

// Update changes the list of items in the inventory to reflect the passing of one day.
func (i *inventory) Update() {
	for _, item := range i.list {
		item.Update()
	}
}
