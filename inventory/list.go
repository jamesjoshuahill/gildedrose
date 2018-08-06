package inventory

// List returns a string with a line for each item including it's name, sell in, and quality.
func (i *inventory) List() []Item {
	return i.items
}
