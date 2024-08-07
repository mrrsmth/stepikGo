package main

import (
	"fmt"
	"math"
	"strconv"
)

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

	// var inputNum int
	// var maxNumber int
	// var maxCount int

	// for fmt.Scan(&inputNum); inputNum != 0; fmt.Scan(&inputNum) {
	// 	if inputNum > maxNumber {
	// 		maxNumber = inputNum
	// 		maxCount = 1
	// 	} else if inputNum == maxNumber {
	// 		maxCount++
	// 	}
	// }

	// fmt.Println(maxCount)

	// Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.
	// Входные данные
	// Вводится 3 натуральных числа n, c, d, каждое из которых не превышает 10000.
	// Выходные данные
	// Вывести первое число от 1 до n включительно, кратное c, но НЕ кратное d. Если такого числа нет - выводить ничего не нужно.

	// var n, c, d int

	// fmt.Scan(&n)
	// fmt.Scan(&c)
	// fmt.Scan(&d)
	// scanThreeNum(n, c, d)

	// На вход подается число типа float64. Вам нужно вывести преобразованное число по правилу: число возводится в квадрат, затем обрезается дробная часть так что остается 4 знака после запятой. Но если число равно или меньше нуля - выводить:
	// "число R не подходит", где R - исходное число с 2 цифрами после запятой и с 2 по ширине. А если число больше чем 10 000 - выводить исходное число в экспоненциальном представлении (знак экспоненты в нижнем регистре).
	// Sample input:
	// -000012.2123
	// Sample output:
	// число -12.21 не подходит
	// Sample input:
	// 1000001
	// Sample output:
	// 1.000001e+06
	// Sample Input:
	// 12.12345678
	// Sample Output:
	// 146.9782

	var num float64
	fmt.Scan(&num)

	// Проверяем, меньше или равно ли число 0
	if num <= 0 {
		fmt.Printf("число %.2f не подходит\n", num)
		return
	}

	// Проверяем, больше ли число 10 000
	if num > 10000 {
		fmt.Printf("%.6e\n", num)
		return
	}

	// Возводим число в квадрат и обрезаем дробную часть до 4 знаков
	result := math.Pow(num, 2)
	resultStr := strconv.FormatFloat(result, 'f', 4, 64)
	fmt.Println(resultStr)

}

func scanThreeNum(n, c, d int) {
	for i := 1; i < n; i++ {
		if i%c == 0 && i%d != 0 {
			fmt.Println(i)
			break
		}
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
