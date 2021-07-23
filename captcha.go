package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	var c1 int
	var c2 int
	var answer int
	c1 = rand.Intn(100)
	c2 = rand.Intn(80)
	fmt.Println("[ CAPTCHA CODE ]")
	fmt.Println(c2, " + ", c1)
	fmt.Scanln(&answer)

	if answer == c1 + c2 {
		fmt.Println("[ CORRECT CAPTCHA! ]")

	} else if answer != c1 + c2 {
		fmt.Println("[ INCORRECT CAPTCHA! ]")

		time.Sleep(2 * time.Second)

		main()
	}

}