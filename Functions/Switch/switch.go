package main

import "fmt"

// Variable expression
// func canIDrink(age int) bool {
// 	switch {
// 	case age < 18:
// 		return false
// 	case age >= 18:
// 		return true
// 	}
// 	return false
// }

func canIDrink(age int) bool {
	switch koreanAge := age + 2; {
	case koreanAge < 18:
		return false
	case koreanAge >= 18:
		return true
	}
	return false
}

func main() {
	fmt.Println(canIDrink(16))
}
