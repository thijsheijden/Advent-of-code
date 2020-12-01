package main

import (
	"fmt"
	"log"
)

var expenses [200]int

func main() {
	for i := 0; i < 200; i++ {
		var j int
		_, err := fmt.Scanf("%d", &j)
		if err != nil {
			log.Fatal(err)
		} else {
			expenses[i] = j
		}
	}

	for i := 0; i < 200; i++ {
		for j := 0; j < 200; j++ {
			for z := 0; z < 200; z++ {
				if expenses[i]+expenses[j]+expenses[z] == 2020 {
					println("Found solution:")
					println(expenses[i] * expenses[j] * expenses[z])
				}
			}

		}
	}
}
