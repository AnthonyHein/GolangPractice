package main

import (
	"fmt"
	"UdemyCourse/LVL13_E2/quote"
    "UdemyCourse/LVL13_E2/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
