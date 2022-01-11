package bitstream_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.darcro.dev/learn-go/bitstream"
)

func TestReadByteAligned(t *testing.T) {
	testDat := []byte{0x11, 0x22, 0x33, 0x44, 0x55}

	var bs = new(bitstream.BitStream)
	bs.Buffer = testDat

	var b = bitstream.ReadByte(bs)

	if b != 0x11 {
		t.Fatal("Could not read first value")
	}
	b = bitstream.ReadByte(bs)
	if b != 0x22 {
		t.Fatal("Could not read second value")
	}
	b = bitstream.ReadByte(bs)
	if b != 0x33 {
		t.Fatal("Could not read third value")
	}
	b = bitstream.ReadByte(bs)
	if b != 0x44 {
		t.Fatal("Could not read fourth value")
	}
	b = bitstream.ReadByte(bs)
	if b != 0x55 {
		t.Fatal("Could not read fifth value")
	}
}

func TestReadBytesAligned(t *testing.T) {
	testDat := []byte{0x11, 0x22, 0x33, 0x44, 0x55}

	var bs = new(bitstream.BitStream)
	bs.Buffer = testDat

	var b = bitstream.ReadBytes(bs, 2)
	if !bytes.Equal(b, []byte{0x11, 0x22}) {
		t.Fatal("Could not read first bytes")
	}

	b = bitstream.ReadBytes(bs, 2)

	if !bytes.Equal(b, []byte{0x33, 0x44}) {
		t.Fatal("Could not read second bytes")
	}
}

func TestReadBool(t *testing.T) {
	testDat := []byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA}

	var bs = new(bitstream.BitStream)
	bs.Buffer = testDat

	var b = bitstream.ReadBoolean(bs)

	if !b {
		t.Fatal("Could not read first value")
	}
	b = bitstream.ReadBoolean(bs)
	if b {
		t.Fatal("Could not read second value")
	}
	b = bitstream.ReadBoolean(bs)
	if !b {
		t.Fatal("Could not read third value")
	}
	b = bitstream.ReadBoolean(bs)
	if b {
		t.Fatal("Could not read fourth value")
	}
	b = bitstream.ReadBoolean(bs)
	if !b {
		t.Fatal("Could not read fifth value")
	}
}

func TestReadByte(t *testing.T) {
	testDat := []byte{0xF0, 0xF0, 0xF0, 0xF0, 0xF0}

	var bs = new(bitstream.BitStream)
	bs.Buffer = testDat

	var bit = bitstream.ReadBoolean(bs)
	if !bit {
		t.Fatal("Initial bit invalid")
	}

	var b = bitstream.ReadByte(bs)
	if b != 0xE1 {
		t.Fatal("Could not read first value: " + hex.EncodeToString([]byte{b}))
	}

	bit = bitstream.ReadBoolean(bs)
	if !bit {
		t.Fatal("Boolean invalid")
	}

	b = bitstream.ReadByte(bs)
	if b != 0xC3 {
		t.Fatal("Could not read second value: " + hex.EncodeToString([]byte{b}))
	}

}
