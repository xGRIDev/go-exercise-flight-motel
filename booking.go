package main

import "fmt"

/*type Booking struct {
	id        string
	hotelName string
	room      float64
	price     float64
}*/

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

func main() {
	//1
	book := [6]string{"ID", "HotelName", "Room", "Gender", "Price", "RateOfHotel"}
	fmt.Println(book)

	//2
	fmt.Println("All-Element : ", book[0:])
	HtlName := []string{book[1]}
	fmt.Println("First-Standalone Element : ", HtlName) //-StandAlone-First
	RoomGender := []string{book[2], book[3]}
	fmt.Println("Combined Element of 2 & 3 : ", RoomGender)

	//3
	htlToRoom := book[0:5]
	htlToRoom2 := book[:2]
	//	fmt.Println("The Element 1-2", htlToRoom)
	fmt.Println("Slice 1 (first & second element) :", htlToRoom)
	fmt.Println("Slice 1 (first & second element) :", htlToRoom2)

	//4
	fmt.Println(cap(book))
	RateHTLToRoom := book[1:5]
	fmt.Println("Resliced Is : ", RateHTLToRoom)

	//5
	mainGoals := []string{"Wrenwood-Motel", "xGRID"}
	fmt.Println("Initialization Goals : ", mainGoals)

	//6
	mainGoals[1] = "Go Checkin"
	mainGoals = append(mainGoals, "Contribute to our motel")
	fmt.Println("Update Direction : ", mainGoals)

}
