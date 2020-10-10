package main

import "fmt"

func main()  {
	// n = number of elements
	// v = value of element to store in array
	// t = total

	// s = array(slice) to store n values
	var n, v, t int
	var s []int
	
	fmt.Scan(&n)

	// bucle to get input, store it in slice, add to total
	for i := 0; i < n; i++ {
		fmt.Scan(&v)
		s = append(s, v)
		t += v
	}
	
	// output is total
	fmt.Println(t)
}