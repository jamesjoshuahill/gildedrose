package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jamesjoshuahill/gildedrose/inventory"
)

func main() {
	daysArg := os.Args[1]
	if daysArg == "" {
		daysArg = "0"
	}

	days, err := strconv.Atoi(daysArg)
	if err != nil {
		log.Fatalln(err)
	}

	for day := 0; day <= days; day++ {
		if day == 0 {
			fmt.Println("OMGHAI!")
		} else {
			fmt.Println()
		}
		fmt.Printf("-------- day %d --------\n", day)
		fmt.Printf("name, sellIn, quality\n")
		inventory.Print()
		inventory.Update()
	}
}
