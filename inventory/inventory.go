package inventory

type inventory struct {
	items []Item
}

// New returns a new inventory with a list of items.
func New(items []Item) *inventory {
	return &inventory{items: items}
}
