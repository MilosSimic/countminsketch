package main

import (
	"fmt"
)

func main() {
	cms := New(4, 7)
	stream := []string{"A", "A", "B", "C", "C", "B", "D", "A"}
	for _, ltr := range stream {
		cms.Add(ltr)
	}

	fmt.Println("Frequency of A:", cms.Query("A"))
}
