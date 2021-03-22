package gonbt

import (
	"encoding/binary"
	"io"
	"unsafe"
)

// Reader nbt
type Reader interface {
	String() (string, error)
	Byte() (byte, error)
	Short() (int16, error)
	Int() (int32, error)
	Long() (int64, error)
	Float() (float32, error)
	Double() (float64, error)
	Bytes() ([]byte, error)
	IntArray() ([]int32, error)
	LongArray() ([]int64, error)
}

type reader struct {
	flux io.Reader
}

// NewReader nbt
func NewReader(r io.Reader) Reader {
	return &reader{
		flux: r,
	}
}

// String reader with nbt format
func (r *reader) String() (string, error) {
	var size uint16
	var data []byte
	var err error

	bsize := make([]byte, unsafe.Sizeof(size))
	if _, err = r.flux.Read(bsize); err != nil {
		return "", err
	}
	size = binary.BigEndian.Uint16(bsize)
	if size == 0 {
		return "", nil
	}
	data = make([]byte, size)
	if _, err = r.flux.Read(data); err != nil {
		return "", err
	}
	return string(data), nil
}

// Byte reader with nbt format
func (r *reader) Byte() (byte, error) {
	var size uint8
	var data []byte
	var err error

	data = make([]byte, unsafe.Sizeof(size))
	if _, err = r.flux.Read(data); err != nil {
		return byte('0'), err
	}
	return data[0], nil
}

// Short reader with nbt format
func (r *reader) Short() (int16, error) {
	var size int16
	var err error

	b := make([]byte, unsafe.Sizeof(size))
	if _, err = r.flux.Read(b); err != nil {
		return int16(0), err
	}
	return int16(binary.BigEndian.Uint16(b)), nil
}

// Int reader with nbt format
func (r *reader) Int() (int32, error) {
	var size int32
	var err error

	b := make([]byte, unsafe.Sizeof(size))
	if _, err = r.flux.Read(b); err != nil {
		return int32(0), err
	}
	return int32(binary.BigEndian.Uint32(b)), nil
}

// Long reader with nbt format
func (r *reader) Long() (int64, error) {
	var size int64
	var err error

	b := make([]byte, unsafe.Sizeof(size))
	if _, err = r.flux.Read(b); err != nil {
		return int64(0), err
	}
	return int64(binary.BigEndian.Uint64(b)), nil
}

// Float reader with nbt format
func (r *reader) Float() (float32, error) {
	var size float32
	var err error

	b := make([]byte, unsafe.Sizeof(size))
	if _, err = r.flux.Read(b); err != nil {
		return float32(0.0), err
	}
	return float32(binary.BigEndian.Uint32(b)), nil
}

// Double reader with nbt format
func (r *reader) Double() (float64, error) {
	var size float64
	var err error

	b := make([]byte, unsafe.Sizeof(size))
	if _, err = r.flux.Read(b); err != nil {
		return float64(0), err
	}
	return float64(binary.BigEndian.Uint64(b)), nil
}

// Bytes reader with nbt format
func (r *reader) Bytes() ([]byte, error) {
	var size int32
	var data []byte
	var err error

	bsize := make([]byte, unsafe.Sizeof(size))
	if _, err = r.flux.Read(bsize); err != nil {
		return []byte{}, err
	}

	size = int32(binary.BigEndian.Uint32(bsize))
	if size == 0 {
		return []byte{}, nil
	}
	data = make([]byte, size)
	if _, err = r.flux.Read(data); err != nil {
		return []byte{}, err
	}
	return data, nil
}

// IntArray reader with nbt format
func (r *reader) IntArray() ([]int32, error) {
	var err error
	var ret []int32
	var nbr int32

	// get number element
	b := make([]byte, unsafe.Sizeof(nbr))
	if _, err = r.flux.Read(b); err != nil {
		return []int32{}, err
	}
	nbr = int32(binary.BigEndian.Uint32(b))

	for i := int32(0); i < nbr; i++ {
		var elem int32

		b := make([]byte, unsafe.Sizeof(elem))
		if _, err = r.flux.Read(b); err != nil {
			return []int32{}, err
		}
		elem = int32(binary.BigEndian.Uint32(b))
		ret = append(ret, elem)
	}
	return ret, nil
}

// LongArray reader with nbt format
func (r *reader) LongArray() ([]int64, error) {
	var err error
	var ret []int64
	var nbr int64

	// get number element
	b := make([]byte, unsafe.Sizeof(nbr))
	if _, err = r.flux.Read(b); err != nil {
		return []int64{}, err
	}
	nbr = int64(binary.BigEndian.Uint64(b))

	for i := int64(0); i < nbr; i++ {
		var elem int64

		b := make([]byte, unsafe.Sizeof(elem))
		if _, err = r.flux.Read(b); err != nil {
			return []int64{}, err
		}
		elem = int64(binary.BigEndian.Uint64(b))
		ret = append(ret, elem)
	}
	return ret, nil
}
