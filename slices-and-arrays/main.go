package main

import "fmt"


func main() {
	// array ( 크기 고정 )
	foods := [3]string{"potato", "pizza", "pasta"}

	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}

	for _, dish := range foods {
		fmt.Println(dish)
	}

	// slice (크기 제한 없음)
	sliceFoods := []string{"potato", "pizza", "pasta"}
	fmt.Printf("%v\n", sliceFoods)
	// 원본을 변경하지 않고 새로운 item을 추가한 slice를 return
	sliceFoods = append(sliceFoods, "tomato")
	fmt.Printf("%v\n", sliceFoods)
}
