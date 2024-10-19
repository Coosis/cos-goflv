package cosgoflv

import (
	"fmt"
)

type FlvAudioTag struct {
	Format SoundFormat
	Rate SamplingRate
	Size SoundSize
	Type SoundType
	Data []byte
}

type SoundFormat uint8
const (
	// 7, 8, 14, 15 are reserved for internal use
	LinearPCMPlatformEndian SoundFormat = iota
	ADPCM
	MP3
	LinearPCMLittleEndian
	Nellymoser16KHzMono
	Nellymoser8KHzMono
	Nellymoser
	G711ALawLogarithmicPCM
	G711MuLawLogarithmicPCM
	Reserved
	AAC
	Speex
	MP38KHz
	DeviceSpecificSound
)
func IsValidSoundFormat(format uint8) error {
	if format > 15 {
		return fmt.Errorf("Invalid sound format: %d", format)
	}
	if format == 7 || format == 8 || format == 14 || format == 15 {
		return fmt.Errorf("Encountered reserved sound format: %d, this is not intended to be handled by the cos-goflv package.", format)
	}
	return nil
}

type SamplingRate uint8
const (
	// Sampling rate. For AAC, it's always 3
	KHZ5P5 SamplingRate = iota
	KHZ11
	KHZ22
	KHZ44
)
func IsValidSamplingRate(rate uint8) error {
	if rate > 3 {
		return fmt.Errorf("Invalid sampling rate: %d", rate)
	}
	return nil
}

type SoundSize uint8
const (
	Bit8 SoundSize = iota
	Bit16
)
func IsValidSoundSize(size uint8) error {
	if size > 1 {
		return fmt.Errorf("Invalid sound size: %d", size)
	}
	return nil
}

type SoundType uint8
const (
	Mono SoundType = iota
	Stereo
)
func IsValidSoundType(soundType uint8) error {
	if soundType > 1 {
		return fmt.Errorf("Invalid sound type: %d", soundType)
	}
	return nil
}

func(audioTag *FlvAudioTag) Parse(data []byte) error {
	if len(data) < 1 {
		return fmt.Errorf("Insufficient data for audio tag.")
	}
	if err := IsValidSoundFormat(data[0] >> 4); err != nil {
		return err
	}
	if err := IsValidSamplingRate((data[0] >> 2) & 0x03); err != nil {
		return err
	}
	if err := IsValidSoundSize((data[0] >> 1) & 0x01); err != nil {
		return err
	}
	if err := IsValidSoundType(data[0] & 0x01); err != nil {
		return err
	}

	audioTag.Format = SoundFormat(data[0] >> 4)
	audioTag.Rate = SamplingRate((data[0] >> 2) & 0x03)
	audioTag.Size = SoundSize((data[0] >> 1) & 0x01)
	audioTag.Type = SoundType(data[0] & 0x01)
	if len(data) == 1 {
		audioTag.Data = nil
		return nil
	}
	audioTag.Data = data[1:]
	return nil
}
