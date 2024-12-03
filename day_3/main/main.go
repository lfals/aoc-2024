package main

import (
	mullitover "day_3/Mull_It_Over"
	parttwo "day_3/part_two"
	"fmt"
)

func main() {
	uncorruptedMul := mullitover.ScanForUncorruptedInstuctions("../input.txt")
	fmt.Println("Uncorrupted Mul:", uncorruptedMul)

	uncorruptedMulDo := parttwo.CalculateEnabledMultiplications("../input.txt")
	fmt.Println("Do Mul", uncorruptedMulDo)

}
