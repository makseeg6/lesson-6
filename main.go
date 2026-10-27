package main

import "fmt"

func main() {
	// #1
	fmt.Println(5 == 5) // true

	fmt.Println(10 != 3) // true

	fmt.Println(7 > 12) // false

	fmt.Println(15 < 20) // true

	fmt.Println(8 >= 8) // true

	fmt.Println(6 <= 4) // false

	fmt.Println((10 > 5) && (3 < 1)) // false

	fmt.Println((10 > 5) || (3 < 1)) // true

	fmt.Println(!(5 == 5)) // false

	fmt.Println(!(7 < 3)) // true

	fmt.Println(true && false) // false

	fmt.Println(false || false) // false

	fmt.Println(true || false) // true

	fmt.Println((4+6 == 10) && (9 > 2)) // true

	fmt.Println((12/3 == 4) || (8 < 5)) // true

	// #2
	age := 23

	hasTicket := true

	canEnter := age >= 18 && hasTicket
	fmt.Println(canEnter)

	// #3
	isLoggedin := true
	isAdmin := false

	hasAccess := isLoggedin && (isAdmin || !isAdmin)
	fmt.Println(hasAccess)
}
