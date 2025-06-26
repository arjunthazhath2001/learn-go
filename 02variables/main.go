package main

import "fmt"

const PubLoginToken = "nhkjnkjn"

func main() {
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("its type is %T \n", isLoggedIn)

	var val int
	fmt.Println(val)
	fmt.Printf("its type is %T \n", val)

	var floval float32 = 256.3233323333333
	fmt.Println(floval)
	fmt.Printf("its type is %T \n", floval)

	var name = "arjun"
	fmt.Println(name)

	numusers := 3000.00
	fmt.Println(numusers)
	fmt.Printf("its type is %T \n", numusers)

	fmt.Println(PubLoginToken)
	fmt.Printf("its type is %T \n", PubLoginToken)

}
