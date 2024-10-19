package cosgoflv

import (
	"fmt"
	"io"
	"encoding/binary"
)

const (
	ASCII_F = 0x46
	ASCII_L = 0x4C
	ASCII_V = 0x56
	V1	  = 0x01
)

func EmptyFlvHeader() *FlvHeader {
	return &FlvHeader{
		Version: V1,
		TypeFlags: 0,
		DataOffset: 9,
	}
}

type FlvHeader struct {
	Version uint8

	TypeFlags uint8

	DataOffset uint32
}

func(header *FlvHeader) Parse(data []byte) error {
	if len(data) < 9 {
		return fmt.Errorf("FLV header too short, file may be corrupted")
	}
	// First 3 bytes are signature, always "FLV"
	if string(data[0:3]) != "FLV" {
		return fmt.Errorf("Invalid FLV signature")
	}

	header.Version = data[3]

	header.TypeFlags = data[4]

	header.DataOffset = binary.BigEndian.Uint32(data[5:9])
	return nil
}

func(header *FlvHeader) Write(w io.Writer) error {
	// Signature "FLV"
	if _, err := w.Write([]byte("FLV")); err != nil {
		return err
	}

	// Version
	if _, err := w.Write([]byte{header.Version}); err != nil {
		return err
	}

	// TypeFlags
	if _, err := w.Write([]byte{header.TypeFlags}); err != nil {
		return err
	}

	// DataOffset
	if err := binary.Write(w, binary.BigEndian, header.DataOffset); err != nil {
		return err
	}

	return nil
}

func(header *FlvHeader) Read(r io.Reader) error {
	// Signature "FLV"
	signature := make([]byte, 3)
	if _, err := io.ReadFull(r, signature); err != nil {
		return err
	}
	if string(signature) != "FLV" {
		return fmt.Errorf("Invalid FLV signature: %s", signature)
	}

	// Version
	if err := binary.Read(r, binary.BigEndian, &header.Version); err != nil {
		return err
	}
	if header.Version != V1 {
		return fmt.Errorf("FLV version %d not supported", header.Version)
	}

	// TypeFlags
	if err := binary.Read(r, binary.BigEndian, &header.TypeFlags); err != nil {
		return err
	}
	if (header.TypeFlags & 0xFA) != 0 {
		return fmt.Errorf("Invalid TypeFlags: %d", header.TypeFlags)
	}

	// DataOffset
	if err := binary.Read(r, binary.BigEndian, &header.DataOffset); err != nil {
		return err
	}
	if header.DataOffset != 9 {
		return fmt.Errorf("Invalid DataOffset: %d", header.DataOffset)
	}

	return nil
}
