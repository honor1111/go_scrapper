package main

import "fmt"

func main() {

	// names := [5]string{"nico", "lynn", "dal"} // array
	//names := []string{"nico", "lynn", "dal"} // slice need append
	names := []string{"nico", "lynn", "dal"} // array
	names = append(names, "volala")          // append didn't revise input(names)

	fmt.Println(names)
}
