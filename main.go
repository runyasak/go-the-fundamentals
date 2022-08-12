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