/*
public and private keys are generated using RSA algorithm
*/

package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

var public_key int
var private_key int
var n int
var primes []int = []int{
	101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151,
	157, 163, 167, 173, 179, 181, 191, 193, 197, 199, 211,
	223, 227, 229, 233, 239, 241, 251, 257, 263, 269, 271,
	277, 281, 283, 293, 307, 311, 313, 317, 331, 337, 347,
	349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409,
	419, 421, 431, 433, 439, 443, 449, 457, 461, 463, 467,
	479, 487, 491, 499, 503, 509, 521, 523, 541, 547, 557,
	563, 569, 571, 577, 587, 593, 599, 601, 607, 613, 617,
	619, 631, 641, 643, 647, 653, 659, 661, 673, 677, 683,
	691, 701, 709, 719, 727, 733, 739, 743, 751, 757, 761,
	769, 773, 787, 797, 809, 811, 821, 823, 827, 829, 839,
	853, 857, 859, 863, 877, 881, 883, 887, 907, 911, 919,
	929, 937, 941, 947, 953, 967, 971, 977, 983, 991, 997,
}

func getPrime() int {
	idx := rand.Intn(143)
	if primes[idx] == -1 {
		return getPrime()
	}
	return primes[idx]
}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}

func init() {
	p := getPrime()
	q := getPrime()
	n = p * q
	fi := (p - 1) * (q - 1)
	e := 2
	for {
		if gcd(e, fi) == 1 {
			break
		}
		e++
	}
	public_key = e

	d := 2
	for {
		if (d*e)%fi == 1 {
			break
		}
		d++
	}
	private_key = d
}

func main() {
	fmt.Println("Public Key  : ", public_key)
	fmt.Println("Private Key : ", private_key)
	fmt.Println("Hex Address : ", SHA256(strconv.Itoa(public_key)))
}
