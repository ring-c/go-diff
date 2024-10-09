package safetensors

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"

	"go-diff/pkg/tensor"
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

	var tensorStorages = tensor.Storages{}

	err = json.Unmarshal(header, &tensorStorages)
	if err != nil {
		err = fmt.Errorf("json.Unmarshal tensorStorages: %w", err)
		return
	}

	println()
	println()
	println()
	println()
	println()
	spew.Dump(tensorStorages.Data[0])
	spew.Dump(tensorStorages.Data[10])
	spew.Dump(tensorStorages.Data[20])

	return
}
