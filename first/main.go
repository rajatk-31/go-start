package main

import "fmt"

// if varibale starts with capital letter then it is considered as public varibale and can be used by any file in folder
const Logintest = "hi"

func main() {
	var user string = "Rajat"
	fmt.Println(user)
	fmt.Printf("The type of user is : %T \n", user)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("The type of isLoggedIn is : %T \n", isLoggedIn)

	//implicit typoes
	var test = "ratfda"
	test2 := "agah"
	fmt.Println(test, test2, Logintest)
}
