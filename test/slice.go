package test;

import (
	"fmt"
)

func testArrayInt() (slice interface{}) {
	slice = make([]int, 100)
	fmt.Println(slice)
	fmt.Println(&slice)
	// slice = append(slice.<,10)
	return slice
}


// func main() {
// 	s1 := make([]int, 10, 10) 
// 	s2 := s1
	

// 	s3 := append(s1, 1000)
// 	s1[0] = 10;
// 	// scores[7] = 9033
	
// 	fmt.Println(s1) 
// 	fmt.Println(s2)
// 	fmt.Println(s3)
// 	// fmt.Print(len(scores	)) 
// 	// fmt.Println(scores)
// }