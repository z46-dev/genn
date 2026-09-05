package shared

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
)

// Reader decodes little-endian protocol values without panicking on malformed input.
type Reader struct {
	bytes  []byte
	offset int
	err    error
}

// NewReader wraps a packet for bounded decoding.
func NewReader(data []byte) (reader *Reader) {
	reader = &Reader{bytes: data}
	return
}

func (r *Reader) GetI8() (value int8) {
	value = int8(r.GetU8())
	return
}

func (r *Reader) GetI16() (value int16) {
	value = int16(r.GetU16())
	return
}

func (r *Reader) GetI32() (value int32) {
	value = int32(r.GetU32())
	return
}

func (r *Reader) GetI64() (value int64) {
	value = int64(r.GetU64())
	return
}

func (r *Reader) GetU8() (value uint8) {
	if !r.ensure(1) {
		return
	}

	value = r.bytes[r.offset]
	r.offset++
	return
}

func (r *Reader) GetU16() (value uint16) {
	if !r.ensure(2) {
		return
	}

	value = binary.LittleEndian.Uint16(r.bytes[r.offset : r.offset+2])
	r.offset += 2
	return
}

func (r *Reader) GetU32() (value uint32) {
	if !r.ensure(4) {
		return
	}

	value = binary.LittleEndian.Uint32(r.bytes[r.offset : r.offset+4])
	r.offset += 4
	return
}

func (r *Reader) GetU64() (value uint64) {
	if !r.ensure(8) {
		return
	}

	value = binary.LittleEndian.Uint64(r.bytes[r.offset : r.offset+8])
	r.offset += 8
	return
}

func (r *Reader) GetF32() (value float32) {
	value = math.Float32frombits(r.GetU32())
	return
}

func (r *Reader) GetF64() (value float64) {
	value = math.Float64frombits(r.GetU64())
	return
}

func (r *Reader) GetStringUTF8() (value string) {
	if r.err != nil {
		return
	}

	var end int = bytes.IndexByte(r.bytes[r.offset:], 0)
	if end < 0 {
		r.err = io.ErrUnexpectedEOF
		return
	}

	value = string(r.bytes[r.offset : r.offset+end])
	r.offset += end + 1
	return
}

func (r *Reader) PeekI8() (value int8) {
	value = int8(r.PeekU8())
	return
}

func (r *Reader) PeekI16() (value int16) {
	value = int16(r.PeekU16())
	return
}

func (r *Reader) PeekI32() (value int32) {
	value = int32(r.PeekU32())
	return
}

func (r *Reader) PeekI64() (value int64) {
	value = int64(r.PeekU64())
	return
}

func (r *Reader) PeekU8() (value uint8) {
	if r.ensure(1) {
		value = r.bytes[r.offset]
	}

	return
}

func (r *Reader) PeekU16() (value uint16) {
	if r.ensure(2) {
		value = binary.LittleEndian.Uint16(r.bytes[r.offset : r.offset+2])
	}

	return
}

func (r *Reader) PeekU32() (value uint32) {
	if r.ensure(4) {
		value = binary.LittleEndian.Uint32(r.bytes[r.offset : r.offset+4])
	}

	return
}

func (r *Reader) PeekU64() (value uint64) {
	if r.ensure(8) {
		value = binary.LittleEndian.Uint64(r.bytes[r.offset : r.offset+8])
	}

	return
}

func (r *Reader) PeekF32() (value float32) {
	value = math.Float32frombits(r.PeekU32())
	return
}

func (r *Reader) PeekF64() (value float64) {
	value = math.Float64frombits(r.PeekU64())
	return
}

// Err reports the first packet-boundary error encountered by the reader.
func (r *Reader) Err() (err error) {
	err = r.err
	return
}

// Remaining reports the unread byte count.
func (r *Reader) Remaining() (remaining int) {
	remaining = len(r.bytes) - r.offset
	return
}

// Rewind moves back by the width of numericType for diagnostic call sites.
func (r *Reader) Rewind(numericType any) *Reader {
	var width int

	switch numericType.(type) {
	case *int8, *uint8, int8, uint8:
		width = 1
	case *int16, *uint16, int16, uint16:
		width = 2
	case *int32, *uint32, *float32, int32, uint32, float32:
		width = 4
	case *int64, *uint64, *float64, int64, uint64, float64:
		width = 8
	}

	r.offset = max(0, r.offset-width)
	return r
}

// ensure validates that size bytes are readable from the current offset.
func (r *Reader) ensure(size int) (ok bool) {
	if r.err != nil {
		return
	}

	if size < 0 || r.offset > len(r.bytes)-size {
		r.err = io.ErrUnexpectedEOF
		return
	}

	ok = true
	return
}
