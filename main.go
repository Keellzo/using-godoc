// main is a simple package with only one function, the function prints the dog years of one human year
package main

import "fmt"

func main() {
	fmt.Println(Years(5))
}

// Years takes dogYears as a parameter and then calculates it too how many years the dog is according to humans
func Years(dogYears int) int {
	return dogYears * 7
}
