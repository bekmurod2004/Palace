package main

import "fmt"



func main() {
	// TwoSumm()
	var allTicks  int = 50
	var zakaz int 
	action := true

	allPlaces := 0

	var name string
	var last string
	var email string
	

	sum := []string{}

	for action {
		var sub = []string{}
		

		fmt.Scanln(&name)
		fmt.Scanln(&last)
		fmt.Scanln(&email)
		

		sub = append(sub, name, last, email)
		sum = append(sum, sub...)
		



		fmt.Print("nechta joy :")
		fmt.Scanln(&zakaz)
	
		
		if zakaz > 0 {
			allPlaces += zakaz
			}


	
			if allPlaces >= 50 {
				fmt.Println("joy qomadi")
				// fmt.Println(sum)
				for i := 0; i < len(sum);  i+= 3 {

					fmt.Println(sum[i])
					fmt.Println(sum[i + 1])
					fmt.Println(sum[i + 2])
					fmt.Println()
					
				}
				return
			}


		allTicks -= zakaz
		fmt.Println("qogan joy : ",allTicks)


		if zakaz == 101 {
			action = false
		}

	
	}
	

	fmt.Println(allPlaces)

	

}