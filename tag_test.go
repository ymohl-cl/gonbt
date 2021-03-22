package gonbt

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// fakeTag to provide a unsupported Tag in test
type fakeTag struct {
	name string
}

func (f *fakeTag) Read(Reader) error        { return nil }
func (f *fakeTag) Write(Writer, bool) error { return nil }

func TestNewTag(t *testing.T) {
	tagName := "testing_name"
	t.Run("should be return an error because tag unknow", func(t *testing.T) {
		expectedError := errorTag

		tag, err := NewTag(byte(42), tagName)
		if assert.Error(t, err); err != nil {
			assert.EqualValues(t, expectedError, err.Error())
			assert.Nil(t, tag)
		}
	})
	t.Run("should be ok with tag END", func(t *testing.T) {
		expectedError := errorEnd

		tag, err := NewTag(TagEnd, tagName)
		if assert.Error(t, err); err != nil {
			assert.EqualValues(t, expectedError, err.Error())
			assert.Nil(t, tag)
		}
	})
	t.Run("should be ok with tag Byte", func(t *testing.T) {
		expectedTag := &ByteT{name: tagName}

		tag, err := NewTag(TagByte, tagName)
		if assert.NoError(t, err); err != nil {
			assert.Nil(t, expectedTag, tag)
		}
	})
	t.Run("should be ok with tag Short", func(t *testing.T) {
		expectedTag := &ShortT{name: tagName}

		tag, err := NewTag(TagShort, tagName)
		if assert.NoError(t, err); err != nil {
			assert.Nil(t, expectedTag, tag)
		}
	})
	t.Run("should be ok with tag Int", func(t *testing.T) {
		expectedTag := &IntT{name: tagName}

		tag, err := NewTag(TagInt, tagName)
		if assert.NoError(t, err); err != nil {
			assert.Nil(t, expectedTag, tag)
		}
	})
	t.Run("should be ok with tag Long", func(t *testing.T) {
		expectedTag := &LongT{name: tagName}

		tag, err := NewTag(TagLong, tagName)
		if assert.NoError(t, err); err != nil {
			assert.Nil(t, expectedTag, tag)
		}
	})
	t.Run("should be ok with tag Float", func(t *testing.T) {
		expectedTag := &FloatT{name: tagName}

		tag, err := NewTag(TagFloat, tagName)
		if assert.NoError(t, err); err != nil {
			assert.Nil(t, expectedTag, tag)
		}
	})
	t.Run("should be ok with tag Double", func(t *testing.T) {
		expectedTag := &DoubleT{name: tagName}

		tag, err := NewTag(TagDouble, tagName)
		if assert.NoError(t, err); err != nil {
			assert.Nil(t, expectedTag, tag)
		}
	})
	t.Run("should be ok with tag Byte array", func(t *testing.T) {
		expectedTag := &ByteArrayT{name: tagName}

		tag, err := NewTag(TagByteArray, tagName)
		if assert.NoError(t, err); err != nil {
			assert.Nil(t, expectedTag, tag)
		}
	})
	t.Run("should be ok with tag String", func(t *testing.T) {
		expectedTag := &StringT{name: tagName}

		tag, err := NewTag(TagString, tagName)
		if assert.NoError(t, err); err != nil {
			assert.Nil(t, expectedTag, tag)
		}
	})
	t.Run("should be ok with tag List", func(t *testing.T) {
		expectedTag := &ListT{name: tagName}

		tag, err := NewTag(TagList, tagName)
		if assert.NoError(t, err); err != nil {
			assert.Nil(t, expectedTag, tag)
		}
	})
	t.Run("should be ok with tag Compound", func(t *testing.T) {
		expectedTag := &CompoundT{name: tagName}

		tag, err := NewTag(TagCompound, tagName)
		if assert.NoError(t, err); err != nil {
			assert.Nil(t, expectedTag, tag)
		}
	})
	t.Run("should be ok with tag Int Array", func(t *testing.T) {
		expectedTag := &IntArrayT{name: tagName}

		tag, err := NewTag(TagIntArray, tagName)
		if assert.NoError(t, err); err != nil {
			assert.Nil(t, expectedTag, tag)
		}
	})
	t.Run("should be ok with tag Long Array", func(t *testing.T) {
		expectedTag := &LongArrayT{name: tagName}

		tag, err := NewTag(TagLongArray, tagName)
		if assert.NoError(t, err); err != nil {
			assert.Nil(t, expectedTag, tag)
		}
	})
}

func TestTagType(t *testing.T) {
	t.Run("should return an error because tag unknow", func(t *testing.T) {
		expectedErr := errorTag
		expectedTagType := byte('0')

		tagT, err := TagType(&fakeTag{})
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedErr, err.Error())
			assert.EqualValues(t, expectedTagType, tagT)
		}
	})
	t.Run("should return an error because tag is nil", func(t *testing.T) {
		expectedErr := errorTag
		expectedTagType := byte('0')

		tagT, err := TagType(nil)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedErr, err.Error())
			assert.EqualValues(t, expectedTagType, tagT)
		}
	})
	t.Run("should be ok with the tag Byte", func(t *testing.T) {
		expectedTagType := TagByte

		tagT, err := TagType(&ByteT{})
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedTagType, tagT)
		}
	})
	t.Run("should be ok with the tag Short", func(t *testing.T) {
		expectedTagType := TagShort

		tagT, err := TagType(&ShortT{})
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedTagType, tagT)
		}
	})
	t.Run("should be ok with the tag Int", func(t *testing.T) {
		expectedTagType := TagInt

		tagT, err := TagType(&IntT{})
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedTagType, tagT)
		}
	})
	t.Run("should be ok with the tag Long", func(t *testing.T) {
		expectedTagType := TagLong

		tagT, err := TagType(&LongT{})
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedTagType, tagT)
		}
	})
	t.Run("should be ok with the tag Float", func(t *testing.T) {
		expectedTagType := TagFloat

		tagT, err := TagType(&FloatT{})
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedTagType, tagT)
		}
	})
	t.Run("should be ok with the tag Double", func(t *testing.T) {
		expectedTagType := TagDouble

		tagT, err := TagType(&DoubleT{})
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedTagType, tagT)
		}
	})
	t.Run("should be ok with the tag Byte array", func(t *testing.T) {
		expectedTagType := TagByteArray

		tagT, err := TagType(&ByteArrayT{})
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedTagType, tagT)
		}
	})
	t.Run("should be ok with the tag String", func(t *testing.T) {
		expectedTagType := TagString

		tagT, err := TagType(&StringT{})
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedTagType, tagT)
		}
	})
	t.Run("should be ok with the tag List", func(t *testing.T) {
		expectedTagType := TagList

		tagT, err := TagType(&ListT{})
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedTagType, tagT)
		}
	})
	t.Run("should be ok with the tag Compound", func(t *testing.T) {
		expectedTagType := TagCompound

		tagT, err := TagType(&CompoundT{})
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedTagType, tagT)
		}
	})
	t.Run("should be ok with the tag int array", func(t *testing.T) {
		expectedTagType := TagIntArray

		tagT, err := TagType(&IntArrayT{})
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedTagType, tagT)
		}
	})
	t.Run("should be ok with the tag Long array", func(t *testing.T) {
		expectedTagType := TagLongArray

		tagT, err := TagType(&LongArrayT{})
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedTagType, tagT)
		}
	})
}

func TestByteT_Read(t *testing.T) {
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mreader := NewMockReader(ctrl)

	t.Run("Should return an error because the reader.Byte() failed", func(t *testing.T) {
		tag := &ByteT{}
		expectedValue := byte('0')

		mreader.EXPECT().Byte().Return(byte('0'), errors.New(expectedMockErr))
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &ByteT{}
		expectedValue := byte('A')

		mreader.EXPECT().Byte().Return(byte('A'), nil)
		err := tag.Read(mreader)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
}

func TestByteT_Write(t *testing.T) {
	tagName := "tag_name"
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mwriter := NewMockWriter(ctrl)

	t.Run("Should return an error because the first call to writer.Byte failed", func(t *testing.T) {
		tag := &ByteT{}

		mwriter.EXPECT().Byte(gomock.Eq(TagByte)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the call to writer.String failed", func(t *testing.T) {
		tag := &ByteT{name: tagName}

		mwriter.EXPECT().Byte(gomock.Eq(TagByte)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the final writer call failed", func(t *testing.T) {
		tag := &ByteT{name: tagName, value: byte('A')}

		mwriter.EXPECT().Byte(gomock.Eq(TagByte)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Byte(gomock.Eq(byte('A'))).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &ByteT{name: tagName, value: byte('A')}

		mwriter.EXPECT().Byte(gomock.Eq(TagByte)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Byte(gomock.Eq(byte('A'))).Return(nil)
		err := tag.Write(mwriter, true)
		assert.NoError(t, err)
	})
}

func TestShortT_Read(t *testing.T) {
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mreader := NewMockReader(ctrl)

	t.Run("Should return an error because the reader.Short() failed", func(t *testing.T) {
		tag := &ShortT{}
		expectedValue := int16(0)

		mreader.EXPECT().Short().Return(int16(0), errors.New(expectedMockErr))
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &ShortT{}
		expectedValue := int16(42)

		mreader.EXPECT().Short().Return(int16(42), nil)
		err := tag.Read(mreader)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
}

func TestShortT_Write(t *testing.T) {
	tagName := "tag_name"
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mwriter := NewMockWriter(ctrl)

	t.Run("Should return an error because the first call to writer.Byte failed", func(t *testing.T) {
		tag := &ShortT{}

		mwriter.EXPECT().Byte(gomock.Eq(TagShort)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the call to writer.String failed", func(t *testing.T) {
		tag := &ShortT{name: tagName}

		mwriter.EXPECT().Byte(gomock.Eq(TagShort)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the final writer call failed", func(t *testing.T) {
		tag := &ShortT{name: tagName, value: int16(42)}

		mwriter.EXPECT().Byte(gomock.Eq(TagShort)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Short(gomock.Eq(int16(42))).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &ShortT{name: tagName, value: int16(42)}

		mwriter.EXPECT().Byte(gomock.Eq(TagShort)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Short(gomock.Eq(int16(42))).Return(nil)
		err := tag.Write(mwriter, true)
		assert.NoError(t, err)
	})
}

func TestIntT_Read(t *testing.T) {
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mreader := NewMockReader(ctrl)

	t.Run("Should return an error because the reader.Int() failed", func(t *testing.T) {
		tag := &IntT{}
		expectedValue := int32(0)

		mreader.EXPECT().Int().Return(int32(0), errors.New(expectedMockErr))
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &IntT{}
		expectedValue := int32(42)

		mreader.EXPECT().Int().Return(int32(42), nil)
		err := tag.Read(mreader)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
}

func TestIntT_Write(t *testing.T) {
	tagName := "tag_name"
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mwriter := NewMockWriter(ctrl)

	t.Run("Should return an error because the first call to writer.Byte failed", func(t *testing.T) {
		tag := &IntT{}

		mwriter.EXPECT().Byte(gomock.Eq(TagInt)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the call to writer.String failed", func(t *testing.T) {
		tag := &IntT{name: tagName}

		mwriter.EXPECT().Byte(gomock.Eq(TagInt)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the final writer call failed", func(t *testing.T) {
		tag := &IntT{name: tagName, value: int32(42)}

		mwriter.EXPECT().Byte(gomock.Eq(TagInt)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Int(gomock.Eq(int32(42))).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &IntT{name: tagName, value: int32(42)}

		mwriter.EXPECT().Byte(gomock.Eq(TagInt)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Int(gomock.Eq(int32(42))).Return(nil)
		err := tag.Write(mwriter, true)
		assert.NoError(t, err)
	})
}

func TestLongT_Read(t *testing.T) {
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mreader := NewMockReader(ctrl)

	t.Run("Should return an error because the reader.Long() failed", func(t *testing.T) {
		tag := &LongT{}
		expectedValue := int64(0)

		mreader.EXPECT().Long().Return(int64(0), errors.New(expectedMockErr))
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &LongT{}
		expectedValue := int64(42)

		mreader.EXPECT().Long().Return(int64(42), nil)
		err := tag.Read(mreader)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
}

func TestLongT_Write(t *testing.T) {
	tagName := "tag_name"
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mwriter := NewMockWriter(ctrl)

	t.Run("Should return an error because the first call to writer.Byte failed", func(t *testing.T) {
		tag := &LongT{}

		mwriter.EXPECT().Byte(gomock.Eq(TagLong)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the call to writer.String failed", func(t *testing.T) {
		tag := &LongT{name: tagName}

		mwriter.EXPECT().Byte(gomock.Eq(TagLong)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the final writer call failed", func(t *testing.T) {
		tag := &LongT{name: tagName, value: int64(42)}

		mwriter.EXPECT().Byte(gomock.Eq(TagLong)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Long(gomock.Eq(int64(42))).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &LongT{name: tagName, value: int64(42)}

		mwriter.EXPECT().Byte(gomock.Eq(TagLong)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Long(gomock.Eq(int64(42))).Return(nil)
		err := tag.Write(mwriter, true)
		assert.NoError(t, err)
	})
}

func TestFloatT_Read(t *testing.T) {
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mreader := NewMockReader(ctrl)

	t.Run("Should return an error because the reader.Float() failed", func(t *testing.T) {
		tag := &FloatT{}
		expectedValue := float32(0)

		mreader.EXPECT().Float().Return(float32(0), errors.New(expectedMockErr))
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &FloatT{}
		expectedValue := float32(42)

		mreader.EXPECT().Float().Return(float32(42), nil)
		err := tag.Read(mreader)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
}

func TestFloatT_Write(t *testing.T) {
	tagName := "tag_name"
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mwriter := NewMockWriter(ctrl)

	t.Run("Should return an error because the first call to writer.Byte failed", func(t *testing.T) {
		tag := &FloatT{}

		mwriter.EXPECT().Byte(gomock.Eq(TagFloat)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the call to writer.String failed", func(t *testing.T) {
		tag := &FloatT{name: tagName}

		mwriter.EXPECT().Byte(gomock.Eq(TagFloat)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the final writer call failed", func(t *testing.T) {
		tag := &FloatT{name: tagName, value: float32(42)}

		mwriter.EXPECT().Byte(gomock.Eq(TagFloat)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Float(gomock.Eq(float32(42))).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &FloatT{name: tagName, value: float32(42)}

		mwriter.EXPECT().Byte(gomock.Eq(TagFloat)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Float(gomock.Eq(float32(42))).Return(nil)
		err := tag.Write(mwriter, true)
		assert.NoError(t, err)
	})
}

func TestDoubleT_Read(t *testing.T) {
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mreader := NewMockReader(ctrl)

	t.Run("Should return an error because the reader.Double() failed", func(t *testing.T) {
		tag := &DoubleT{}
		expectedValue := float64(0)

		mreader.EXPECT().Double().Return(float64(0), errors.New(expectedMockErr))
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &DoubleT{}
		expectedValue := float64(42)

		mreader.EXPECT().Double().Return(float64(42), nil)
		err := tag.Read(mreader)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
}

func TestDoubleT_Write(t *testing.T) {
	tagName := "tag_name"
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mwriter := NewMockWriter(ctrl)

	t.Run("Should return an error because the first call to writer.Byte failed", func(t *testing.T) {
		tag := &DoubleT{}

		mwriter.EXPECT().Byte(gomock.Eq(TagDouble)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the call to writer.String failed", func(t *testing.T) {
		tag := &DoubleT{name: tagName}

		mwriter.EXPECT().Byte(gomock.Eq(TagDouble)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the final writer call failed", func(t *testing.T) {
		tag := &DoubleT{name: tagName, value: float64(42)}

		mwriter.EXPECT().Byte(gomock.Eq(TagDouble)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Double(gomock.Eq(float64(42))).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &DoubleT{name: tagName, value: float64(42)}

		mwriter.EXPECT().Byte(gomock.Eq(TagDouble)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Double(gomock.Eq(float64(42))).Return(nil)
		err := tag.Write(mwriter, true)
		assert.NoError(t, err)
	})
}

func TestByteArrayT_Read(t *testing.T) {
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mreader := NewMockReader(ctrl)

	t.Run("Should return an error because the reader.ByteArray() failed", func(t *testing.T) {
		tag := &ByteArrayT{}
		expectedValue := []byte{}

		mreader.EXPECT().Bytes().Return([]byte{}, errors.New(expectedMockErr))
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &ByteArrayT{}
		expectedValue := []byte{0x0a, 0x0a}

		mreader.EXPECT().Bytes().Return([]byte{0x0a, 0x0a}, nil)
		err := tag.Read(mreader)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
}

func TestByteArrayT_Write(t *testing.T) {
	tagName := "tag_name"
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mwriter := NewMockWriter(ctrl)

	t.Run("Should return an error because the first call to writer.Byte failed", func(t *testing.T) {
		tag := &ByteArrayT{}

		mwriter.EXPECT().Byte(gomock.Eq(TagByteArray)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the call to writer.String failed", func(t *testing.T) {
		tag := &ByteArrayT{name: tagName}

		mwriter.EXPECT().Byte(gomock.Eq(TagByteArray)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the final writer call failed", func(t *testing.T) {
		tag := &ByteArrayT{name: tagName, value: []byte{0x0a, 0x0}}

		mwriter.EXPECT().Byte(gomock.Eq(TagByteArray)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Bytes(gomock.Eq([]byte{0x0a, 0x0})).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &ByteArrayT{name: tagName, value: []byte{0x0a, 0x0}}

		mwriter.EXPECT().Byte(gomock.Eq(TagByteArray)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Bytes(gomock.Eq([]byte{0x0a, 0x0})).Return(nil)
		err := tag.Write(mwriter, true)
		assert.NoError(t, err)
	})
}

func TestStringT_Read(t *testing.T) {
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mreader := NewMockReader(ctrl)

	t.Run("Should return an error because the reader.String() failed", func(t *testing.T) {
		tag := &StringT{}
		expectedValue := ""

		mreader.EXPECT().String().Return("", errors.New(expectedMockErr))
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &StringT{}
		expectedValue := "Hello world !"

		mreader.EXPECT().String().Return("Hello world !", nil)
		err := tag.Read(mreader)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
}

func TestStringT_Write(t *testing.T) {
	tagName := "tag_name"
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mwriter := NewMockWriter(ctrl)

	t.Run("Should return an error because the first call to writer.Byte failed", func(t *testing.T) {
		tag := &StringT{}

		mwriter.EXPECT().Byte(gomock.Eq(TagString)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the call to writer.String failed", func(t *testing.T) {
		tag := &StringT{name: tagName}

		mwriter.EXPECT().Byte(gomock.Eq(TagString)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the final writer call failed", func(t *testing.T) {
		tag := &StringT{name: tagName, value: "Hello World !"}

		mwriter.EXPECT().Byte(gomock.Eq(TagString)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq("Hello World !")).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &StringT{name: tagName, value: "Hello World !"}

		mwriter.EXPECT().Byte(gomock.Eq(TagString)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq("Hello World !")).Return(nil)
		err := tag.Write(mwriter, true)
		assert.NoError(t, err)
	})
}

func TestListT_Read(t *testing.T) {
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mreader := NewMockReader(ctrl)

	t.Run("Should return an error because the reader.Byte() failed", func(t *testing.T) {
		tag := &ListT{}

		mreader.EXPECT().Byte().Return(byte('0'), errors.New(expectedMockErr))
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
			assert.Empty(t, tag.value)
		}
	})
	t.Run("Should return an error because the reader.Int() failed", func(t *testing.T) {
		tag := &ListT{}

		mreader.EXPECT().Byte().Return(byte(TagFloat), nil)
		mreader.EXPECT().Int().Return(int32(0), errors.New(expectedMockErr))
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
			assert.Empty(t, tag.value)
		}
	})
	t.Run("Should return an error because NewTag() failed", func(t *testing.T) {
		tag := &ListT{}

		mreader.EXPECT().Byte().Return(byte(42), nil)
		mreader.EXPECT().Int().Return(int32(3), nil)
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, errorTag, err.Error())
			assert.Empty(t, tag.value)
		}
	})
	t.Run("Should return an error because elem.Read() failed", func(t *testing.T) {
		tag := &ListT{}

		mreader.EXPECT().Byte().Return(byte(TagString), nil)
		mreader.EXPECT().Int().Return(int32(3), nil)
		mreader.EXPECT().String().Return("", errors.New(expectedMockErr))
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
			assert.Empty(t, tag.value)
		}
	})
	t.Run("Should be ok with a complete list", func(t *testing.T) {
		tag := &ListT{}
		expectedValue := []interface{}{&StringT{value: "coucou"}, &StringT{value: "Hello"}, &StringT{value: "Yo"}}

		mreader.EXPECT().Byte().Return(byte(TagString), nil)
		mreader.EXPECT().Int().Return(int32(3), nil)
		mreader.EXPECT().String().Return("coucou", nil).Times(1)
		mreader.EXPECT().String().Return("Hello", nil).Times(1)
		mreader.EXPECT().String().Return("Yo", nil).Times(1)
		err := tag.Read(mreader)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
}

func TestListT_Write(t *testing.T) {
	tagName := "tag_name"
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mwriter := NewMockWriter(ctrl)

	t.Run("Should return an error because the first call to writer.Byte failed", func(t *testing.T) {
		tag := &ListT{}

		mwriter.EXPECT().Byte(gomock.Eq(TagList)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the call to writer.String failed", func(t *testing.T) {
		tag := &ListT{name: tagName}

		mwriter.EXPECT().Byte(gomock.Eq(TagList)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the type list is not supported", func(t *testing.T) {
		tag := &ListT{name: tagName, value: []interface{}{"primitive", "type", "not", "supported", "without", "tagparent type"}}

		mwriter.EXPECT().Byte(gomock.Eq(TagList)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, errorTag, err.Error())
		}
	})
	t.Run("Should return an error because TagType function return an error", func(t *testing.T) {
		tag := &ListT{name: tagName, value: []interface{}{&fakeTag{name: "fake_type"}}}

		mwriter.EXPECT().Byte(gomock.Eq(TagList)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, errorTag, err.Error())
		}
	})
	t.Run("Should return an error because second call to writer.Byte() failed", func(t *testing.T) {
		tag := &ListT{name: tagName, value: []interface{}{&StringT{value: "coucou"}, &StringT{value: "Hello"}, &StringT{value: "Yo"}}}

		mwriter.EXPECT().Byte(gomock.Eq(TagList)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Byte(gomock.Eq(TagString)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because second call to writer.Int() failed", func(t *testing.T) {
		tag := &ListT{name: tagName, value: []interface{}{&StringT{value: "coucou"}, &StringT{value: "Hello"}, &StringT{value: "Yo"}}}

		mwriter.EXPECT().Byte(gomock.Eq(TagList)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Byte(gomock.Eq(TagString)).Return(nil)
		mwriter.EXPECT().Int(gomock.Eq(int32(3))).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because embeded elements Writer failed", func(t *testing.T) {
		tag := &ListT{name: tagName, value: []interface{}{&StringT{value: "coucou"}, &StringT{value: "Hello"}, &StringT{value: "Yo"}}}

		mwriter.EXPECT().Byte(gomock.Eq(TagList)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Byte(gomock.Eq(TagString)).Return(nil)
		mwriter.EXPECT().Int(gomock.Eq(int32(3))).Return(nil)
		mwriter.EXPECT().String(gomock.Eq("coucou")).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})

	t.Run("Should be ok", func(t *testing.T) {
		tag := &ListT{name: tagName, value: []interface{}{&StringT{value: "coucou"}, &StringT{value: "Hello"}, &StringT{value: "Yo"}}}

		mwriter.EXPECT().Byte(gomock.Eq(TagList)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Byte(gomock.Eq(TagString)).Return(nil)
		mwriter.EXPECT().Int(gomock.Eq(int32(3))).Return(nil)
		mwriter.EXPECT().String(gomock.Eq("coucou")).Return(nil).Times(1)
		mwriter.EXPECT().String(gomock.Eq("Hello")).Return(nil).Times(1)
		mwriter.EXPECT().String(gomock.Eq("Yo")).Return(nil).Times(1)
		err := tag.Write(mwriter, true)
		assert.NoError(t, err)
	})
	t.Run("Should be ok with an empty list", func(t *testing.T) {
		tag := &ListT{name: tagName, value: []interface{}{}}

		mwriter.EXPECT().Byte(gomock.Eq(TagList)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Byte(gomock.Eq(byte(0x00))).Return(nil)
		mwriter.EXPECT().Int(gomock.Eq(int32(0))).Return(nil)
		err := tag.Write(mwriter, true)
		assert.NoError(t, err)
	})
}

func TestCompoundT_Read(t *testing.T) {
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mreader := NewMockReader(ctrl)

	t.Run("Should return an error because the reader.Byte() failed", func(t *testing.T) {
		tag := &CompoundT{}
		expectedValue := map[string]interface{}{}

		mreader.EXPECT().Byte().Return(byte('0'), errors.New(expectedMockErr))
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})

	t.Run("Should return an error becaue reader.String() failed", func(t *testing.T) {
		tag := &CompoundT{}
		expectedValue := map[string]interface{}{}

		mreader.EXPECT().Byte().Return(TagString, nil)
		mreader.EXPECT().String().Return("", errors.New(expectedMockErr))
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
	t.Run("Should return an error becaue NewTag failed", func(t *testing.T) {
		tag := &CompoundT{}
		expectedValue := map[string]interface{}{}

		mreader.EXPECT().Byte().Return(byte(0x2a), nil)
		mreader.EXPECT().String().Return("tag_name", nil)
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, errorTag, err.Error())
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
	t.Run("Should return an error becaue elem.Read failed", func(t *testing.T) {
		tag := &CompoundT{}
		expectedValue := map[string]interface{}{}

		mreader.EXPECT().Byte().Return(byte(TagString), nil)
		mreader.EXPECT().String().Return("tag_name", nil)
		mreader.EXPECT().String().Return("", errors.New(expectedMockErr))
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
	t.Run("Should be ok with an empty list", func(t *testing.T) {
		tag := &CompoundT{}
		expectedValue := map[string]interface{}{}

		mreader.EXPECT().Byte().Return(byte(TagEnd), nil)
		err := tag.Read(mreader)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &CompoundT{}
		expectedValue := map[string]interface{}{
			"tag_name1": &StringT{name: "tag_name1", value: "coucou1"},
			"tag_name2": &StringT{name: "tag_name2", value: "coucou2"},
		}

		mreader.EXPECT().Byte().Return(byte(TagString), nil)
		mreader.EXPECT().String().Return("tag_name1", nil)
		mreader.EXPECT().String().Return("coucou1", nil)
		mreader.EXPECT().Byte().Return(byte(TagString), nil)
		mreader.EXPECT().String().Return("tag_name2", nil)
		mreader.EXPECT().String().Return("coucou2", nil)
		mreader.EXPECT().Byte().Return(byte(TagEnd), nil)
		err := tag.Read(mreader)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})

}

func TestCompoundT_Write(t *testing.T) {
	tagName := "tag_name"
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mwriter := NewMockWriter(ctrl)

	t.Run("Should return an error because the first call to writer.Byte failed", func(t *testing.T) {
		tag := &CompoundT{}

		mwriter.EXPECT().Byte(gomock.Eq(TagCompound)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the call to writer.String failed", func(t *testing.T) {
		tag := &CompoundT{name: tagName}

		mwriter.EXPECT().Byte(gomock.Eq(TagCompound)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the content value is not a tag", func(t *testing.T) {
		tag := &CompoundT{name: tagName, value: map[string]interface{}{
			"tag_name1": "hello"},
		}

		mwriter.EXPECT().Byte(gomock.Eq(TagCompound)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, errorTag, err.Error())
		}
	})
	t.Run("Should return an error because the content value is tag unsuported", func(t *testing.T) {
		tag := &CompoundT{name: tagName, value: map[string]interface{}{
			"tag_name1": &fakeTag{name: "tag_name1"}},
		}

		mwriter.EXPECT().Byte(gomock.Eq(TagCompound)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, errorTag, err.Error())
		}
	})
	t.Run("Should return an error because the writer tag type by writer.Byte() call failed", func(t *testing.T) {
		tag := &CompoundT{name: tagName, value: map[string]interface{}{
			"tag_name1": &StringT{name: "tag_name1", value: "coucou1"},
		}}

		mwriter.EXPECT().Byte(gomock.Eq(TagCompound)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Byte(gomock.Eq(TagString)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the writer key by writer.String() call failed", func(t *testing.T) {
		tag := &CompoundT{name: tagName, value: map[string]interface{}{
			"tag_name1": &StringT{name: "tag_name1", value: "coucou1"},
		}}

		mwriter.EXPECT().Byte(gomock.Eq(TagCompound)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Byte(gomock.Eq(TagString)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq("tag_name1")).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the elem.Writer() call failed", func(t *testing.T) {
		tag := &CompoundT{name: tagName, value: map[string]interface{}{
			"tag_name1": &StringT{name: "tag_name1", value: "coucou1"},
		}}

		mwriter.EXPECT().Byte(gomock.Eq(TagCompound)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		// write the embbeded element (type string)
		mwriter.EXPECT().Byte(gomock.Eq(TagString)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq("tag_name1")).Return(nil)
		mwriter.EXPECT().String(gomock.Eq("coucou1")).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the write.Byte() to TagEnd failed", func(t *testing.T) {
		tag := &CompoundT{name: tagName, value: map[string]interface{}{
			"tag_name1": &StringT{name: "tag_name1", value: "coucou1"},
		}}

		mwriter.EXPECT().Byte(gomock.Eq(TagCompound)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		// write the embbeded element (type string)
		// elem 1
		mwriter.EXPECT().Byte(gomock.Eq(TagString)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq("tag_name1")).Return(nil)
		mwriter.EXPECT().String(gomock.Eq("coucou1")).Return(nil)
		// tag end writing
		mwriter.EXPECT().Byte(gomock.Eq(TagEnd)).Return(errors.New(expectedMockErr))

		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})

	t.Run("Should be ok", func(t *testing.T) {
		tag := &CompoundT{name: tagName, value: map[string]interface{}{
			"tag_name1": &StringT{name: "tag_name1", value: "coucou1"},
		}}

		mwriter.EXPECT().Byte(gomock.Eq(TagCompound)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		// write the embbeded element (type string)
		// elem 1
		mwriter.EXPECT().Byte(gomock.Eq(TagString)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq("tag_name1")).Return(nil)
		mwriter.EXPECT().String(gomock.Eq("coucou1")).Return(nil)
		// tag end writing
		mwriter.EXPECT().Byte(gomock.Eq(TagEnd)).Return(nil)

		err := tag.Write(mwriter, true)
		assert.NoError(t, err)
	})
	t.Run("Should be ok with an empty value", func(t *testing.T) {
		tag := &CompoundT{name: tagName, value: map[string]interface{}{}}

		mwriter.EXPECT().Byte(gomock.Eq(TagCompound)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		// tag end writing
		mwriter.EXPECT().Byte(gomock.Eq(TagEnd)).Return(nil)

		err := tag.Write(mwriter, true)
		assert.NoError(t, err)
	})

}

func TestIntArrayT_Read(t *testing.T) {
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mreader := NewMockReader(ctrl)

	t.Run("Should return an error because the reader.IntArray() failed", func(t *testing.T) {
		tag := &IntArrayT{}
		expectedValue := []int32{}

		mreader.EXPECT().IntArray().Return([]int32{}, errors.New(expectedMockErr))
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &IntArrayT{}
		expectedValue := []int32{11, 42, 59}

		mreader.EXPECT().IntArray().Return([]int32{11, 42, 59}, nil)
		err := tag.Read(mreader)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
}

func TestIntArrayT_Write(t *testing.T) {
	tagName := "tag_name"
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mwriter := NewMockWriter(ctrl)

	t.Run("Should return an error because the first call to writer.Byte failed", func(t *testing.T) {
		tag := &IntArrayT{}

		mwriter.EXPECT().Byte(gomock.Eq(TagIntArray)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the call to writer.String failed", func(t *testing.T) {
		tag := &IntArrayT{name: tagName}

		mwriter.EXPECT().Byte(gomock.Eq(TagIntArray)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the first call to writer.Int failed", func(t *testing.T) {
		tag := &IntArrayT{name: tagName}

		mwriter.EXPECT().Byte(gomock.Eq(TagIntArray)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Int(gomock.Eq(int32(0))).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the second call to writer.Int failed", func(t *testing.T) {
		tag := &IntArrayT{name: tagName, value: []int32{42}}

		mwriter.EXPECT().Byte(gomock.Eq(TagIntArray)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Int(gomock.Eq(int32(1))).Return(nil)
		mwriter.EXPECT().Int(gomock.Eq(int32(42))).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &IntArrayT{name: tagName, value: []int32{42, 3, 33}}
		expectedValue := []int32{42, 3, 33}

		mwriter.EXPECT().Byte(gomock.Eq(TagIntArray)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Int(gomock.Eq(int32(3))).Return(nil)
		mwriter.EXPECT().Int(gomock.Eq(int32(42))).Return(nil)
		mwriter.EXPECT().Int(gomock.Eq(int32(3))).Return(nil)
		mwriter.EXPECT().Int(gomock.Eq(int32(33))).Return(nil)
		err := tag.Write(mwriter, true)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
	t.Run("Should be ok with empty list", func(t *testing.T) {
		tag := &IntArrayT{name: tagName, value: []int32{}}
		expectedValue := []int32{}

		mwriter.EXPECT().Byte(gomock.Eq(TagIntArray)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Int(gomock.Eq(int32(0))).Return(nil)
		err := tag.Write(mwriter, true)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
}

func TestLongArrayT_Read(t *testing.T) {
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mreader := NewMockReader(ctrl)

	t.Run("Should return an error because the reader.LongArray() failed", func(t *testing.T) {
		tag := &LongArrayT{}
		expectedValue := []int64{}

		mreader.EXPECT().LongArray().Return([]int64{}, errors.New(expectedMockErr))
		err := tag.Read(mreader)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &LongArrayT{}
		expectedValue := []int64{11, 42, 59}

		mreader.EXPECT().LongArray().Return([]int64{11, 42, 59}, nil)
		err := tag.Read(mreader)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
}

func TestLongArrayT_Write(t *testing.T) {
	tagName := "tag_name"
	expectedMockErr := "expected_mock_error"
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mwriter := NewMockWriter(ctrl)

	t.Run("Should return an error because the first call to writer.Byte failed", func(t *testing.T) {
		tag := &LongArrayT{}

		mwriter.EXPECT().Byte(gomock.Eq(TagLongArray)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the call to writer.String failed", func(t *testing.T) {
		tag := &LongArrayT{name: tagName}

		mwriter.EXPECT().Byte(gomock.Eq(TagLongArray)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the first call to writer.Int failed", func(t *testing.T) {
		tag := &LongArrayT{name: tagName}

		mwriter.EXPECT().Byte(gomock.Eq(TagLongArray)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Long(gomock.Eq(int64(0))).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should return an error because the second call to writer.Int failed", func(t *testing.T) {
		tag := &LongArrayT{name: tagName, value: []int64{42}}

		mwriter.EXPECT().Byte(gomock.Eq(TagLongArray)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Long(gomock.Eq(int64(1))).Return(nil)
		mwriter.EXPECT().Long(gomock.Eq(int64(42))).Return(errors.New(expectedMockErr))
		err := tag.Write(mwriter, true)
		if assert.Error(t, err) {
			assert.EqualValues(t, expectedMockErr, err.Error())
		}
	})
	t.Run("Should be ok", func(t *testing.T) {
		tag := &LongArrayT{name: tagName, value: []int64{42, 3, 33}}
		expectedValue := []int64{42, 3, 33}

		mwriter.EXPECT().Byte(gomock.Eq(TagLongArray)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Long(gomock.Eq(int64(3))).Return(nil)
		mwriter.EXPECT().Long(gomock.Eq(int64(42))).Return(nil)
		mwriter.EXPECT().Long(gomock.Eq(int64(3))).Return(nil)
		mwriter.EXPECT().Long(gomock.Eq(int64(33))).Return(nil)
		err := tag.Write(mwriter, true)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})
	t.Run("Should be ok with empty list", func(t *testing.T) {
		tag := &LongArrayT{name: tagName, value: []int64{}}
		expectedValue := []int64{}

		mwriter.EXPECT().Byte(gomock.Eq(TagLongArray)).Return(nil)
		mwriter.EXPECT().String(gomock.Eq(tagName)).Return(nil)
		mwriter.EXPECT().Long(gomock.Eq(int64(0))).Return(nil)
		err := tag.Write(mwriter, true)
		if assert.NoError(t, err) {
			assert.EqualValues(t, expectedValue, tag.value)
		}
	})

}
