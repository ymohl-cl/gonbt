package gonbt

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// yes i was testes my private functions :)
func TestWriter_String(t *testing.T) {
	// Write should not return an error and will be always nil.
	t.Run("should be ok with a full string", func(t *testing.T) {
		data := []byte{}
		r := bytes.NewBuffer(data)
		w := &writer{flux: r}

		expectedByte := append([]byte{0x00, 0x0d}, []byte("Hello World !")...)
		err := w.String("Hello World !")
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedByte, r.Bytes())
		}
	})
	t.Run("should be ok with an empty string", func(t *testing.T) {
		data := []byte{}
		r := bytes.NewBuffer(data)
		w := &writer{flux: r}

		expectedByte := []byte{0x00, 0x00}
		err := w.String("")
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedByte, r.Bytes())
		}
	})
}

func TestWriter_Byte(t *testing.T) {
	// Write should not return an error and will be always nil.
	t.Run("should be ok", func(t *testing.T) {
		data := []byte{}
		r := bytes.NewBuffer(data)
		w := &writer{flux: r}

		expectedByte := append([]byte{'A'})
		err := w.Byte(byte('A'))
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedByte, r.Bytes())
		}
	})
}

func TestWriter_Short(t *testing.T) {
	// Write should not return an error and will be always nil.
	t.Run("should be ok", func(t *testing.T) {
		data := []byte{}
		r := bytes.NewBuffer(data)
		w := &writer{flux: r}

		expectedByte := []byte{0x00, 0x2a}
		err := w.Short(int16(42))
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedByte, r.Bytes())
		}
	})
}

func TestWriter_Int(t *testing.T) {
	// Write should not return an error and will be always nil.
	t.Run("should be ok", func(t *testing.T) {
		data := []byte{}
		r := bytes.NewBuffer(data)
		w := &writer{flux: r}

		expectedByte := []byte{0x00, 0x00, 0x00, 0x2a}
		err := w.Int(int32(42))
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedByte, r.Bytes())
		}
	})
}

func TestWriter_Long(t *testing.T) {
	// Write should not return an error and will be always nil.
	t.Run("should be ok", func(t *testing.T) {
		data := []byte{}
		r := bytes.NewBuffer(data)
		w := &writer{flux: r}

		expectedByte := []byte{0x00, 0x00, 0x00, 0x00, 0x0a, 0x0a, 0x00, 0x0a}
		err := w.Long(int64(168427530))
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedByte, r.Bytes())
		}
	})
}

func TestWriter_Float(t *testing.T) {
	// Write should not return an error and will be always nil.
	t.Run("should be ok", func(t *testing.T) {
		data := []byte{}
		r := bytes.NewBuffer(data)
		w := &writer{flux: r}

		expectedByte := []byte{0x42, 0x28, 0xcc, 0xcd}
		err := w.Float(42.2)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedByte, r.Bytes())
		}
	})
}

func TestWriter_Double(t *testing.T) {
	// Write should not return an error and will be always nil.
	t.Run("should be ok", func(t *testing.T) {
		data := []byte{}
		r := bytes.NewBuffer(data)
		w := &writer{flux: r}

		expectedByte := []byte{0x40, 0x45, 0x19, 0x99, 0x99, 0x99, 0x99, 0x9a}
		err := w.Double(42.2)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedByte, r.Bytes())
		}
	})
}

func TestWriter_Bytes(t *testing.T) {
	// Write should not return an error and will be always nil.
	t.Run("should be ok with a full string", func(t *testing.T) {
		data := []byte{}
		r := bytes.NewBuffer(data)
		w := &writer{flux: r}

		expectedByte := append([]byte{0x00, 0x00, 0x00, 0x0d}, []byte("Hello World !")...)
		err := w.Bytes([]byte("Hello World !"))
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedByte, r.Bytes())
		}
	})
	t.Run("should be ok with an empty string", func(t *testing.T) {
		data := []byte{}
		r := bytes.NewBuffer(data)
		w := &writer{flux: r}

		expectedByte := []byte{0x00, 0x00, 0x00, 0x00}
		err := w.Bytes([]byte{})
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedByte, r.Bytes())
		}
	})
}

func TestWriter_IntArray(t *testing.T) {
	// Write should not return an error and will be always nil.
	t.Run("should be ok with an empty list", func(t *testing.T) {
		data := []byte{}
		r := bytes.NewBuffer(data)
		w := &writer{flux: r}

		expectedByte := []byte{0x00, 0x00, 0x00, 0x00}
		err := w.IntArray([]int32{})
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedByte, r.Bytes())
		}
	})
	t.Run("should be ok with 3 elements", func(t *testing.T) {
		data := []byte{}
		r := bytes.NewBuffer(data)
		w := &writer{flux: r}

		expectedByte := []byte{0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x2a}
		err := w.IntArray([]int32{11, 10, 42})
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedByte, r.Bytes())
		}
	})
}

func TestWriter_LongArray(t *testing.T) {
	// Write should not return an error and will be always nil.
	t.Run("should be ok with an empty list", func(t *testing.T) {
		data := []byte{}
		r := bytes.NewBuffer(data)
		w := &writer{flux: r}

		expectedByte := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
		err := w.LongArray([]int64{})
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedByte, r.Bytes())
		}
	})
	t.Run("should be ok with 3 elements", func(t *testing.T) {
		data := []byte{}
		r := bytes.NewBuffer(data)
		w := &writer{flux: r}

		expectedByte := []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03}
		expectedByte = append(expectedByte, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0b}...)
		expectedByte = append(expectedByte, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a}...)
		expectedByte = append(expectedByte, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2a}...)
		err := w.LongArray([]int64{11, 10, 42})
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedByte, r.Bytes())
		}
	})
}
