package cosgoflv

import (
	"fmt"
)

type AVCPacketType uint8
const (
	AVC_SequenceHeader AVCPacketType = iota
	AVC_NALU
	AVC_EndOfSequence
)

type FlvAvcVideoPacket struct {
	AVCPacketType AVCPacketType
	CompositionTime int32
	Data []byte
}

func(avp *FlvAvcVideoPacket) Parse(data []byte) error {
	if len(data) < 4 {
		return fmt.Errorf("Invalid FlvAvcVideoPacket. Data length is less than 4")
	}
	avp.AVCPacketType = AVCPacketType(data[0])
	avp.CompositionTime = int32(data[1]) << 16 | int32(data[2]) << 8 | int32(data[3])
	if len(data) == 4 {
		avp.Data = nil
		return nil
	}
	avp.Data = data[4:]
	return nil
}
