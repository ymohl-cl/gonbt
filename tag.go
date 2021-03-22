package gonbt

import (
	"errors"
)

// TagType values
const (
	TagEnd byte = iota
	TagByte
	TagShort
	TagInt
	TagLong
	TagFloat
	TagDouble
	TagByteArray
	TagString
	TagList
	TagCompound
	TagIntArray
	TagLongArray
)

// Tag interface to provide a nbt reader / writer
type Tag interface {
	Read(reader Reader) error
	Write(writer Writer, printInfo bool) error
}

// NewTag instance
func NewTag(tagT byte, name string) (Tag, error) {
	switch tagT {
	case TagEnd:
		return nil, errors.New(errorEnd)
	case TagByte:
		return &ByteT{name: name}, nil
	case TagShort:
		return &ShortT{name: name}, nil
	case TagInt:
		return &IntT{name: name}, nil
	case TagLong:
		return &LongT{name: name}, nil
	case TagFloat:
		return &FloatT{name: name}, nil
	case TagDouble:
		return &DoubleT{name: name}, nil
	case TagByteArray:
		return &ByteArrayT{name: name}, nil
	case TagString:
		return &StringT{name: name}, nil
	case TagList:
		return &ListT{name: name}, nil
	case TagCompound:
		return &CompoundT{name: name}, nil
	case TagIntArray:
		return &IntArrayT{name: name}, nil
	case TagLongArray:
		return &LongArrayT{name: name}, nil
	default:
		return nil, errors.New(errorTag)
	}
}

// TagType return the tag type from the Tag parameter
func TagType(tag Tag) (byte, error) {
	switch tag.(type) {
	case *ByteT:
		return TagByte, nil
	case *ShortT:
		return TagShort, nil
	case *IntT:
		return TagInt, nil
	case *LongT:
		return TagLong, nil
	case *FloatT:
		return TagFloat, nil
	case *DoubleT:
		return TagDouble, nil
	case *ByteArrayT:
		return TagByteArray, nil
	case *StringT:
		return TagString, nil
	case *ListT:
		return TagList, nil
	case *CompoundT:
		return TagCompound, nil
	case *IntArrayT:
		return TagIntArray, nil
	case *LongArrayT:
		return TagLongArray, nil
	default:
		return byte('0'), errors.New(errorTag)
	}
}

// EndT to end type: 0
type EndT int

// ByteT to byte type: 1
type ByteT struct {
	name  string
	value byte
}

// ShortT to short type: 2
type ShortT struct {
	name  string
	value int16
}

// IntT to int type: 3
type IntT struct {
	name  string
	value int32
}

// LongT to long type: 4
type LongT struct {
	name  string
	value int64
}

// FloatT to float type: 5
type FloatT struct {
	name  string
	value float32
}

// DoubleT to double type: 6
type DoubleT struct {
	name  string
	value float64
}

// ByteArrayT to byte array type: 7
type ByteArrayT struct {
	name  string
	value []byte
}

// StringT to string type: 8
type StringT struct {
	name  string
	value string
}

// ListT to list type: 9
type ListT struct {
	name  string
	value []interface{}
}

// CompoundT to compound type: 10
type CompoundT struct {
	name  string
	value map[string]interface{}
}

// IntArrayT to int array type: 11
type IntArrayT struct {
	name  string
	value []int32
}

// LongArrayT to long array type: 12
type LongArrayT struct {
	name  string
	value []int64
}

// 1 		TAG_Byte 	1 byte / 8 bits, signed 	<number>b or <number>B 	A signed integral type. Sometimes used for booleans. 	Full range of -(27) to (27 - 1)
// (-128 to 127)
func (t *ByteT) Read(reader Reader) error {
	var err error

	if t.value, err = reader.Byte(); err != nil {
		return err
	}
	return nil
}

func (t *ByteT) Write(writer Writer, printInfo bool) error {
	var err error

	if printInfo {
		if err = writer.Byte(TagByte); err != nil {
			return err
		}
		if err = writer.String(t.name); err != nil {
			return err
		}
	}
	if err = writer.Byte(t.value); err != nil {
		return err
	}
	return nil
}

// 2 		TAG_Short 	2 bytes / 16 bits, signed, big endian 	<number>s or <number>S 	A signed integral type. 	Full range of -(215) to (215 - 1)
// (-32,768 to 32,767)
func (t *ShortT) Read(reader Reader) error {
	var err error

	if t.value, err = reader.Short(); err != nil {
		return err
	}
	return nil
}

func (t *ShortT) Write(writer Writer, printInfo bool) error {
	var err error

	if printInfo {
		if err = writer.Byte(TagShort); err != nil {
			return err
		}
		if err = writer.String(t.name); err != nil {
			return err
		}
	}
	if err = writer.Short(t.value); err != nil {
		return err
	}
	return nil
}

// 3 		TAG_Int 	4 bytes / 32 bits, signed, big endian 	<number> 	A signed integral type. 	Full range of -(231) to (231 - 1)
// (-2,147,483,648 to 2,147,483,647)
func (t *IntT) Read(reader Reader) error {
	var err error

	if t.value, err = reader.Int(); err != nil {
		return err
	}
	return nil
}

func (t *IntT) Write(writer Writer, printInfo bool) error {
	var err error

	if printInfo {
		if err = writer.Byte(TagInt); err != nil {
			return err
		}
		if err = writer.String(t.name); err != nil {
			return err
		}
	}
	if err = writer.Int(t.value); err != nil {
		return err
	}
	return nil
}

// 4 		TAG_Long 	8 bytes / 64 bits, signed, big endian 	<number>l or <number>L 	A signed integral type. 	Full range of -(263) to (263 - 1)
// (-9,223,372,036,854,775,808 to 9,223,372,036,854,775,807)
func (t *LongT) Read(reader Reader) error {
	var err error

	if t.value, err = reader.Long(); err != nil {
		return err
	}
	return nil
}

func (t *LongT) Write(writer Writer, printInfo bool) error {
	var err error

	if printInfo {
		if err = writer.Byte(TagLong); err != nil {
			return err
		}
		if err = writer.String(t.name); err != nil {
			return err
		}
	}
	if err = writer.Long(t.value); err != nil {
		return err
	}
	return nil
}

// 5 		TAG_Float 	4 bytes / 32 bits, signed, big endian, IEEE 754-2008, binary32 	<number>f or <number>F 	A signed floating point type. 	Precision varies throughout number line;
// See Single-precision floating-point format. Maximum value about 3.4*1038
func (t *FloatT) Read(reader Reader) error {
	var err error

	if t.value, err = reader.Float(); err != nil {
		return err
	}
	return nil
}

func (t *FloatT) Write(writer Writer, printInfo bool) error {
	var err error

	if printInfo {
		if err = writer.Byte(TagFloat); err != nil {
			return err
		}
		if err = writer.String(t.name); err != nil {
			return err
		}
	}
	if err = writer.Float(t.value); err != nil {
		return err
	}
	return nil
}

//  6 		TAG_Double 	8 bytes / 64 bits, signed, big endian, IEEE 754-2008, binary64 	<decimal number>, <number>d or <number>D 	A signed floating point type. 	Precision varies throughout number line;
// See Double-precision floating-point format. Maximum value about 1.8*10308
func (t *DoubleT) Read(reader Reader) error {
	var err error

	if t.value, err = reader.Double(); err != nil {
		return err
	}
	return nil
}

func (t *DoubleT) Write(writer Writer, printInfo bool) error {
	var err error

	if printInfo {
		if err = writer.Byte(TagDouble); err != nil {
			return err
		}
		if err = writer.String(t.name); err != nil {
			return err
		}
	}
	if err = writer.Double(t.value); err != nil {
		return err
	}
	return nil
}

// 7 		TAG_Byte_Array 	TAG_Int's payload size, then size TAG_Byte's payloads. 	[B;<byte>,<byte>,...] 	An array of bytes. 	Maximum number of elements ranges between (231 - 9) and (231 - 1) (2,147,483,639 and 2,147,483,647), depending on the specific JVM.
func (t *ByteArrayT) Read(reader Reader) error {
	var err error

	if t.value, err = reader.Bytes(); err != nil {
		return err
	}
	return nil
}

func (t *ByteArrayT) Write(writer Writer, printInfo bool) error {
	var err error

	if printInfo {
		if err = writer.Byte(TagByteArray); err != nil {
			return err
		}
		if err = writer.String(t.name); err != nil {
			return err
		}
	}
	if err = writer.Bytes(t.value); err != nil {
		return err
	}
	return nil
}

// 8 		TAG_String 	A TAG_Short-like, but instead unsigned[2] payload length, then a UTF-8 string resembled by length bytes. 	<a-zA-Z0-9 text>, "<text>" (" within needs to be escaped to \"), or '<text>' (' within needs to be escaped to \') 	A UTF-8 string. It has a size, rather than being null terminated. 	65,535 bytes interpretable as UTF-8 (see modified UTF-8 format; most commonly-used characters are a single byte).
func (t *StringT) Read(reader Reader) error {
	var err error

	if t.value, err = reader.String(); err != nil {
		return err
	}
	return nil
}

func (t *StringT) Write(writer Writer, printInfo bool) error {
	var err error

	if printInfo {
		if err = writer.Byte(TagString); err != nil {
			return err
		}
		if err = writer.String(t.name); err != nil {
			return err
		}
	}
	if err = writer.String(t.value); err != nil {
		return err
	}
	return nil
}

// 9 		TAG_List 	TAG_Byte's payload tagId, then TAG_Int's payload size, then size tags' payloads, all of type tagId. 	[<value>,<value>,...] 	A list of tag payloads, without repeated tag IDs or any tag names. 	Due to JVM limitations and the implementation of ArrayList, the maximum number of list elements is (231 - 9), or 2,147,483,639. Also note that List and Compound tags may not be nested beyond a depth of 512.
func (t *ListT) Read(reader Reader) error {
	var err error
	var tagT byte
	var nbr int32

	// get the the tagType
	if tagT, err = reader.Byte(); err != nil {
		return err
	}
	// get number element
	if nbr, err = reader.Int(); err != nil {
		return err
	}

	for i := int32(0); i < nbr; i++ {
		var elem Tag
		if elem, err = NewTag(tagT, ""); err != nil {
			return err
		}
		if err = elem.Read(reader); err != nil {
			return err
		}
		t.value = append(t.value, elem)
	}
	return nil
}

func (t *ListT) Write(writer Writer, printInfo bool) error {
	var err error
	var tagT byte
	var nbr int32

	if printInfo {
		if err = writer.Byte(TagList); err != nil {
			return err
		}
		if err = writer.String(t.name); err != nil {
			return err
		}
	}

	nbr = int32(len(t.value))
	if nbr > 0 {
		if _, ok := t.value[0].(Tag); !ok {
			return errors.New(errorTag)
		}
		if tagT, err = TagType(t.value[0].(Tag)); err != nil {
			return err
		}
	}
	if err = writer.Byte(tagT); err != nil {
		return err
	}
	if err = writer.Int(nbr); err != nil {
		return err
	}

	for _, t := range t.value {
		if err = t.(Tag).Write(writer, false); err != nil {
			return err
		}
	}
	return nil
}

// 10 		TAG_Compound 	Fully formed tags, followed by a TAG_End. 	{<tag name>:<value>,<tag name>:<value>,...} 	A list of fully formed tags, including their IDs, names, and payloads. No two tags may have the same name. 	Unlike lists, there is no hard limit to the number of tags within a Compound (of course, there is always the implicit limit of virtual memory). Note, however, that Compound and List tags may not be nested beyond a depth of 512.
func (t *CompoundT) Read(reader Reader) error {
	var err error
	var tagT byte
	var name string

	t.value = make(map[string]interface{})
	for tagT, err = reader.Byte(); tagT != TagEnd && err == nil; tagT, err = reader.Byte() {
		if name, err = reader.String(); err != nil {
			return err
		}
		var elem Tag
		if elem, err = NewTag(tagT, name); err != nil {
			return err
		}
		if err = elem.Read(reader); err != nil {
			return err
		}
		t.value[name] = elem
	}
	if err != nil {
		return err
	}
	return nil
}

func (t *CompoundT) Write(writer Writer, printInfo bool) error {
	var err error

	if printInfo {
		if err = writer.Byte(TagCompound); err != nil {
			return err
		}
		if err = writer.String(t.name); err != nil {
			return err
		}
	}
	for key, value := range t.value {
		var tagT byte

		if _, ok := value.(Tag); !ok {
			return errors.New(errorTag)
		}
		if tagT, err = TagType(value.(Tag)); err != nil {
			return err
		}
		if err = writer.Byte(tagT); err != nil {
			return err
		}
		if err = writer.String(key); err != nil {
			return err
		}
		if err = value.(Tag).Write(writer, false); err != nil {
			return err
		}
	}
	if err = writer.Byte(TagEnd); err != nil {
		return err
	}
	return nil
}

// 11 		TAG_Int_Array 	TAG_Int's payload size, then size TAG_Int's payloads. 	[I;<integer>,<integer>,...] 	An array of TAG_Int's payloads. 	Maximum number of elements ranges between (231 - 9) and (231 - 1) (2,147,483,639 and 2,147,483,647), depending on the specific JVM.
func (t *IntArrayT) Read(reader Reader) error {
	var err error

	if t.value, err = reader.IntArray(); err != nil {
		return err
	}
	return nil
}

func (t *IntArrayT) Write(writer Writer, printInfo bool) error {
	var err error

	if printInfo {
		if err = writer.Byte(TagIntArray); err != nil {
			return err
		}
		if err = writer.String(t.name); err != nil {
			return err
		}
	}
	if err = writer.Int(int32(len(t.value))); err != nil {
		return err
	}
	for _, v := range t.value {
		if err = writer.Int(v); err != nil {
			return err
		}
	}
	return nil
}

// 12 		TAG_Long_Array 	TAG_
func (t *LongArrayT) Read(reader Reader) error {
	var err error

	if t.value, err = reader.LongArray(); err != nil {
		return err
	}
	return nil
}

func (t *LongArrayT) Write(writer Writer, printInfo bool) error {
	var err error

	if printInfo {
		if err = writer.Byte(TagLongArray); err != nil {
			return err
		}
		if err = writer.String(t.name); err != nil {
			return err
		}
	}
	if err = writer.Long(int64(len(t.value))); err != nil {
		return err
	}
	for _, v := range t.value {
		if err = writer.Long(v); err != nil {
			return err
		}
	}
	return nil
}
