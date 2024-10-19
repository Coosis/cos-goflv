package cosgoflv

import (
	"fmt"
)

type AACPacketType uint8
const (
	AAC_SequenceHeader AACPacketType = iota
	AAC_Raw
)
func IsValidAACPacketType(aacPacketType AACPacketType) error {
	switch aacPacketType {
	case AAC_SequenceHeader, AAC_Raw:
		return nil
	default:
		return fmt.Errorf("Invalid AACPacketType: %d", aacPacketType)
	}
}

type FlvAACAudioData struct {
	AACPacketType AACPacketType
	Data []byte
}

func(aacdata *FlvAACAudioData) Parse(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Invalid FlvAACAudioData. Data length is less than 1")
	}
	aacPacketType := uint8(data[0])
	if err := IsValidAACPacketType(AACPacketType(aacPacketType)); err != nil {
		return err
	}
	aacdata.AACPacketType = AACPacketType(aacPacketType)
	if len(data) == 1 {
		aacdata.Data = nil
		return nil
	}
	aacdata.Data = data[1:]
	return nil
}
