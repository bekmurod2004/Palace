package main

import (
	"fmt"

	"github.com/TwiN/go-color"
)

func RenderPlace(boom []int) {
	var all = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// boom := []int{1,50,12,34}

	for i := 0; i < len(all); i++ {
		if all[i]%10 == 0 {
			print(color.Ize(color.Green, all[i-1]))
			println()

		} else {
			if all[i] == 1 {
				print(" ")
			} else {
				print(color.Ize(color.Green, all[i-1]))
				print(" ")

			}
		}

		for j := 0; j < len(boom); j++ {

			if all[i] == boom[j] {
				if i == 0 {

					fmt.Print(" ")
				}

				if all[i]%10 == 0 {
					i++

					print(color.Ize(color.Red, all[i-1]))
					println()

				} else {

					i++
					print(color.Ize(color.Red, all[i-1]))
					fmt.Print(" ")

				}

			}

		}

	}
	fmt.Println()

	// frePl := len(all) - 1 - len(boom)

	// fmt.Println("50 ta joydan ",frePl,"joy qoldi")

}
