package inventory

type inventory struct {
	items []Item
}

// New returns a new inventory with a list of items.
func New(items []Item) *inventory {
	return &inventory{items: items}
}

// List returns a string with a line for each item including it's name, sell in, and quality.
func (i *inventory) List() []Item {
	return i.items
}
