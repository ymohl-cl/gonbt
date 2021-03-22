package gonbt

import (
	"encoding/binary"
	"io"
	"math"
	"unsafe"
)

// Writer nbt
type Writer interface {
	String(string) error
	Byte(byte) error
	Short(int16) error
	Int(int32) error
	Long(int64) error
	Float(float32) error
	Double(float64) error
	Bytes([]byte) error
	IntArray([]int32) error
	LongArray([]int64) error
}

type writer struct {
	flux io.Writer
}

// NewWriter nbt
func NewWriter(driver io.Writer) Writer {
	return &writer{
		flux: driver,
	}
}

// String write with nbt format
func (w *writer) String(str string) error {
	var size uint16
	var err error

	bsize := make([]byte, unsafe.Sizeof(size))
	size = uint16(len(str))
	binary.BigEndian.PutUint16(bsize, size)
	if _, err = w.flux.Write(bsize); err != nil {
		return err
	}
	if _, err = w.flux.Write([]byte(str)); err != nil {
		return err
	}
	return nil
}

// Byte write with nbt format
func (w *writer) Byte(b byte) error {
	var err error

	if _, err = w.flux.Write([]byte{b}); err != nil {
		return err
	}
	return nil
}

// Short write with nbt format
func (w *writer) Short(v int16) error {
	var err error

	b := make([]byte, unsafe.Sizeof(v))
	binary.BigEndian.PutUint16(b, uint16(v))
	if _, err = w.flux.Write(b); err != nil {
		return err
	}
	return nil
}

// Int write with nbt format
func (w *writer) Int(v int32) error {
	var size int32
	var err error

	b := make([]byte, unsafe.Sizeof(size))
	binary.BigEndian.PutUint32(b, uint32(v))
	if _, err = w.flux.Write(b); err != nil {
		return err
	}
	return nil
}

// Long write with nbt format
func (w *writer) Long(v int64) error {
	var size int64
	var err error

	b := make([]byte, unsafe.Sizeof(size))
	binary.BigEndian.PutUint64(b, uint64(v))
	if _, err = w.flux.Write(b); err != nil {
		return err
	}
	return nil
}

// Float write with nbt format
func (w *writer) Float(v float32) error {
	var size float32
	var err error

	b := make([]byte, unsafe.Sizeof(size))
	binary.BigEndian.PutUint32(b, math.Float32bits(v))
	if _, err = w.flux.Write(b); err != nil {
		return err
	}
	return nil
}

// Double write with nbt format
func (w *writer) Double(v float64) error {
	var size float64
	var err error

	b := make([]byte, unsafe.Sizeof(size))
	binary.BigEndian.PutUint64(b, math.Float64bits(v))
	if _, err = w.flux.Write(b); err != nil {
		return err
	}
	return nil
}

// Bytes write with nbt format
func (w *writer) Bytes(v []byte) error {
	var size int32
	var err error

	bsize := make([]byte, unsafe.Sizeof(size))
	size = int32(len(v))
	binary.BigEndian.PutUint32(bsize, uint32(size))
	if _, err = w.flux.Write(bsize); err != nil {
		return err
	}
	if _, err = w.flux.Write(v); err != nil {
		return err
	}
	return nil
}

// IntArray write with nbt format
func (w *writer) IntArray(values []int32) error {
	var size int32
	var err error

	// get number element
	buf := make([]byte, unsafe.Sizeof(size))
	size = int32(len(values))
	binary.BigEndian.PutUint32(buf, uint32(size))
	if _, err = w.flux.Write(buf); err != nil {
		return err
	}
	for _, v := range values {
		buf = make([]byte, unsafe.Sizeof(v))
		binary.BigEndian.PutUint32(buf, uint32(v))
		if _, err = w.flux.Write(buf); err != nil {
			return err
		}
	}
	return nil
}

// LongArray write with nbt format
func (w *writer) LongArray(values []int64) error {
	var size int64
	var err error

	// get number element
	buf := make([]byte, unsafe.Sizeof(size))
	size = int64(len(values))
	binary.BigEndian.PutUint64(buf, uint64(size))
	if _, err = w.flux.Write(buf); err != nil {
		return err
	}
	for _, v := range values {
		buf = make([]byte, unsafe.Sizeof(v))
		binary.BigEndian.PutUint64(buf, uint64(v))
		if _, err = w.flux.Write(buf); err != nil {
			return err
		}
	}
	return nil
}
