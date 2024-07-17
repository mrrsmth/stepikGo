package main

import "fmt"

func main() {
	// var a int
	// fmt.Println("Input your number:")
	// fmt.Scan(&a)
	// b := a*2 + 100
	// fmt.Println(b)

	var c float32 = 9
	switch {
	case 1 <= c && c <= 9:
		fmt.Print("от 1 до 9 ")
		c--
		fallthrough
	case c == 10.2:
		fmt.Print("пройден ")
		fmt.Print(c)
	case 100 <= c && c <= 250:
		fmt.Print("от 100 до 250 ")
		fmt.Print(c)
	case 1000 <= c && c <= 6000:
		c += 999
		fmt.Print("от 1000 до 6000 ")
		fallthrough
	default:
		fmt.Print("default ")
	}

	var a int
	fmt.Scan(&a)
	switch {
	case a == 0:
		fmt.Print("Ноль")
	case a > 0:
		fmt.Print("Число положительное")
	case a < 0:
		fmt.Print("Число отрицательное")
	default:
		fmt.Print("default ")
	}
}
