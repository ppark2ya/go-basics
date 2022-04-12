package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age int
	favFood []string
}

func repeatMe(names ...string) {
	fmt.Println(names)
}

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

/**
* @see naked return
*/
func lenAndUpper(name string) (length int, uppercase string) {
	// 함수의 실행이 끝났을 때  추가적으로 실행 -> func의 return 이 끝난 후 실행
	defer fmt.Println("I'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func multiply(a int, b int) int {
	return a * b
}

func superAdd(numbers ...int) int {
	total := 0
	for index, number := range numbers {
		fmt.Println(index, number)
		total += number
	}

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }

	return total
}

func main() {
	// const name string = "jtpark"
	// 축약형은 func 에서만 사용 가능하고 변수에만 사용가능하다.
	// name := "jtpark"	// var name string = "jtpark"
	// name = "test"
	// fmt.Println(name)
	// fmt.Println(multiply(2, 2))

	// totalLength, upperName := lenAndUpper(name)
	// totalLength, _ := lenAndUpper(name)	-> return 데이터 중에서 사용하지 않는 값이 있을 경우 _(underscore)를 사용한다.
	// fmt.Println(totalLength, upperName)

	// repeatMe("Jtpark", "ghlee", "jslee")

	// total := superAdd(1, 2, 3, 4, 5, 6)
	// fmt.Println("superAdd return value: ", total)

	// fmt.Println(canIDrink(15))

	// pointer
	a := 2
	b := &a
	// 메모리 주소 접근: &
	// fmt.Println(&a, b, *b)

	// 메모리에 저장된 값 확인: *
	*b = 20
	// fmt.Println(a)

	// basic array
	// names := [5]string{"jtpark", "dka", "foo"}

	// slice
	names := []string{"jtpark", "dka", "foo"}
	names = append(names, "hello")
	// fmt.Println(names)

	// map[key]value
	nico := map[string]string{"name": "nico", "age": "12"}
	// fmt.Println(nico)

	for key, value := range nico {
		fmt.Println(key, value)
	}

	favFood := []string{"milk", "ramen"}
	// aPerson := person{"jtpark", 18, favFood}
	aPerson := person{name: "jtpark", age: 18, favFood: favFood}
	fmt.Println(aPerson)
}

func canIDrink(age int) bool {
	// variable expression
	if koreanAge := age + 2; koreanAge < 18 {
		return false
	}

	return true
}

func canIDrinkSwitch(age int) bool {
	switch {
	case age < 10:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}

	switch koreanAge := age + 2; koreanAge {
	case 10:
		return false
	case 18:
		return true
	}

	return false
}
