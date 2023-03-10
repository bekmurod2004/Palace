package main

import "fmt"

func TwoSumm() {
	arr := []int{0,2,7,8}
	help := 0

	tar := 10

	ans := []int{}

	for i := 0; i < len(arr); i++ {
		// fmt.Println(arr[i])
		for j := 0; j < len(arr); j++ {
			if arr[i] !=arr[j] {
				help = arr[i] +arr[j]
				if  tar ==help{
					ans = append(ans,i,j)
					fmt.Println(ans)
					return
					
				}

			
				
				
			}
			
		}
		fmt.Println()
		
	}

	
	
}