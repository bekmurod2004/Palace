package main

import (
	"fmt"
	"strconv"
)



func main() {
	getUsers()
	
}
func getUsers(){
	
	 action := true

	 var halqil string

	 var (
		ism string
		fam string
		eml string
		nech string 
		qay string
	 )

	 allP := 20
	 zakazOlJoyla := []int{}

	 allUsersPaptis := []string{}
	 allUsersMaptis := []string{}

	 for action {
		var joy string  



	    shabUse := make(map[string]string)

		fmt.Print("Tanla odam 'qoshish' || 'korish' ")
		fmt.Scanln(&halqil)

		if halqil == "qoshish" {
			fmt.Print("Ismi: ")
			fmt.Scanln(&ism)
			fmt.Print("Familiyasi: ")
			fmt.Scanln(&fam)
			fmt.Print("Email: ")
			fmt.Scanln(&eml)
			fmt.Print("Nechta joy: ")
			fmt.Scanln(&nech)

			numN,_ := strconv.Atoi(nech)

			allP -= numN


			if allP <= 0 {
				action = false
			}

			for i := 0; i < numN; i++ {

				fmt.Print("Qaysi joyla: ")
			    fmt.Scanln(&qay)
				
				joy +=qay
				joy += " "

				numQ,_ := strconv.Atoi(qay)

				
				zakazOlJoyla = append(zakazOlJoyla, numQ)
				
			}


			shabUse["Ismi"] = ism
			shabUse["Familiya"] = fam
			shabUse["Email"] = eml
			shabUse["Nechta_Joy"] = nech
			shabUse["Qaysi_Joylar"] = joy
			

			// fmt.Println(shabUse)
			fmt.Println("------> 50 tadan ",allP,"joy qoldi <------")

			fmt.Println()
			fmt.Println("--------------------------------------------------------")
			RenderPlace(zakazOlJoyla)
			fmt.Println("--------------------------------------------------------")
			fmt.Println()
			// fmt.Println(zakazOlJoyla)
			// fmt.Println()

			for i, v := range shabUse {
				allUsersPaptis = append(allUsersPaptis, v)
				allUsersMaptis = append(allUsersMaptis, i)
			}


		}else if halqil == "korish"{
			
			fmt.Println()
			fmt.Println("--------------------------------------------------------")
			RenderPlace(zakazOlJoyla)
			fmt.Println("--------------------------------------------------------")
			fmt.Println()

			
			
		}




	 }

	 fmt.Println("         ----------- All Users -----------")

	 for i := 0; i < len(allUsersPaptis); i+= 5 {
		fmt.Println(allUsersMaptis[i]," :", allUsersPaptis[i])
		fmt.Println(allUsersMaptis[i+1]," :", allUsersPaptis[i+1])
		fmt.Println(allUsersMaptis[i+2]," :", allUsersPaptis[i+2])
		fmt.Println(allUsersMaptis[i+3]," :", allUsersPaptis[i+3])
		fmt.Println(allUsersMaptis[i+4]," :", allUsersPaptis[i+4])

		fmt.Println()
		
	 }

}