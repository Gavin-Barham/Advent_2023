package main

func problemOne(filePath string) int {
	input := getInputFromFile(filePath)
	matrix := formatInputTo2DArray(input)
	expandedInput := expandUniverse(matrix)
	galaxyVerticesArr := findAllGalaxies(expandedInput)
	sum := 0
	for i := 0; i < len(galaxyVerticesArr)-1; i++ {
		for j := i + 1; j < len(galaxyVerticesArr); j++ {
			x1 := galaxyVerticesArr[i]["x"]
			y1 := galaxyVerticesArr[i]["y"]

			x2 := galaxyVerticesArr[j]["x"]
			y2 := galaxyVerticesArr[j]["y"]

			var x int
			if x2 > x1 {
				x = x2 - x1
			} else {
				x = x1 - x2
			}

			var y int
			if y2 > y1 {
				y = y2 - y1
			} else {
				y = y1 - y2
			}

			sum += x + y
		}
	}
	return sum
}
