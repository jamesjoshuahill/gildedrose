package inventory

type inventory struct {
	items []WrappedItem
}

// New returns a new inventory with a list of items.
func New(items []Item) *inventory {
	var wrapped []WrappedItem
	for _, item := range items {
		wrapped = append(wrapped, WrappedItem{item})
	}
	return &inventory{items: wrapped}
}

// List returns a string with a line for each item including it's name, sell in, and quality.
func (i *inventory) List() []Item {
	var unwrapped []Item
	for _, wrapped := range i.items {
		unwrapped = append(unwrapped, wrapped.Item)
	}
	return unwrapped
}

type WrappedItem struct {
	Item
}

func (w *WrappedItem) Update() {

}
