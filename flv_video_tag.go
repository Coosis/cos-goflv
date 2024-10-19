package cosgoflv

import (
	"fmt"
)

type FlvVideoTag struct {
	FrameType FrameType
	CodecID CodecID
	VideoData []byte
}

type FrameType uint8
const (
	KeyFrame FrameType = iota+1
	InterFrame
	DisposableInterFrame
	GeneratedKeyFrame
	// Equivalent to video info/command frame
	VideoFrame
)
func IsValidFrameType(frameType uint8) error {
	if frameType > 5 {
		return fmt.Errorf("Invalid frame type: %d", frameType)
	}
	if frameType < 1 {
		return fmt.Errorf("Invalid frame type: %d", frameType)
	}
	return nil
}

type CodecID uint8
const (
	JPEG CodecID = iota+1
	SorensonH263
	ScreenVideo
	On2VP6
	On2VP6Alpha
	ScreenVideov2
	AVC
)
func IsValidCodecID(codecID uint8) error {
	if codecID > 7 {
		return fmt.Errorf("Invalid codec ID: %d", codecID)
	}
	if codecID < 1 {
		return fmt.Errorf("Invalid codec ID: %d", codecID)
	}
	return nil
}

func(videoTag *FlvVideoTag) Parse(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Insufficient data for video tag.")
	}
	err := IsValidFrameType(data[0] >> 4)
	if err != nil {
		return err
	}
	err = IsValidCodecID(data[0] & 0x0f)
	if err != nil {
		return err
	}
	videoTag.FrameType = FrameType(data[0] >> 4)
	videoTag.CodecID = CodecID(data[0] & 0x0f)
	if len(data) == 1 {
		videoTag.VideoData = nil
		return nil
	}
	videoTag.VideoData = data[1:]
	return nil
}
