package gonbt

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// constant compression type
const (
	CompressGZIP = "gzip"
	CompressZLIB = "zlib"
	CompressNone = "none"
)

// read and write data with the RFC NBT describe here:
// https://minecraft.gamepedia.com/NBT_format

// Unmarshal data
func Unmarshal(data []byte) (Tag, error) {
	var err error
	var reader Reader
	var tagT byte
	var name string
	var t Tag
	var driver io.Reader

	driver = bytes.NewReader(data)
	contentType := http.DetectContentType(data)
	if contentType == "application/x-gzip" {
		var gReader *gzip.Reader
		if gReader, err = gzip.NewReader(driver); err != nil {
			return nil, err
		}
		defer gReader.Close()

		if data, err = ioutil.ReadAll(gReader); err != nil {
			return nil, err
		}
		driver = bytes.NewBuffer(data)
	}

	reader = NewReader(driver)
	if tagT, err = reader.Byte(); err != nil {
		return nil, err
	}
	if name, err = reader.String(); err != nil {
		return nil, err
	}

	if t, err = NewTag(tagT, name); err != nil && err.Error() != errorEnd {
		return nil, err
	}
	if err == nil {
		if err = t.Read(reader); err != nil {
			return nil, err
		}
	}

	return t, nil
}

// Marshal data
func Marshal(t Tag, compress string) ([]byte, error) {
	var err error
	var driver io.Writer
	var buf *bytes.Buffer
	var writer Writer
	var output []byte

	buf = bytes.NewBuffer([]byte{})
	driver = buf

	writer = NewWriter(driver)
	if err = t.Write(writer, true); err != nil {
		return []byte{}, err
	}

	switch compress {
	case CompressGZIP:
		gzipBuf := bytes.NewBuffer([]byte{})
		var gw *gzip.Writer
		gw, err = gzip.NewWriterLevel(gzipBuf, gzip.BestSpeed)
		defer gw.Close()
		if _, err = gw.Write(buf.Bytes()); err != nil {
			return []byte{}, nil
		}
		output = gzipBuf.Bytes()
	case CompressZLIB:
	case CompressNone:
		output = buf.Bytes()
	default:
		return []byte{}, errors.New(errorCompressType)
	}

	return output, nil
}
