package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "hey user howaya"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter rating")

	input, _ := reader.ReadString('\n')
	fmt.Println("tq", input)

}
