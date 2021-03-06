/*
Have the function FirstFactorial(num) take the num 
parameter being passed and return the factorial of it 
(e.g. if num = 4, return (4 * 3 * 2 * 1)). For the test cases, 
the range will be between 1 and 18 and the input will always 
be an integer. 
*/

package main

import (
	"fmt"
)

func FirstFactorial(num int) (result int) {
	//base case
	if num == 0{
		return 1
	} else { //recursive case
		return FirstFactorial(num - 1) * num
	}
	return result
}

func main() {
    fmt.Println(FirstFactorial(5)) //120
    fmt.Println(FirstFactorial(6)) //720
}