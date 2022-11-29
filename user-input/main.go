package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Phone  int
	Status bool
	Age    int
}

func main() {
	// welcome := "Welcome to user imput"
	// fmt.Println(welcome)

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter something: ")

	// //comma ok syntax or error ok syntax
	// input, _ := reader.ReadString('\n')
	// fmt.Println("Thanks for entering : ", input)

	//conmveriosn
	// fmt.Println("Please enter a rating (1-5) : ")
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')
	// fmt.Println("Thanks for entering : ", input)
	// numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("Your rating is +1: ", numRating+1)
	// }

	//time
	// presentTime := time.Now()
	// fmt.Println(presentTime)
	// fmt.Println(presentTime.Format("01-02-2006"))
	// createDate := time.Date(2022, time.August, 23, 10, 10, 0, 0, time.UTC)
	// fmt.Println(createDate)

	//pointres
	// myNumber := 23
	// var ptr = &myNumber
	// fmt.Println(ptr, *ptr)
	// *ptr = *ptr + 2
	// fmt.Println(myNumber)

	//slices
	// var fruits = []string{}
	// fruits = append(fruits, "Mango", "Banana", "Peach", "Anything")
	// fmt.Printf("type of fruit list is : %T", fruits)
	// fruits = append(fruits[1:])
	// fmt.Println(fruits)
	// fruits = append(fruits[1:2])

	// highScores := make([]int, 4)
	// highScores[0] = 234
	// highScores[1] = 334
	// highScores[2] = 634
	// highScores[3] = 534
	// // highScores[4] = 534 //will not work

	// highScores = append(highScores, 345, 657, 754) //this will work. go assigns new places

	// fmt.Println(highScores)
	// sort.Ints(highScores)
	// fmt.Println(highScores)

	// //remove a value from a slice based on index
	// courses := []string{"javascript", "swift", "golang", "java", "rust"}
	// var index int16 = 1
	// courses = append(courses[:index], courses[index+1:]...)
	// fmt.Println(courses)

	// //maps
	// languages := make(map[string]string)
	// languages["js"] = "javascript"
	// languages["go"] = "golang"
	// languages["py"] = "python"

	// fmt.Println(languages)

	// // delete(languages, "py")
	// // fmt.Println(languages)

	// //loops
	// for key, value := range languages {
	// 	fmt.Println(key, "---", value)
	// }

	// //structs
	// newUser := User{"Rajat", "rajatky07@gmail.com", 8003650500, true, 27}
	// fmt.Println(newUser)

	// //can also use newUser.Name (same as objects in js)
	// fmt.Printf("%+v", newUser)

	// if num := 3; num < 10 {
	// 	fmt.Println("Number is less than 10")
	// } else {
	// 	fmt.Println("Number is greater than 10")
	// }

	//switch case
	// rand.Seed(time.Now().UnixNano())
	// diceNumber := rand.Intn(6) + 1
	// fmt.Println(diceNumber)
	// switch diceNumber {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// 	fallthrough
	// case 4:
	// 	fmt.Println("Four")
	// case 5:
	// 	fmt.Println("Five")
	// case 6:
	// 	fmt.Println("Six")
	// default:
	// 	fmt.Println("Nothing")
	// }

	//looing and break;
	// var days = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(d, days[d])
	// }
	// for i := range days {

	// for _, day := range days {
	// 	fmt.Println(day)
	// }
	// 	value := 1
	// 	for value < 10 {
	// 		if value == 5 {
	// 			value++
	// 			continue
	// 		}
	// 		if value == 3 {
	// 			goto test
	// 		}
	// 		fmt.Println(value)
	// 		value++
	// 	}

	// test:
	// 	fmt.Println("Test label")

	//functions
	// greeter()
	// result := add(3, 5)
	// fmt.Println(result)

	// pro, something := addUnlimited(4, 6, 7, 3, 6)
	// fmt.Println(pro, something)

	//METHODS
	newUser := User{"Rajat", "rajatky07@gmail.com", 8003650500, true, 27}
	fmt.Println(newUser)

	//can also use newUser.Name (same as objects in js)
	fmt.Printf("%+v \n", newUser)
	newUser.GetStatus()

}

func (u User) GetStatus() {
	fmt.Println("Is user active : ", u.Status)
}

// if want to change in actual struct then use pointers
func (u User) updateMail() {
	u.Email = "test@go.dev"
}

// func add(one int, two int) int {
// 	return one + two
// }
// func addUnlimited(val ...int) (int, string) {
// 	total := 0

// 	for _, value := range val {
// 		total += value
// 	}
// 	return total, "hello"
// }

// func greeter() {
// 	fmt.Println("Hello from external function.")
// }
