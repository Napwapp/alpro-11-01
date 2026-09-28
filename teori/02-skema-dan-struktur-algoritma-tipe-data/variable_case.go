package main 
import "fmt"

func main () {
	var name string

	var ( 
		fullName = "Muhamad Nawaf Abduh"
		firstName = "Muhamad"
	)
	
	name = "Muhamad Nawaf abduh" 
	var lastName = "S.kom"
	middleName := "Nawaf"

	fmt.Println("Nama :", name)
	fmt.Println("Nama tengah:", middleName)
	fmt.Println("Nama belakang:", lastName)
	fmt.Println(fullName)
	fmt.Println(firstName)
}