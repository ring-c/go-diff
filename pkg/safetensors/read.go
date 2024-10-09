package safetensors

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

// Read - Reading safetensors file
// See: https://huggingface.co/docs/safetensors/index
func Read(filename string) (err error) {
	f, err := os.Open(filename)
	if err != nil {
		err = fmt.Errorf("os.Open: %w", err)
		return
	}
	defer func() {
		_ = f.Close()
	}()

	var headerSizeData = make([]byte, 8)
	_, err = f.Read(headerSizeData)
	if err != nil {
		err = fmt.Errorf("f.Read headerSizeData: %w", err)
		return
	}

	headerSize := binary.LittleEndian.Uint64(headerSizeData)

	spew.Dump(headerSize)

	header := make([]byte, headerSize)
	_, err = f.Read(header)
	if err != nil {
		err = fmt.Errorf("f.Read header: %w", err)
		return
	}

	spew.Dump(string(header))

	return
}
