package inventory

type inventory struct {
	items []Item
}

// New returns a new inventory with the default list of items.
func New() *inventory {
	return &inventory{
		items: items,
	}
}
