// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var firstPrime int
// 	firstPrime = 2
// 	num := 10
// 	prime := [11]bool{}
// 	for i := 0; i < num+1; i++ {
// 		prime[i] = true
// 	}
// 	var det float64 = float64(firstPrime)
// 	var terminator float64 = float64(num)
// 	var index int
// 	max := math.Pow(det, 2)
// 	for max <= terminator {
// 		if prime[firstPrime] == true {
// 			for x := max; x < terminator+1; x += x {
// 				index = int(x)
// 				prime[index] = false
// 			}
// 		}
// 		firstPrime = firstPrime + 1
// 		var det float64 = float64(firstPrime)
// 		max = math.Pow(det, 2)

// 	}
// 	fmt.Printf("%d\n", firstPrime)
// 	fmt.Printf("%v", prime)

// }

package main

import (
	"fmt"
	"math"
)

func printPrimeNumbers(num1, num2 int) {
	if num2 < 1 {
		fmt.Println("Numbers must be greater than or equal to 1.")
		return
	}
	for num1 <= num2 {
		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(num1))); i++ {
			if num1%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", num1)
		}
		num1++
	}
	fmt.Println()
}

func main() {
	//var i int
	//_, err := fmt.Scanf("%d", &i)
	var i int
	fmt.Scan(&i)
	fmt.Println("read number", i, "from stdin")
	num1 := 0
	num2 := i

	printPrimeNumbers(num1, num2)
}
