package rc4

func NewCipher(key []byte) *Cipher {
	var c Cipher

	N := len(key)
	k := make([]byte, 256)

	for i := 0; i < 256; i++ {
		c.s[i] = byte(i)
		k[i] = byte(key[i%N])
	}

	var j uint8 = 0

	for i := 0; i < 256; i++ {
		j = j + c.s[i] + k[i]

		// Swap
		c.s[i], c.s[j] = c.s[j], c.s[i]
	}
	return &c
}

func (c *Cipher) StepForward() byte {
	i, j := c.i, c.j

	i = i + 1
	j = j + c.s[i]

	// Swap
	c.s[i], c.s[j] = c.s[j], c.s[i]

	t := c.s[i] + c.s[j]

	// Set it back
	c.i, c.j = i, j

	return c.s[t]
}

func (c *Cipher) StepRounds(n int) byte {
	var b byte

	for i := 0; i <= n; i++ {
		b = c.StepForward()
	}

	return b
}

func (c *Cipher) Apply(dst, src []byte) {
	for k, v := range src {
		dst[k] = v ^ c.StepForward()
	}
}
