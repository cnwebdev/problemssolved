/*************************************************************************************************
 * If we list all the nautual number below 10 that are multiples of 3 and 5, we get 3, 5, 6, and 9.
 * the sume of these multiples is 23
 */

package main

import "fmt"

func main() {
	multiples3n5 := []int{}

	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			multiples3n5 = append(multiples3n5, i)
		}
	}

	fmt.Println("Multiples of 3 an 5 in 1 to 1000 are:")
	sum := 0
	for i, n := range multiples3n5 {
		fmt.Printf("#%-5d %d\n", i, n)
		sum += n
	}
	fmt.Printf("The sume of multiple of 3 and 5 in 1 to 1000 is %d\n", sum)
}
