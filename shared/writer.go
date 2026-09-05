package shared

import "math"

type Writer struct { // Little endian
	Bytes []byte
}

func (w *Writer) Reset() {
	w.Bytes = w.Bytes[:0]
}

func (w *Writer) SetI8(value int8) {
	w.Bytes = append(w.Bytes, byte(value))
}

func (w *Writer) SetI16(value int16) {
	w.Bytes = append(w.Bytes, byte(value>>8), byte(value))
}

func (w *Writer) SetI32(value int32) {
	w.Bytes = append(w.Bytes, byte(value>>24), byte(value>>16), byte(value>>8), byte(value))
}

func (w *Writer) SetI64(value int64) {
	w.Bytes = append(w.Bytes, byte(value>>56), byte(value>>48), byte(value>>40), byte(value>>32), byte(value>>24), byte(value>>16), byte(value>>8), byte(value))
}

func (w *Writer) SetU8(value uint8) {
	w.Bytes = append(w.Bytes, value)
}

func (w *Writer) SetU16(value uint16) {
	w.Bytes = append(w.Bytes, byte(value), byte(value>>8))
}

func (w *Writer) SetU32(value uint32) {
	w.Bytes = append(w.Bytes, byte(value), byte(value>>8), byte(value>>16), byte(value>>24))
}

func (w *Writer) SetU64(value uint64) {
	w.Bytes = append(w.Bytes, byte(value), byte(value>>8), byte(value>>16), byte(value>>24), byte(value>>32), byte(value>>40), byte(value>>48), byte(value>>56))
}

func (w *Writer) SetF32(value float32) {
	w.SetU32(math.Float32bits(value))
}

func (w *Writer) SetF64(value float64) {
	w.SetU64(math.Float64bits(value))
}

func (w *Writer) SetStringUTF8(value string) {
	w.Bytes = append(w.Bytes, []byte(value)...)
	w.SetU8(0)
}

func (w *Writer) Append(other *Writer) {
	w.Bytes = append(w.Bytes, other.Bytes...)
}

func (w *Writer) GetBytes() []byte {
	return w.Bytes
}

func (w *Writer) GetLength() int {
	return len(w.Bytes)
}
