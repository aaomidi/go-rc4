package rc4

type Cipher struct {
	s    [256]byte
	i, j uint8
}

