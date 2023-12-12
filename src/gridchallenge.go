package main

/*
 * Complete the 'gridChallenge' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING_ARRAY grid as parameter.
 */

func gridChallenge(grid []string) string {
	// Write your code here
	rowSize := len(grid)
	if rowSize < 1 {
		return ""
	}

	for i, rowStr := range grid {
		sortString(&rowStr, 0, len(rowStr)-1)
		grid[i] = rowStr
	}

	colSize := len(grid[0])
	for j := 0; j < colSize; j++ {
		for i := 1; i < rowSize; i++ {
			if grid[i][j] < grid[i-1][j] {
				return "NO"
			}
		}
	}

	return "YES"
}

func sortString(s *string, left int, right int) {
	if left < right {
		pivotIndex := partitionAroundPivotV2(s, left, right)
		sortString(s, left, pivotIndex-1)
		sortString(s, pivotIndex+1, right)
	}
}

func partitionAroundPivotV2(s *string, left int, right int) int {
	greater := left
	runes := []rune(*s)
	pivot := runes[right]
	for left <= right-1 {
		if runes[left] < pivot {
			temp := runes[left]
			runes[left] = runes[greater]
			runes[greater] = temp
			greater += 1
		}
		left += 1
	}
	runes[right] = runes[greater]
	runes[greater] = pivot

	*s = string(runes)

	return greater
}

// func main() {
// 	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

// 	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
// 	checkError(err)

// 	defer stdout.Close()

// 	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

// 	tTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 	checkError(err)
// 	t := int32(tTemp)

// 	for tItr := 0; tItr < int(t); tItr++ {
// 		nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
// 		checkError(err)
// 		n := int32(nTemp)

// 		var grid []string

// 		for i := 0; i < int(n); i++ {
// 			gridItem := readLine(reader)
// 			grid = append(grid, gridItem)
// 		}

// 		result := gridChallenge(grid)

// 		fmt.Fprintf(writer, "%s\n", result)
// 	}

// 	writer.Flush()
// }

// func readLine(reader *bufio.Reader) string {
// 	str, _, err := reader.ReadLine()
// 	if err == io.EOF {
// 		return ""
// 	}

// 	return strings.TrimRight(string(str), "\r\n")
// }

// func checkError(err error) {
// 	if err != nil {
// 		panic(err)
// 	}
// }
