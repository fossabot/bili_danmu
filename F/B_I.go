package F

import (
	"bytes"
	"encoding/binary"

	p "github.com/qydysky/part"
)

/*
整数 字节转换区
64 8字节
32 4字节
16 2字节
*/
func Itob64(num int64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		p.Logf().E(err)
	}
	return buffer.Bytes()
}

func Itob32(num int32) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		p.Logf().E(err)
	}
	return buffer.Bytes()
}

func Itob16(num int16) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err != nil {
		p.Logf().E(err)
	}
	return buffer.Bytes()
}

func Btoi64(b []byte, offset int) int64 {
	for len(b) < 8 {
		b = append([]byte{0x00}, b...)
	}
	return btoi64(b[offset : offset+8])
}

func Btoi(b []byte, offset int, size int) int64 {
	if size > 8 {
		panic("最大8位")
	}
	var buf = b[offset : offset+size]
	for len(buf) < 8 {
		buf = append([]byte{0x00}, buf...)
	}
	return btoi64(buf)
}

func Btoui32(b []byte, offset int) uint32 {
	return btoui32(b[offset : offset+4])
}

func Btoi32(b []byte, offset int) int32 {
	for len(b) < 4 {
		b = append([]byte{0x00}, b...)
	}
	return int32(btoui32(b[offset : offset+4]))
}

func Btoui16(b []byte, offset int) uint16 {
	for len(b) < 2 {
		b = append([]byte{0x00}, b...)
	}
	return btoui16(b[offset : offset+2])
}

func Btoi16(b []byte, offset int) int16 {
	for len(b) < 2 {
		b = append([]byte{0x00}, b...)
	}
	return int16(btoui16(b[offset : offset+2]))
}

func btoi64(b []byte) int64 {
	var buffer int64
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &buffer)
	if err != nil {
		p.Logf().E(err)
	}
	return buffer
}

func btoui32(b []byte) uint32 {
	var buffer uint32
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &buffer)
	if err != nil {
		p.Logf().E(err)
	}
	return buffer
}

func btoui16(b []byte) uint16 {
	var buffer uint16
	err := binary.Read(bytes.NewReader(b), binary.BigEndian, &buffer)
	if err != nil {
		p.Logf().E(err)
	}
	return buffer
}
