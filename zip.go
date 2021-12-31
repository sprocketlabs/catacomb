package catacomb

import (
	"bytes"
	"compress/gzip"
	"log"
)

func compress(data []byte) []byte {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(data); err != nil {
		log.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		log.Fatal(err)
	}
	return b.Bytes()
}

func decompress(compressedData []byte) []byte {
	var b bytes.Buffer
	gz, err := gzip.NewReader(&b)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := gz.Read(compressedData); err != nil {
		log.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		log.Fatal(err)
	}
	return b.Bytes()
}
