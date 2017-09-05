// Package fng generates random data given a file's bytes.
package fng

import (
	"errors"
	"hash/adler32"
	"math/rand"
)

// Default charset: (0-9 + a-z + A-Z).
var Charset = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n',
	'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
	'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

// Returns a random string by using adler32 checksum as seed for each character.
func GenerateString(data []byte, charset []byte, lenght int) (string, error) {
	var chars []byte
	sliced, err := sliceBytes(data, lenght)
	if err != nil {
		return "", err
	}

	for _, x := range sliced {
		chars = append(chars, randomChar(charset, int64(adler32.Checksum(x))))
	}

	return string(chars), nil
}

func sliceBytes(data []byte, times int) ([][]byte, error) {
	var start, end int
	fileSize := len(data)
	chunkSize := int(fileSize / times)

	// FIXME: improve error message
	if times > fileSize {
		times = fileSize
		return nil, errors.New("chunk size too high")
	}

	bytesMap := make([][]byte, times)

	bytesMap[0] = data[0:chunkSize]

	for i := 1; i < times; i++ {
		start = chunkSize * i
		end = chunkSize * (i + 1)
		bytesMap[i] = data[start:end]
	}

	return bytesMap, nil
}

func randomChar(charset []byte, seed int64) byte {
	rand.Seed(seed)

	return charset[rand.Intn(len(charset))]
}
