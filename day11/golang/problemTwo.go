package main

func problemTwo(filepath string) int {
	input := getInputFromFile(filepath)
	matrix := formatInputTo2DArray(input)
	emptyColumnsIndexArr := getEmptySpaceColumns(matrix)
	emptyRowsIndexArr := getEmptySpaceRows(matrix)
	galaxyVerticesArr := findAllGalaxies(matrix)
	sum := 0
	for i := 0; i < len(galaxyVerticesArr)-1; i++ {
		for j := i + 1; j < len(galaxyVerticesArr); j++ {
			x1 := galaxyVerticesArr[i]["x"]
			y1 := galaxyVerticesArr[i]["y"]

			x2 := galaxyVerticesArr[j]["x"]
			y2 := galaxyVerticesArr[j]["y"]

			x := calculateSumOfTwoPointsOnSameAxesWithExpansion(x1, x2, emptyRowsIndexArr)

			y := calculateSumOfTwoPointsOnSameAxesWithExpansion(y1, y2, emptyColumnsIndexArr)

			sum += x + y
		}
	}
	return sum
}
func calculateSumOfTwoPointsOnSameAxesWithExpansion(num1 int, num2 int, emptyIndexArr []int) int {
	var sum int
	if num2 > num1 {
		for k := num1; k < num2; k++ {
			if contains(emptyIndexArr, k) {
				sum += 1000000
			} else {
				sum += 1
			}
		}
	} else {
		for k := num2; k < num1; k++ {
			if contains(emptyIndexArr, k) {
				sum += 1000000
			} else {
				sum += 1
			}
		}
	}
	return sum
}
