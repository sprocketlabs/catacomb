package catacomb

import (
	"compress/gzip"
	"fmt"
	"os"
)

func compress(data []byte) {
	original := "bird and frog"

	// Open a file for writing.
	f, _ := os.Create("C:\\programs\\file.gz")

	// Create gzip writer.
	w := gzip.NewWriter(f)

	// Write bytes in compressed form to the file.
	w.Write([]byte(original))

	// Close the file.
	w.Close()

	fmt.Println("DONE")
}

func decompress(compressedData []byte) {
	// Open the gzip file.
	f, _ := os.Open("C:\\programs\\file.gz")

	// Create new reader to decompress gzip.
	reader, _ := gzip.NewReader(f)

	// Empty byte slice.
	result := make([]byte, 100)

	// Read in data.
	count, _ := reader.Read(result)

	// Print our decompressed data.
	fmt.Println(count)
	fmt.Println(string(result))

}
