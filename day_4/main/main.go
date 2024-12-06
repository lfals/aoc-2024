package main

import (
	ceressearch "day_4/Ceres_Search"
	"fmt"
)

func main() {
	totalXmas := ceressearch.CountXmas("../input.txt")
	fmt.Println("Part one:", totalXmas)
}
