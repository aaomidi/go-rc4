package main

import (
	"fmt"
	x "github.com/aaomidi/rc4-go/rc4"
)

func main() {
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
