package main

import "fmt"

func characterFrequencyInText(s string) string {
	res := ""
	charArr := make([]int, 128)

	for _, r := range s {
		index := r - ' '
		fmt.Printf("%d\n", index)
		charArr[index] += 1
	}

	fmt.Println(charArr)

	for i, count := range charArr {
		if count > 0 {
			fmt.Printf("%c - %d \n", i+' ', count)
		}
	}

	return res
}
