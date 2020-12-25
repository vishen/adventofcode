package main

import "fmt"

const secret = 20201227

func main() {
	assert(crack(7, 5764801) == 8, "incorrect loop size")
	assert(crack(7, 17807724) == 11, "incorrect loop size")

	assert(transform(5764801, 11) == 14897079, "incorrect result")
	assert(transform(17807724, 8) == 14897079, "incorrect result")

	// Part 1
	// 335121 -> door or card public key
	// 363891 -> door or card public key

	l1 := crack(7, 335121)
	key := transform(363891, l1)
	fmt.Println(key)
}

func transform(subject, loop int) int {
	val := 1
	for i := 0; i < loop; i++ {
		val *= subject
		if val > secret {
			val %= secret
		}
	}
	fmt.Println("transform", subject, loop, val)
	return val
}

func crack(subject, publicKey int) int {
	val := 1
	i := 1
	for {
		val *= subject
		if val > secret {
			val %= secret
		}
		if val == publicKey {
			break
		}
		i++
	}

	fmt.Println("crack", subject, publicKey, i, val)
	assert(publicKey == val, "public key not equal to val")
	return i
}

func assert(expr bool, msg string) {
	if !expr {
		panic(msg)
	}
}
