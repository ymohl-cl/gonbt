package gonbt

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader_String(t *testing.T) {
	t.Run("read return an error because reader is empty", func(t *testing.T) {
		data := []byte{}
		r := &reader{flux: bytes.NewReader(data)}

		expectedError := "EOF"
		str, err := r.String()
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedError, err.Error())
			assert.Empty(t, str)
		}
	})
	t.Run("read return an error because reader is empty after the bsize loaded", func(t *testing.T) {
		data := []byte{0x00, 0x0a}
		r := &reader{flux: bytes.NewReader(data)}

		expectedError := "EOF"
		str, err := r.String()
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedError, err.Error())
			assert.Empty(t, str)
		}
	})
	t.Run("Should be ok with an empty string", func(t *testing.T) {
		data := []byte{0x00, 0x00}
		r := &reader{flux: bytes.NewReader(data)}

		str, err := r.String()
		if assert.NoError(t, err) {
			assert.Empty(t, str)
		}
	})
	t.Run("Should be", func(t *testing.T) {
		data := append([]byte{0x00, 0x0a}, []byte("Hello YOU!")...)
		r := &reader{flux: bytes.NewReader(data)}

		expectedSTR := "Hello YOU!"
		str, err := r.String()
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedSTR, str)
		}
	})
}

func TestReader_Byte(t *testing.T) {
	t.Run("should return an error because the data is empty", func(t *testing.T) {
		data := []byte{}
		r := &reader{flux: bytes.NewReader(data)}

		expectedError := "EOF"
		expectedByte := byte('0')
		ret, err := r.Byte()
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedError, err.Error())
			assert.EqualValues(t, expectedByte, ret)
		}
	})
	t.Run("should be ok", func(t *testing.T) {
		data := []byte{0x0a}
		r := &reader{flux: bytes.NewReader(data)}

		expectedByte := byte(10)
		ret, err := r.Byte()
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedByte, ret)
		}
	})
}

func TestReader_Short(t *testing.T) {
	t.Run("should return an error because the data is empty", func(t *testing.T) {
		data := []byte{}
		r := &reader{flux: bytes.NewReader(data)}

		expectedError := "EOF"
		expectedRet := int16(0)
		ret, err := r.Short()
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedError, err.Error())
			assert.EqualValues(t, expectedRet, ret)
		}
	})
	t.Run("should be ok", func(t *testing.T) {
		data := []byte{0x00, 0x0a}
		r := &reader{flux: bytes.NewReader(data)}

		expectedRet := int16(10)
		ret, err := r.Short()
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRet, ret)
		}
	})
}

func TestReader_Int(t *testing.T) {
	t.Run("should return an error because the data is empty", func(t *testing.T) {
		data := []byte{}
		r := &reader{flux: bytes.NewReader(data)}

		expectedError := "EOF"
		expectedRet := int32(0)
		ret, err := r.Int()
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedError, err.Error())
			assert.EqualValues(t, expectedRet, ret)
		}
	})
	t.Run("should be ok", func(t *testing.T) {
		data := []byte{0x0a, 0x0a, 0x00, 0x0a}
		r := &reader{flux: bytes.NewReader(data)}

		expectedRet := int32(168427530)
		ret, err := r.Int()
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRet, ret)
		}
	})
}

func TestReader_Long(t *testing.T) {
	t.Run("should return an error because the data is empty", func(t *testing.T) {
		data := []byte{}
		r := &reader{flux: bytes.NewReader(data)}

		expectedError := "EOF"
		expectedRet := int64(0)
		ret, err := r.Long()
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedError, err.Error())
			assert.EqualValues(t, expectedRet, ret)
		}
	})
	t.Run("should be ok", func(t *testing.T) {
		data := []byte{0x00, 0x00, 0x00, 0x00, 0x0a, 0x0a, 0x00, 0x0a}
		r := &reader{flux: bytes.NewReader(data)}

		expectedRet := int64(168427530)
		ret, err := r.Long()
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRet, ret)
		}
	})
}

func TestReader_Float(t *testing.T) {
	t.Run("should return an error because the data is empty", func(t *testing.T) {
		data := []byte{}
		r := &reader{flux: bytes.NewReader(data)}

		expectedError := "EOF"
		expectedRet := float32(0)
		ret, err := r.Float()
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedError, err.Error())
			assert.EqualValues(t, expectedRet, ret)
		}
	})
	t.Run("should be ok", func(t *testing.T) {
		data := []byte{0x0a, 0x0a, 0x00, 0x0a}
		r := &reader{flux: bytes.NewReader(data)}

		expectedRet := float32(168427530)
		ret, err := r.Float()
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRet, ret)
		}
	})
}

func TestReader_Double(t *testing.T) {
	t.Run("should return an error because the data is empty", func(t *testing.T) {
		data := []byte{}
		r := &reader{flux: bytes.NewReader(data)}

		expectedError := "EOF"
		expectedRet := float64(0)
		ret, err := r.Double()
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedError, err.Error())
			assert.EqualValues(t, expectedRet, ret)
		}
	})
	t.Run("should be ok", func(t *testing.T) {
		data := []byte{0x00, 0x00, 0x00, 0x00, 0x0a, 0x0a, 0x00, 0x0a}
		r := &reader{flux: bytes.NewReader(data)}

		expectedRet := float64(168427530)
		ret, err := r.Double()
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRet, ret)
		}
	})
}

func TestReader_Bytes(t *testing.T) {
	t.Run("read return an error because reader is empty", func(t *testing.T) {
		data := []byte{}
		r := &reader{flux: bytes.NewReader(data)}

		expectedError := "EOF"
		ret, err := r.Bytes()
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedError, err.Error())
			assert.Empty(t, ret)
		}
	})
	t.Run("read return an error because reader is empty after the bsize loaded", func(t *testing.T) {
		data := []byte{0x00, 0x00, 0x00, 0x0a}
		r := &reader{flux: bytes.NewReader(data)}

		expectedError := "EOF"
		ret, err := r.Bytes()
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedError, err.Error())
			assert.Empty(t, ret)
		}
	})
	t.Run("Should be ok with an empty string", func(t *testing.T) {
		data := []byte{0x00, 0x00, 0x00, 0x00}
		r := &reader{flux: bytes.NewReader(data)}

		ret, err := r.Bytes()
		if assert.NoError(t, err) {
			assert.Empty(t, ret)
		}
	})
	t.Run("Should be", func(t *testing.T) {
		data := append([]byte{0x00, 0x00, 0x00, 0x0a}, []byte("Hello YOU!")...)
		r := &reader{flux: bytes.NewReader(data)}

		expectedRet := []byte("Hello YOU!")
		ret, err := r.Bytes()
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRet, ret)
		}
	})
}

func TestReader_IntArray(t *testing.T) {
	t.Run("should return an error because flux is empty", func(t *testing.T) {
		data := []byte{}
		r := &reader{flux: bytes.NewReader(data)}

		expectedError := "EOF"
		ret, err := r.IntArray()
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedError, err.Error())
			assert.Empty(t, ret)
		}
	})
	t.Run("should return an error because the flux is corrompted", func(t *testing.T) {
		data := []byte{0x00, 0x00, 0x00, 0x0a}
		r := &reader{flux: bytes.NewReader(data)}

		expectedError := "EOF"
		ret, err := r.IntArray()
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedError, err.Error())
			assert.Empty(t, ret)
		}
	})
	t.Run("should be ok with an empty list", func(t *testing.T) {
		data := []byte{0x00, 0x00, 0x00, 0x00}
		r := &reader{flux: bytes.NewReader(data)}

		ret, err := r.IntArray()
		if assert.NoError(t, err) {
			assert.Empty(t, ret)
		}
	})
	t.Run("should be", func(t *testing.T) {
		data := []byte{0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x2a}
		r := &reader{flux: bytes.NewReader(data)}

		expectedRet := []int32{11, 10, 42}
		ret, err := r.IntArray()
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRet, ret)
		}
	})
}

func TestReader_LongArray(t *testing.T) {
	t.Run("should return an error because flux is empty", func(t *testing.T) {
		data := []byte{}
		r := &reader{flux: bytes.NewReader(data)}

		expectedError := "EOF"
		ret, err := r.LongArray()
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedError, err.Error())
			assert.Empty(t, ret)
		}
	})
	t.Run("should return an error because the flux is corrompted", func(t *testing.T) {
		data := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a}
		r := &reader{flux: bytes.NewReader(data)}

		expectedError := "EOF"
		ret, err := r.LongArray()
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedError, err.Error())
			assert.Empty(t, ret)
		}
	})
	t.Run("should be ok with an empty list", func(t *testing.T) {
		data := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		r := &reader{flux: bytes.NewReader(data)}

		ret, err := r.LongArray()
		if assert.NoError(t, err) {
			assert.Empty(t, ret)
		}
	})
	t.Run("should be", func(t *testing.T) {
		data := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03}
		data = append(data, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0b}...)
		data = append(data, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a}...)
		data = append(data, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2a}...)
		r := &reader{flux: bytes.NewReader(data)}

		expectedRet := []int64{11, 10, 42}
		ret, err := r.LongArray()
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedRet, ret)
		}
	})
}
