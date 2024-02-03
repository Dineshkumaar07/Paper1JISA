package algorithms

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func gcd(a int, b int) int {
	if a < b {
		return gcd(b, a)
	} else if a%b == 0 {
		return b
	} else {
		return gcd(b, a%b)
	}
}

func genKey(q int) int {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	min := 100
	max := q
	key := rng.Intn(max-min+1) + min
	for gcd(q, key) != 1 {
		key = rng.Intn(max-min+1) + min
	}
	return key
}

func power(a int, b int, c int) int {
	x := 1
	y := a

	for b > 0 {
		if b%2 != 0 {
			x = (x * y) % c
		}
		y = (y * y) % c
		b = int(b / 2)
	}
	return x % c

}

func encrypt(msg string, q int, h int, g int) ([]int, int) {
	var enMsg strings.Builder
	enMsg.WriteString("")
	modified := []int{}

	k := genKey(q)
	s := power(h, k, q)
	p := power(g, k, q)

	for _, char := range msg {
		enMsg.WriteRune(char)
	}

	for _, char := range enMsg.String() {
		newChar := rune(s * int(char))
		modified = append(modified, int(newChar))
	}
	return modified, p

}

func decrypt(enMsg []int, p int, key int, q int) []rune {
	deMsg := []rune{}
	h := power(p, key, q)
	for _, char := range enMsg {
		newChar := rune(int(char) / h)
		deMsg = append(deMsg, newChar)
	}
	return deMsg

}

func Elgamal() {
	msg := "Dinesh M S"
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	minq := 100
	maxq := 100000

	q := rng.Intn(maxq-minq+1) + minq
	ming := 2
	maxg := q
	g := rng.Intn(maxg-ming+1) + ming

	key := genKey(q)
	h := power(g, key, q)

	enMsg, p := encrypt(msg, q, h, g)
	fmt.Println("Message: ", msg)
	fmt.Println("Encrypted Text : ", enMsg)
	deMsg := decrypt(enMsg, p, key, q)
	fmt.Println("Decrypted: ", string(deMsg))

}
