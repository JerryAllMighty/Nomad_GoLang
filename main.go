package main

import (
	"fmt"

	"github.com/learngo/accounts"
)

// func multiply(a int, b int) int {
// 	return a * b
// }

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

// var named string = "gg"

// func lenAndUpper(name string) (length int, upperCase string) {
// 	defer fmt.Println("I'm done")
// 	length = len(name)
// 	upperCase = strings.ToUpper(name)
// 	return
// }

// func repeatMe(words ...string) {
// 	fmt.Println(words)
// }

// func superAdd(numbers ...int) int {
// 	total := 0
// 	for i := 0; i < len(numbers); i++ {
// 		fmt.Println(numbers[i])
// 		total += numbers[i]
// 	}
// 	return total
// }

//	func canIDrink(age int) bool {
//		switch koreanAge := age + 2; koreanAge {
//		case 10:
//			return false
//		case 18:
//			return true
//		}
//		return false
//	}
//
//	func canIDrink(age int) bool {
//		switch koreanAge := age + 2; {
//		case koreanAge < 18:
//			return false
//		}
//		return true
//	}

// type person struct {
// 	name    string
// 	age     int
// 	favFood []string
// }

func main() {

	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account.Balance())
	// favFood := []string{"kimchi", "ramen"}
	// nico := person{name: "nico", age: 18, favFood: favFood}
	// fmt.Println(nico)
	// nico := map[string]string{"name": "kim", "age": "12"}

	// for _, value := range nico {
	// 	fmt.Println(value)
	// }

	// names := []string{"a", "b", "c"}
	// names[3] = "d"
	// names[4] = "e"

	// names = append(names, "f")

	// fmt.Println(names)

	// a := 2
	// b := a
	// c := &a
	// *c = 90
	// fmt.Println(a, b, *c)
	//fmt.Println(canIDrink(18))
	// result := superAdd(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// print(result)
	// repeatMe("java", "python", "javascript", "and then there's me")
	// a, b := lenAndUpper("abcd")
	// fmt.Println(a, b)
	// fmt.Println(multiply(2, 2))
	// something.SayHello()
	// something.sayBye()

	// const name string = "ymk"
	// name = "abcd"
	// fmt.Println(name)

	// var name string = "ymk"
	// name := "ymk"
	// name = "abcd"
	// fmt.Println(named)
	// totalLength, upperName := lenAndUpper("nico")
	// fmt.Println(totalLength, upperName)

}
