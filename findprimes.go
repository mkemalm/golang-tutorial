/*
This program finds prime sums of even numbers.
*/
package main

import (
	"fmt"
	"os"
	"strconv"
)

func isItPrime(num int) bool {
	half := num / 2
	for i := 2; i <= half; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func primeArray(min int, max int) []int {
	var primes []int
	for i := min; i <= max; i++ {
		if isItPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes
}

func main() {
	var prime_part1 []int
	var prime_part2 []int
	var count int

	if len(os.Args) == 1 {
		println("Usage: findprimes NUMBER")
		os.Exit(1)
	}

	number_s := os.Args[1]
	number_i, error := strconv.Atoi(number_s)
	if error != nil {
		fmt.Println("Error conversion")
		os.Exit(1)
	}

	if number_i%2 != 0 {
		println("Number is not even")
		os.Exit(1)
	}

	if number_i <= 2 {
		println("Number should be greater then 2")
		os.Exit(1)
	}

	prime_half := number_i / 2
	prime_part1 = primeArray(2, prime_half)
	prime_part2 = primeArray(prime_half, number_i)
	println("\n" + number_s + " = \n")
	for i := range prime_part1 {
		for j := range prime_part2 {
			if (prime_part1[i] + prime_part2[j]) == number_i {
				println(strconv.Itoa(prime_part1[i]) + " + " + strconv.Itoa(prime_part2[j]))
				count++
			}
		}
	}

	println("\nTotal:" + strconv.Itoa(count))
}
