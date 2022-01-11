package bitstream

type BitStream struct {
	Buffer    []byte
	Length    uint
	byteIndex uint
	bitIndex  uint
}

var (
	hiMask = [9]byte{0x00FF, 0x00FE, 0x00FC, 0x00F8, 0x00F0, 0x00E0, 0x00C0, 0x0080, 0x0000}
	loMask = [9]byte{0x0000, 0x0001, 0x0003, 0x0007, 0x000F, 0x001F, 0x003F, 0x007F, 0x00FF}
)

func ReadByte(b *BitStream) byte {
	if b.bitIndex == 0 {
		b.byteIndex++
		return b.Buffer[b.byteIndex-1]
	}

	hi := (b.Buffer[b.byteIndex] << b.bitIndex) & hiMask[b.bitIndex]
	b.byteIndex++
	lo := (b.Buffer[b.byteIndex] >> (8 - b.bitIndex)) & loMask[b.bitIndex]

	return hi | lo
}

func ReadBytes(b *BitStream, bytes uint) []byte {
	out := make([]byte, bytes)
	if b.bitIndex == 0 {
		copy(out, b.Buffer[b.byteIndex:b.byteIndex+bytes])
		b.byteIndex += bytes
		return out
	}

	for i := 0; i < int(bytes); i++ {
		out[i] = ReadByte(b)
	}

	return out
}

func PeekBit(b *BitStream) byte {
	return (b.Buffer[b.byteIndex] >> (7 - b.bitIndex)) & 0x01
}

func ReadBoolean(b *BitStream) bool {
	bit := PeekBit(b)
	increment(b, 1)
	return bit == 0x01
}

func increment(b *BitStream, bits uint) {
	b.bitIndex += bits
	for b.bitIndex > 7 {
		b.bitIndex -= 8
		b.byteIndex++
	}
}
