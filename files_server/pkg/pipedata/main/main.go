package main

import "fmt"

func CountValues(list []string) map[string]int {

	duplicate_frequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	return duplicate_frequency
}

func main() {
	duplicate := []string{"Hello", "World", "GoodBye", "World", "We", "Love", "Love", "You"}
	c := CountValues(duplicate)

	fmt.Println("Items: ", c)

}
