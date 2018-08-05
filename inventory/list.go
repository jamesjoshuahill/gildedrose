package inventory

import "fmt"

func (i *inventory) List() (l string) {
	for _, item := range i.items {
		l += fmt.Sprintf("%s, %d, %d\n", item.name, item.sellIn, item.quality)
	}
	return
}
