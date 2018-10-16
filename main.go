package main

import (
	"fmt"
	x "github.com/aaomidi/go-rc4/rc4"
)

func main() {
	test1()
	test2()
}

func test1() {
	c1 := x.NewCipher([]byte{0x1A, 0x2B, 0x3C, 0x4D, 0x5E, 0x6F, 0x77})
	c2 := x.NewCipher([]byte{0x1A, 0x2B, 0x3C, 0x4D, 0x5E, 0x6F, 0x77})

	msg := []byte("Hey this is Amir")

	// Get rid of the problematic 256 bytes
	c1.StepRounds(256)
	c2.StepRounds(256)

	encrypted := make([]byte, len(msg))
	decrypted := make([]byte, len(msg))

	c1.Apply(encrypted, msg)
	c2.Apply(decrypted, encrypted)

	fmt.Println(string(decrypted))
}

func test2() {
	c := x.NewCipher([]byte{0x1A, 0x2B, 0x3C, 0x4D, 0x5E, 0x6F, 0x77})

	fmt.Println("Round 0:")
	c.PrintReport()

	fmt.Println("Round 100:")
	c.StepRounds(100)
	c.PrintReport()

	fmt.Println("Round 1000:")
	c.StepRounds(900)
	c.PrintReport()

}
