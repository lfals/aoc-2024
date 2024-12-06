package main

import (
	ceressearch "day_4/Ceres_Search"
	part_two "day_4/part_two"
	"fmt"
)

func main() {
	totalXmas := ceressearch.CountXmas("../input.txt")
	fmt.Println("Part one:", totalXmas)

	totalX_mas := part_two.CalculateX_MAS("../input.txt")
	fmt.Println("Part one:", totalX_mas)
}
