package tensor

type Storage struct {
	Name     string
	NE       [SDMaxDims]int64
	NDims    int
	DataType GGMLType
}

func (storage *Storage) SetGGMLType(dType string) {
	storage.DataType = GGML_TYPE_ERROR

	switch dType {
	case "F16":
		storage.DataType = GGML_TYPE_F16
	case "BF16":
		storage.DataType = GGML_TYPE_F32
	case "F32":
		storage.DataType = GGML_TYPE_F32
	case "F8_E4M3":
		storage.DataType = GGML_TYPE_F16
	}
}

type GGMLType int

// NOTE: always add types at the end of the enum to keep backward compatibility
// Go NOTE: moved GGML_TYPE_F32 from 0 to 4, so we can use 0 as error
const (
	GGML_TYPE_ERROR GGMLType = 0

	GGML_TYPE_F16  GGMLType = 1
	GGML_TYPE_Q4_0 GGMLType = 2
	GGML_TYPE_Q4_1 GGMLType = 3

	GGML_TYPE_F32 GGMLType = 4 // moved from 0

	// GGML_TYPE_Q4_2 = 4, support has been removed
	// GGML_TYPE_Q4_3 = 5, support has been removed

	GGML_TYPE_Q5_0     GGMLType = 6
	GGML_TYPE_Q5_1     GGMLType = 7
	GGML_TYPE_Q8_0     GGMLType = 8
	GGML_TYPE_Q8_1     GGMLType = 9
	GGML_TYPE_Q2_K     GGMLType = 10
	GGML_TYPE_Q3_K     GGMLType = 11
	GGML_TYPE_Q4_K     GGMLType = 12
	GGML_TYPE_Q5_K     GGMLType = 13
	GGML_TYPE_Q6_K     GGMLType = 14
	GGML_TYPE_Q8_K     GGMLType = 15
	GGML_TYPE_IQ2_XXS  GGMLType = 16
	GGML_TYPE_IQ2_XS   GGMLType = 17
	GGML_TYPE_IQ3_XXS  GGMLType = 18
	GGML_TYPE_IQ1_S    GGMLType = 19
	GGML_TYPE_IQ4_NL   GGMLType = 20
	GGML_TYPE_IQ3_S    GGMLType = 21
	GGML_TYPE_IQ2_S    GGMLType = 22
	GGML_TYPE_IQ4_XS   GGMLType = 23
	GGML_TYPE_I8       GGMLType = 24
	GGML_TYPE_I16      GGMLType = 25
	GGML_TYPE_I32      GGMLType = 26
	GGML_TYPE_I64      GGMLType = 27
	GGML_TYPE_F64      GGMLType = 28
	GGML_TYPE_IQ1_M    GGMLType = 29
	GGML_TYPE_BF16     GGMLType = 30
	GGML_TYPE_Q4_0_4_4 GGMLType = 31
	GGML_TYPE_Q4_0_4_8 GGMLType = 32
	GGML_TYPE_Q4_0_8_8 GGMLType = 33
)
