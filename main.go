package main

import (
	"fmt"
)

func main() {
	stream := []string{"A", "A", "B", "C", "C", "B", "D", "A"}

	cms := New(4, 7)
	for _, ltr := range stream {
		cms.Add(ltr)
	}
	fmt.Println("Frequency of A [no estimets]:", cms.Query("A"))

	ecms := NewWithEstiments(0.001, 1e-09)
	for _, ltr := range stream {
		ecms.Add(ltr)
	}
	fmt.Println("Frequency of A [with estimets]:", ecms.Query("A"))
}
