package main

func replaceAndRemove(arr []string) []string {
	writeIndex := 0
	aCount := 0

	for _, c := range arr {
		if c == "-" {
			break
		}
		if c != "b" {
			arr[writeIndex] = c
			writeIndex++
		}

		if c == "a" {
			aCount++
		}
	}

	finalLength := writeIndex + aCount
	begin := writeIndex - 1
	writeIndex = finalLength - 1

	for begin >= 0 {
		if arr[begin] == "a" {
			arr[writeIndex] = "d"
			arr[writeIndex-1] = "d"
			writeIndex -= 2
		} else {
			arr[writeIndex] = arr[begin]
			writeIndex--
		}
		begin--
	}

	return arr
}
