package misc

import (
	"io"
	"os"
)

func Rand(sz int) []byte {
	reader, err := os.Open("/dev/urandom")
	buff := make([]byte, sz)

	if err != nil {
		panic(err)
	}


	io.ReadFull(reader, buff)
	return buff
}
