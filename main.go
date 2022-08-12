package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, World")
	fmt.Println(greeting("Pallat"))
	fmt.Println(greetingWithAge("Pallat", 30))
	fmt.Println(greetingWithAgeAndDrink("Pallat", 30, "Cola"))
	fmt.Println(isOdd(1))
	fmt.Println(sumOfFirst(3))
	fmt.Println(isPrime(1))
	fmt.Println(isPrime(2))
	fmt.Println(isPrime(3))
	fmt.Println(isPrime(4))
}


// greeting("Pallat") should return "Hello, Pallat"
func greeting(name string) string {
	return "Hello, " + name
}

// greetingWithAge("Pallat", 30) should return "Hello, Pallat. You are 30 years old."
func greetingWithAge(name string, age uint) string {
	return "Hello, " + name + ". You are " + strconv.FormatUint(uint64(age), 10) + " years old."
}

// greetingWithAgeAndDrink("Pallat", 30, "Cola") should return "Hello, Pallat. You are 30 years old and your favorite drink is Cola."
func greetingWithAgeAndDrink(name string, age uint, drink string) string {
	return "Hello, " + name + ". You are " + strconv.FormatUint(uint64(age), 10) + " years old and your favorite drink is " + drink + "."
}

func isOdd(n int) bool {
	if result := n % 2 == 1; result {
		return true
	}

	return false
}

// It should return sum of n, n-1, n-2, ..., 1
// sumOfFirst(3) should return 3+2+1=6
func sumOfFirst(n int) int {
	result := 0

	for i := 0; i < n; i++ {
		result += n - i
	}

	return result
}


// A prime number is a number greater than 1 with only two factors – themselves and 1
func isPrime(n int) bool {
	count := 0

	for i := 2; i <= n/2; i++ {
			if n % i == 0 {
					count++
					break
			}
	}

	return count == 0 && n != 1
}