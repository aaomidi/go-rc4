package rc4

import (
	"fmt"
)

func (c *Cipher) PrintTable() {
	fmt.Printf("Table: \n")
	for i := 0; i < 256; i++ {
		fmt.Printf("\t %#x", c.GetTable()[i])
		if (i+1)%16 == 0 {
			fmt.Println()
		}
	}
}

func (c *Cipher) GetTable() [256]byte {
	return c.s
}
func (c *Cipher) GetI() uint8 {
	return c.i
}

func (c *Cipher) GetJ() uint8 {
	return c.i
}

func (c *Cipher) PrintReport() {
	fmt.Printf("i: %d, j:%d\n", c.GetI(), c.GetJ())
	c.PrintTable()
}
