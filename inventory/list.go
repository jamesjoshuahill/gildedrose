package inventory

import "fmt"

// List returns a string with a line for each item including it's name, sell in, and quality.
func (i *inventory) List() (l string) {
	for _, item := range i.items {
		l += fmt.Sprintf("%s, %d, %d\n", item.name, item.sellIn, item.quality)
	}
	return
}
