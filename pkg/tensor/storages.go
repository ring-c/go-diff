package tensor

import (
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

const SDMaxDims = 5

type Storages struct {
	Data []Storage
}

func (storages *Storages) UnmarshalJSON(headerData []byte) (err error) {

	// spew.Dump(string(data))

	var elements = make(map[string]headerElementData)

	err = json.Unmarshal(headerData, &elements)
	if err != nil {
		err = fmt.Errorf("json.Unmarshal elements: %w", err)
		return
	}

	storages.Data = make([]Storage, 0)

	for name, data := range elements {
		if name == "__metadata__" {
			continue
		}

		// if (is_unused_tensor(name)) {
		// 	continue;
		// }
		// println()
		// println()
		// spew.Dump(name)
		spew.Dump(data)

		var storage = Storage{
			Name: name,
		}

		storage.SetGGMLType(data.DType)

		if storage.DataType == GGML_TYPE_ERROR {
			err = fmt.Errorf("storage.DataType == GGML_TYPE_ERROR: %s", name)
			return
		}

		if len(data.Shape) > SDMaxDims {
			err = fmt.Errorf("invalid tensor (SDMaxDims): %s", name)
			return
		}

		var nDims = len(data.Shape)

		for i := range storage.NE {
			storage.NE[i] = 1
		}

		for i := 0; i < nDims; i++ {
			storage.NE[i] = data.Shape[i]
		}

		if nDims == 5 {
			if storage.NE[3] == 1 && storage.NE[4] == 1 {
				nDims = 4
			} else {
				err = fmt.Errorf("invalid tensor (nDims check): %s", name)
				return
			}
		}

		// ggml_n_dims returns 1 for scalars
		if nDims == 0 {
			nDims = 1
		}

		storage.NDims = nDims

		storages.Data = append(storages.Data, storage)
	}

	return
}

//goland:noinspection SpellCheckingInspection
type headerElementData struct {
	DType       string  `json:"dtype"`
	DataOffsets []int64 `json:"data_offsets"`
	Shape       []int64 `json:"shape"`
}
