package main

import "fmt"

func main() {

	dimeter, circumfarence, area := Multifunction(10)

	fmt.Println(dimeter, circumfarence, area)

}

func Multifunction(r float32) (float32, float32, float32) {
	dimeter := 2 * r
	circumfarence := 2 * 3.14 * r
	area := 3.14 * r * r

	return dimeter, circumfarence, area

}
