package prefix

import "io"

type Prefixer interface {
	// Returns field length encoded into []byte
	EncodeLength(maxLen, length int) ([]byte, error)

	// Returns the size of the field (number of characters, HEX-digits, bytes)
	DecodeLength(maxLen int, data []byte) (int, error)

	// Reads length from the Reader
	// TODO: return read bytes
	ReadLength(maxLen int, r io.Reader) (int, error)

	// Returns the number of bytes that takes to encode the length
	Length() int

	// Returns human readable information about length prefixer
	Inspect() string
}

type Prefixers struct {
	Fixed Prefixer
	L     Prefixer
	LL    Prefixer
	LLL   Prefixer
	LLLL  Prefixer
}

type PrefixerBuilder func(int) Prefixer
