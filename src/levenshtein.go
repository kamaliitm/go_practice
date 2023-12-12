package main

func findLevenshteinDistance(a string, b string) int32 {
	// aRunes := []rune(a)
	// bRunes := []rune(b)

	aLen := len(a)
	bLen := len(b)
	distanceBetweenIndexes := make([][]int32, aLen)
	for aIndex := range a {
		distanceBetweenIndexes[aIndex] = make([]int32, bLen)
		for bIndex := range b {
			distanceBetweenIndexes[aIndex][bIndex] = -1
		}
	}

	return calculateDistanceBetweenPrefixes(a, aLen-1, b, bLen-1, &distanceBetweenIndexes)
}

func calculateDistanceBetweenPrefixes(a string, aIndex int, b string, bIndex int, distanceBetweenIndexesPtr *[][]int32) int32 {
	if aIndex < 0 {
		return int32(bIndex) + 1 // add all characters in b
	} else if bIndex < 0 {
		return int32(aIndex) + 1 // delete all characters from a
	}

	distanceBetweenIndexes := *distanceBetweenIndexesPtr

	if distanceBetweenIndexes[aIndex][bIndex] == -1 {
		if a[aIndex] == b[bIndex] {
			distanceBetweenIndexes[aIndex][bIndex] = calculateDistanceBetweenPrefixes(
				a, aIndex-1, b, bIndex-1, distanceBetweenIndexesPtr,
			)
		} else {
			addCharDistance := calculateDistanceBetweenPrefixes(a, aIndex, b, bIndex-1, distanceBetweenIndexesPtr)
			deleteCharDistance := calculateDistanceBetweenPrefixes(a, aIndex-1, b, bIndex, distanceBetweenIndexesPtr)
			editCharDistance := calculateDistanceBetweenPrefixes(a, aIndex-1, b, bIndex-1, distanceBetweenIndexesPtr)

			minDistance := addCharDistance
			if deleteCharDistance < minDistance {
				minDistance = deleteCharDistance
			}
			if editCharDistance < minDistance {
				minDistance = editCharDistance
			}

			distanceBetweenIndexes[aIndex][bIndex] = 1 + minDistance
		}
	}

	return distanceBetweenIndexes[aIndex][bIndex]

}
