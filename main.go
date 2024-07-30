package main

import "fmt"

func main() {
	// var a int
	// fmt.Println("Input your number:")
	// fmt.Scan(&a)
	// b := a*2 + 100
	// fmt.Println(b)

	// var c float32 = 9
	// switch {
	// case 1 <= c && c <= 9:
	// 	fmt.Print("от 1 до 9 ")
	// 	c--
	// 	fallthrough
	// case c == 10.2:
	// 	fmt.Print("пройден ")
	// 	fmt.Print(c)
	// case 100 <= c && c <= 250:
	// 	fmt.Print("от 100 до 250 ")
	// 	fmt.Print(c)
	// case 1000 <= c && c <= 6000:
	// 	c += 999
	// 	fmt.Print("от 1000 до 6000 ")
	// 	fallthrough
	// default:
	// 	fmt.Print("default ")
	// }

	// var a int
	// fmt.Scan(&a)
	// switch {
	// case a == 0:
	// 	fmt.Print("Ноль")
	// case a > 0:
	// 	fmt.Print("Число положительное")
	// case a < 0:
	// 	fmt.Print("Число отрицательное")
	// default:
	// 	fmt.Print("default ")
	// }

	// var a int
	// fmt.Scan(&a)

	// for i := 1; i < 11; i++ {
	// 	a = i * i
	// 	fmt.Println(a)
	// }

	// Требуется написать программу, при выполнении которой с клавиатуры считываются два натуральных числа A и B (каждое не более 100, A < B). Вывести сумму всех чисел от A до B  включительно.
	// 	Sample Input:
	// 1 5
	// Sample Output:
	// 15

	// var a, b int

	// fmt.Scan(&a)
	// fmt.Scan(&b)

	// fmt.Println(sumNum(a, b))

	// Напишите программу, которая в последовательности чисел находит сумму двузначных чисел, кратных 8.
	// Программа в первой строке получает на вход число n - количество чисел в последовательности, во второй строке -- n чисел, входящих в данную последовательность.

	// var n int
	// fmt.Scan(&n)
	// sum := 0

	// for i := 0; i < n; i++ {
	// 	var dirtyNum int
	// 	fmt.Scan(&dirtyNum)

	// 	if dirtyNum >= 10 && dirtyNum <= 99 && dirtyNum%8 == 0 {
	// 		sum += dirtyNum
	// 	}
	// }
	// fmt.Println(sum)

	// var num int
	// fmt.Scan(&num)
	// fmt.Println(forSum(num))

	// 	Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.
	// Формат входных данных:
	// Вводится непустая последовательность натуральных чисел, оканчивающаяся числом 0 (само число 0 в последовательность не входит, а служит как признак ее окончания).
	// Формат выходных данных:
	// Выведите ответ на задачу.

	// var inputNum int
	// slice := []int{}
	// fmt.Scan(&inputNum)
	// for inputNum != 0 {
	// 	slice = append(slice, inputNum)
	// }
	// fmt.Println(slice)

	var inputNum, lastNumber int
	// count := 0
	fmt.Scan(&inputNum)
	for fmt.Scan(&inputNum); inputNum != 0; fmt.Scan(&inputNum) {
		lastNumber = inputNum
		fmt.Println(lastNumber)
	}
}

func forSum(num int) int {
	sum := 0

	for i := 0; i < num; i++ {
		var dirtyNum int
		fmt.Scan(&dirtyNum)

		if dirtyNum >= 10 && dirtyNum <= 99 && dirtyNum%8 == 0 {
			sum += dirtyNum
		}
	}
	return sum
}

func sumNum(a, b int) int {
	sum := 0

	for i := a; i <= b; i++ {
		sum += i
	}
	return sum
}
